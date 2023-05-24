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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
