/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "dev test mode.don't use!!!",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Test()
		sfile.TestMkdir()
		sfile.TestUploadPrivateFile()
		sfile.TestPullFile()
		sfile.TestDeleteFile()
		sfile.TestDeleteDir()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
