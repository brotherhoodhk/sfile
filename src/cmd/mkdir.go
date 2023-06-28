/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	"fmt"
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// mkdirCmd represents the mkdir command
var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "make a private dir in cloud",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			sfile.Error()
			return
		}
		// CommonAgreenment(args[1], 40)
		if auth, ok := sfile.GetAuthInfo(); ok {
			sfile.CommonAgreenmentSecure(args[0], auth, 840)
		} else {
			fmt.Println(sfile.AUTHGETWARN)
		}
	},
}

func init() {
	rootCmd.AddCommand(mkdirCmd)
}
