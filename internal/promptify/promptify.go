package promptify

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	gitignore "github.com/sabhiram/go-gitignore"
)

// Options defines the configurable parameters for generating the prompt.
type Options struct {
	// MaxDepth is the maximum depth of directory traversal.
	// If <= 0, there is no limit.
	MaxDepth int

	// RootDir is the directory from which to begin collecting file information.
	RootDir string

	// FileFormat is the format/template for each file's content, e.g.:
	//   "<FILE name=\"{{.FileName}}\">\n{{.Content}}\n</FILE>"
	FileFormat string

	// PromptIntro is the template that is prepended before listing all file formats, e.g.:
	//   "The contents below represent a directory '{{.Root}}' and its file contents..."
	PromptIntro string

	// IgnorePatterns is a list of file/directory name patterns that should be ignored
	// in addition to the normal .gitignore (and the implicit .git folder).
	// Patterns can be e.g. "*.md" or "node_modules".
	IgnorePatterns []string

	// DryRun, if true, indicates that we only want to list the files
	// that would be included (skipping content retrieval and templating).
	DryRun bool
}

// PromptifyData is the data used by the top-level PromptIntro template.
type PromptifyData struct {
	Root       string
	FileFormat string
}

// FileData is the data passed into the FileFormat template for each file.
type FileData struct {
	FileName string
	Content  string
}

// Promptify generates a single string containing an introduction (via PromptIntro)
// and the contents of each file (via FileFormat), respecting .gitignore, max depth,
// and additional ignore patterns. It also always ignores the .git folder by default.
//
// If opts.DryRun == true, then it simply returns a list of included file names
// (one per line) without reading file contents or rendering templates.
func Promptify(opts Options) (string, error) {
	if opts.RootDir == "" {
		return "", errors.New("root directory must be specified")
	}

	// 1. Prepare the .gitignore matcher (ignore errors if .gitignore not found).
	var ign *gitignore.GitIgnore
	gitIgnorePath := filepath.Join(opts.RootDir, ".gitignore")
	if fi, err := os.Stat(gitIgnorePath); err == nil && !fi.IsDir() {
		ign, _ = gitignore.CompileIgnoreFile(gitIgnorePath)
	}

	// 2. Collect all files up to max depth (if > 0), respecting .gitignore and user ignore patterns.
	fileInfos, err := collectFiles(opts.RootDir, ign, opts.IgnorePatterns, opts.MaxDepth)
	if err != nil {
		return "", err
	}

	// 3. If we're just doing a dry run, return the filenames only.
	if opts.DryRun {
		var buf bytes.Buffer
		for _, path := range fileInfos {
			buf.WriteString(path + "\n")
		}
		return buf.String(), nil
	}

	// 4. Otherwise, parse templates and render the actual prompt.

	// 4a. Parse the introduction template.
	introTmpl, err := template.New("intro").Parse(opts.PromptIntro)
	if err != nil {
		return "", fmt.Errorf("failed to parse PromptIntro template: %w", err)
	}

	// 4b. Parse the file format template.
	fileTmpl, err := template.New("fileFormat").Parse(opts.FileFormat)
	if err != nil {
		return "", fmt.Errorf("failed to parse FileFormat template: %w", err)
	}

	// 5. Build the result in a string buffer.
	var buf bytes.Buffer

	// 5a. Render the intro template first.
	introData := PromptifyData{
		Root:       opts.RootDir,
		FileFormat: opts.FileFormat,
	}
	if err := introTmpl.Execute(&buf, introData); err != nil {
		return "", fmt.Errorf("failed to execute PromptIntro template: %w", err)
	}
	if !strings.HasSuffix(buf.String(), "\n") {
		buf.WriteString("\n")
	}
	buf.WriteString("\n")

	// 5b. For each file, read contents & apply file template.
	for _, path := range fileInfos {
		content, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("failed to read file %q: %w", path, err)
		}

		fileData := FileData{
			FileName: path,
			Content:  string(content),
		}

		if err := fileTmpl.Execute(&buf, fileData); err != nil {
			return "", fmt.Errorf("failed to execute FileFormat template for file %q: %w", path, err)
		}
		if !strings.HasSuffix(buf.String(), "\n") {
			buf.WriteString("\n")
		}
		buf.WriteString("\n")
	}

	return buf.String(), nil
}

// collectFiles walks through the directory up to maxDepth (if > 0),
// collecting files that are *not* ignored by the .gitignore parser (if provided),
// and are not matched by any custom ignore patterns. Also implicitly ignores the .git folder.
func collectFiles(root string, ign *gitignore.GitIgnore, userIgnores []string, maxDepth int) ([]string, error) {
	var result []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		// We don't want to process the root as a file
		if path == root {
			return nil
		}

		// If this is the .git folder, skip it (and everything inside).
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		// Compute depth by counting separators in the relative path
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		depth := len(strings.Split(rel, string(filepath.Separator)))

		// If maxDepth > 0, skip deeper dirs/files
		if maxDepth > 0 && depth > maxDepth {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// 1) .gitignore check
		if ign != nil && ign.MatchesPath(rel) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// 2) user-specified ignore patterns
		for _, pattern := range userIgnores {
			matched, matchErr := filepath.Match(pattern, info.Name())
			if matchErr != nil {
				// If the pattern is invalid, skip or handle as needed
				continue
			}
			if matched {
				// If the name matches the pattern, skip this file/dir
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// If we're a directory, keep walking
		if info.IsDir() {
			return nil
		}

		// If we're a file, add it
		result = append(result, path)
		return nil
	})

	if err != nil && err != io.EOF {
		return nil, err
	}
	return result, nil
}
