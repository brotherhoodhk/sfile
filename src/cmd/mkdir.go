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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mkdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mkdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
