/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	"fmt"
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull a file from cloud",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			sfile.PullFile(args[0])
		case 2:
			if args[0] == "--private" && sfile.Isprivatefilename(args[1]) {
				// CommonExchangeFile(args[2], 42)
				if auth, ok := sfile.GetAuthInfo(); ok {
					sfile.CommonExchangeFilePlus(args[1], auth, 842)
				} else {
					fmt.Println(sfile.AUTHGETWARN)
				}
			} else {
				sfile.Error()
			}
		default:
			sfile.Error()
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
