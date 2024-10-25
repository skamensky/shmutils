package cmd

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/skamensky/shmutils/internal/calc"
	"github.com/skamensky/shmutils/internal/docker"
	"github.com/skamensky/shmutils/internal/promptify"
	"github.com/skamensky/shmutils/internal/tz"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shmutils",
	Short: "My personal utility packages. Maybe someone else will find this interesting. Probably not.",
	//Long: `The CLI `,
	// TODO when run without args, use promptui to select a command
	//Run: func(cmd *cobra.Command, args []string) {
	//
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// todo make a clid.HandleError function

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shmutils.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dockerCmd := &cobra.Command{
		Use:   "docker",
		Short: "Docker utilities",
		Run: func(cmd *cobra.Command, args []string) {
			docker.ShellIntoContainer()
		},
	}
	calcCmd := &cobra.Command{
		Use:   `calc "[expression]"`,
		Short: "Calculator",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := calc.Calculate(strings.Join(args, " "), len(args) < 1)
			if err != nil {
				fmt.Printf("Error calculating expression: %s\n", err)
			} else {
				fmt.Println(res)
			}
		},
	}
	randPass := &cobra.Command{
		Use: `randpass`,

		Short: "Random password generator",
		Run: func(cmd *cobra.Command, args []string) {
			shouldSpecial, err := cmd.Flags().GetBool("special")
			if err != nil {
				fmt.Println("Error getting special flag")
				return
			}
			shouldNumbers, err := cmd.Flags().GetBool("numbers")
			if err != nil {
				fmt.Println("Error getting numbers flag")
				return
			}
			length, err := cmd.Flags().GetInt("length")
			if err != nil {
				fmt.Println("Error getting length flag")
				return
			}
			shouldUpper, err := cmd.Flags().GetBool("upper")
			if err != nil {
				fmt.Println("Error getting uppercase flag")
				return
			}
			shouldLower, err := cmd.Flags().GetBool("lower")
			if err != nil {
				fmt.Println("Error getting lowercase flag")
				return
			}

			shouldConfusing, err := cmd.Flags().GetBool("confusing")
			if err != nil {
				fmt.Println("Error getting letters confusing")
				return
			}
			specialCharSet := "!@#$%^&*()_+"
			// missing some characters on purpose which are found in the "confusing" set
			numberCharSet := "23456789"
			upperCharSet := "ABCDEFGHJKLMNPQRSTUVWXYZ"
			lowerCharSet := "abcdefghijkmnpqrstuvwxyz"
			confusingSet := "0Oo1Il"
			characterSets := []string{}

			if shouldSpecial {
				characterSets = append(characterSets, specialCharSet)
			}
			if shouldNumbers {
				characterSets = append(characterSets, numberCharSet)
			}
			if shouldUpper {
				characterSets = append(characterSets, upperCharSet)
			}
			if shouldLower {
				characterSets = append(characterSets, lowerCharSet)
			}
			if shouldConfusing {
				characterSets = append(characterSets, confusingSet)
			}
			if len(characterSets) == 0 {
				fmt.Println("You must select at least one character type")
				return
			}

			rand.Seed(time.Now().UnixNano())
			result := ""
			for i := 0; i < length; i++ {
				chosenSet := rand.Intn(len(characterSets))
				chosenChar := rand.Intn(len(characterSets[chosenSet]))
				result += string(characterSets[chosenSet][chosenChar])
			}
			fmt.Println(result)
		},
	}
	randPass.Flags().Int("length", 16, "Length of password")
	randPass.Flags().Bool("special", false, "Include special characters")
	randPass.Flags().Bool("numbers", true, "Include numbers")
	randPass.Flags().Bool("upper", true, "Include uppercase letters")
	randPass.Flags().Bool("lower", false, "Include lowercase letters")
	randPass.Flags().Bool("confusing", false, "Include visually confusing characters like lowercase L and uppercase I")

	tzCmd := &cobra.Command{
		Use:   `tz "[query]"`,
		Short: "Converts times between timezones",
		Run: func(cmd *cobra.Command, args []string) {
			query := strings.Join(args, " ")
			res, err := tz.ExecuteQuery(query)
			if err != nil {
				fmt.Printf("Error executing query: %s\n", err)
			} else {
				fmt.Println(res)
			}
		},
	}

	webServer := &cobra.Command{
		Use:   `server`,
		Short: "Starts an http file system server",
		Run: func(cmd *cobra.Command, args []string) {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				fmt.Println("Error getting port flag")
				return
			}
			dir, err := cmd.Flags().GetString("dir")
			if err != nil {
				fmt.Println("Error getting dir flag")
				return
			}
			fmt.Printf("Starting fileserver on 127.0.0.1:%d serving directory %s\n", port, dir)
			err = http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
			if err != nil {
				fmt.Println("Failed to start server", err)
				return
			}
		},
	}
	webServer.Flags().Int("port", 8080, "Port to listen on")
	webServer.Flags().String("dir", ".", "Directory to serve")

	promptifyCmd := &cobra.Command{
		Use:   "promptify [directory]",
		Short: "Generates a directory tree prompt with file contents, respecting .gitignore",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			maxDepth, err := cmd.Flags().GetInt("maxdepth")
			if err != nil {
				fmt.Println("Error getting maxdepth flag:", err)
				return
			}

			fileFormat, err := cmd.Flags().GetString("format")
			if err != nil {
				fmt.Println("Error getting format flag:", err)
				return
			}

			promptIntro, err := cmd.Flags().GetString("intro")

			if err != nil {
				fmt.Println("Error getting intro flag:", err)
				return
			}

			rootDir := args[0]

			opts := promptify.Options{
				MaxDepth:    maxDepth,
				RootDir:     rootDir,
				FileFormat:  fileFormat,
				PromptIntro: promptIntro,
			}

			result, err := promptify.Promptify(opts)
			if err != nil {
				fmt.Printf("Error generating prompt: %s\n", err)
				return
			}

			fmt.Print(result)
		},
	}

	promptifyCmd.Flags().Int("maxdepth", 1, "Maximum depth to traverse (default 1)")
	promptifyCmd.Flags().String("format", "<FILE name=\"{{.FileName}}\">\n{{.Content}}\n</FILE>",
		"Format template for file contents")
	promptifyCmd.Flags().String("intro", "The contents below represent a directory '{{.Root}}' and its file contents. Files are delimited by ```{{.FileFormat}}```\n", "Introduction template")

	rootCmd.AddCommand(promptifyCmd)
	rootCmd.AddCommand(promptifyCmd)
	rootCmd.AddCommand(dockerCmd)
	rootCmd.AddCommand(calcCmd)
	rootCmd.AddCommand(randPass)
	rootCmd.AddCommand(tzCmd)
	rootCmd.AddCommand(webServer)
}
