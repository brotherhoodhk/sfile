/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get file from local sfile virtual directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			sfile.Error()
			return
		} else {
			for _, name := range args {
				sfile.GetFile(name)
			}
		}
		// sfile.GetFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
