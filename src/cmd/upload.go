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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
