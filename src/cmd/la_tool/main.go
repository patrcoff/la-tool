package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/patrcoff/la-tool/src/cli/commands"
	"github.com/patrcoff/la-tool/src/pkg/config"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "logicappstool",
		Short: "A tool for managing Azure Logic App configurations",
		Long: `logicappstool is a CLI tool for comparing and validating 
Azure Logic App configurations across different sources including 
directories, git repositories, and deployed instances.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.latool.yaml)")

	rootCmd.AddCommand(commands.DiffCmd())
	rootCmd.AddCommand(commands.ValidateCmd())
}

func initConfig() {
	if err := config.Init(cfgFile); err != nil {
		fmt.Println("Error initializing config:", err)
		os.Exit(1)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
