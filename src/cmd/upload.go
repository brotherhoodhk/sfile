/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	sfile "sfile/sfile_body"
	"strings"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload a file to cloud",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 2:
			sfile.UploadFile(args[1])
		case 3:
			if args[1] == "--private" {
				if strings.ContainsRune(args[2], '/') {
					sfile.Uploadprivatefile(args[2])
				}
			}
		default:
			sfile.Error()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
