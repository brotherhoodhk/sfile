/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show all file record from local sfile virtual directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sfile.ShowList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
