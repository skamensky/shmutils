package cmd

import (
	"fmt"
	"os"
	"strings"

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
	rootCmd.AddCommand(dockerCmd)
	rootCmd.AddCommand(calcCmd)
}
