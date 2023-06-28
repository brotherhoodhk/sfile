/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search file from cloud",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			sfile.Search(args[1])
		} else {
			sfile.Error()
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
