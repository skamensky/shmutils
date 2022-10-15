package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/skamensky/shmutils/internal/calc"
	"github.com/skamensky/shmutils/internal/docker"
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
			shouldLetters, err := cmd.Flags().GetBool("letters")
			if err != nil {
				fmt.Println("Error getting letters flag")
				return
			}
			specialChacterSet := "!@#$%^&*()_+"
			numberCharacterSet := "1234567890"
			letterCharacterSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

			characterSets := []string{}

			if shouldSpecial {
				characterSets = append(characterSets, specialChacterSet)
			}
			if shouldNumbers {
				characterSets = append(characterSets, numberCharacterSet)
			}
			if shouldLetters {
				characterSets = append(characterSets, letterCharacterSet)
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
	randPass.Flags().Bool("letters", true, "Include letters")

	rootCmd.AddCommand(dockerCmd)
	rootCmd.AddCommand(calcCmd)
	rootCmd.AddCommand(randPass)
}
