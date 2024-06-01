package util

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "2clip",
	Short: "2clip is a simple clipboard manager",
	Long:  `2clip is a simple CLI tool for managing your clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2clip is a simple CLI tool for managing your clipboard")
	},
}