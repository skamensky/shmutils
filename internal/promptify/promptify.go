// File: internal/promptify/promptify.go

package promptify

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

type Options struct {
	MaxDepth    int
	RootDir     string
	FileFormat  string
	PromptIntro string
}

type FileNode struct {
	Name     string
	Path     string
	Content  string
	IsDir    bool
	Children []*FileNode
}

type FormatData struct {
	FileName string
	Content  string
}

type IntroData struct {
	Root       string
	FileFormat string
}

// loadGitignore loads the .gitignore from the root directory and returns a matcher
func loadGitignore(rootDir string) (gitignore.Matcher, error) {
	patterns := []gitignore.Pattern{}

	gitignorePath := filepath.Join(rootDir, ".gitignore")
	if _, err := os.Stat(gitignorePath); err == nil {
		content, err := os.ReadFile(gitignorePath)
		if err != nil {
			return nil, fmt.Errorf("error reading .gitignore: %w", err)
		}

		for _, line := range strings.Split(string(content), "\n") {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				patterns = append(patterns, gitignore.ParsePattern(line, nil))
			}
		}
	}

	return gitignore.NewMatcher(patterns), nil
}

func Promptify(opts Options) (string, error) {
	if opts.MaxDepth < 1 {
		return "", fmt.Errorf("maxdepth must be at least 1")
	}

	// Load .gitignore once at the root
	matcher, err := loadGitignore(opts.RootDir)
	if err != nil {
		return "", err
	}

	// Build the file tree
	tree := &FileNode{
		Name:     filepath.Base(opts.RootDir),
		Path:     opts.RootDir,
		IsDir:    true,
		Children: make([]*FileNode, 0),
	}

	err = buildTree(tree, opts.MaxDepth, 1, matcher, opts.RootDir)
	if err != nil {
		return "", fmt.Errorf("error building file tree: %w", err)
	}

	// Generate the prompt
	var sb strings.Builder

	// Add prompt introduction if provided
	if opts.PromptIntro != "" {
		introTmpl, err := template.New("intro").Parse(opts.PromptIntro)
		if err != nil {
			return "", fmt.Errorf("error parsing intro template: %w", err)
		}

		introData := IntroData{
			Root:       opts.RootDir,
			FileFormat: opts.FileFormat,
		}

		var introBuf bytes.Buffer
		if err := introTmpl.Execute(&introBuf, introData); err != nil {
			return "", fmt.Errorf("error executing intro template: %w", err)
		}
		sb.WriteString(introBuf.String())
		sb.WriteString("\n\n")
	}

	// Parse file format template
	formatTmpl, err := template.New("format").Parse(opts.FileFormat)
	if err != nil {
		return "", fmt.Errorf("error parsing file format template: %w", err)
	}

	err = generatePrompt(tree, formatTmpl, &sb)
	if err != nil {
		return "", fmt.Errorf("error generating prompt: %w", err)
	}

	return sb.String(), nil
}

func buildTree(node *FileNode, maxDepth, currentDepth int, matcher gitignore.Matcher, rootDir string) error {
	if currentDepth > maxDepth {
		return nil
	}

	entries, err := os.ReadDir(node.Path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Name() == ".git" {
			continue
		}

		path := filepath.Join(node.Path, entry.Name())

		// Get path relative to root directory for gitignore matching
		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			continue
		}

		// Convert to forward slashes for consistent matching
		relPath = filepath.ToSlash(relPath)

		// Check if file/directory should be ignored
		if matcher.Match(strings.Split(relPath, "/"), entry.IsDir()) {
			continue
		}

		fileNode := &FileNode{
			Name:     entry.Name(),
			Path:     path,
			IsDir:    entry.IsDir(),
			Children: make([]*FileNode, 0),
		}

		if !entry.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			fileNode.Content = string(content)
		} else {
			err := buildTree(fileNode, maxDepth, currentDepth+1, matcher, rootDir)
			if err != nil {
				continue
			}
		}

		node.Children = append(node.Children, fileNode)
	}

	return nil
}

func generatePrompt(node *FileNode, formatTmpl *template.Template, sb *strings.Builder) error {
	if !node.IsDir {
		data := FormatData{
			FileName: node.Path,
			Content:  node.Content,
		}

		var buf bytes.Buffer
		if err := formatTmpl.Execute(&buf, data); err != nil {
			return err
		}
		sb.WriteString(buf.String())
		sb.WriteString("\n")
	}

	for _, child := range node.Children {
		if err := generatePrompt(child, formatTmpl, sb); err != nil {
			return err
		}
	}

	return nil
}
