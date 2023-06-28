/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	"fmt"
	sfile "sfile/sfile_body"
	"strings"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean cloud directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 1:
			sfile.CleanFile(args[1])
		case 2:
			if args[0] == "--private" {
				if strings.ContainsRune(args[1], '/') {
					// CommonAgreenment(args[2], 431)
					if auth, ok := sfile.GetAuthInfo(); ok {
						sfile.CommonAgreenmentSecure(args[1], auth, 8431)
					} else {
						fmt.Println(sfile.AUTHGETWARN)
					}
				} else {
					fmt.Println("filename is not correct")
				}
			} else if args[0] == "-r" {
				if strings.ContainsRune(args[1], '/') {
					fmt.Println("your dirname is not correct")
				} else {
					// CommonAgreenment(args[2], 43)
					if auth, ok := sfile.GetAuthInfo(); ok {
						sfile.CommonAgreenmentSecure(args[1], auth, 843)
					} else {
						fmt.Println(sfile.AUTHGETWARN)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
