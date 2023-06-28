/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove local file record which in sfile virtual directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			sfile.Error()
			return
		} else {
			sfile.RemoveFile(args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
