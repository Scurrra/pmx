/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	pmx_version = "0.0.1"
	noInteract  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pmx",
	Version: pmx_version,
	Short:   "Project manager for C/C++.",
	Long: `
Project manager for C/C++, that helps you to create, document, test and deliver 
binary applications and libraries. 
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)

	rootCmd.PersistentFlags().BoolVar(&noInteract, "no-interact", false, "Ask for missing flags in interactive mod or not.")

	initCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the project.")
	initCmd.Flags().StringVar(&projectType, "project_type", "binary", "Project type.") // prompt is preferred
	initCmd.Flags().StringSliceVar(&authors, "authors", []string{}, "Authors of the project.")
	initCmd.Flags().StringVarP(&version, "version", "v", "v0.0.1", "Version of the project.")
	initCmd.Flags().StringVar(&license, "license", "", "License of the project.")
	initCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the project.")
	initCmd.Flags().StringVarP(&repository, "repository", "r", "", "Url of the remote repository.")
}
