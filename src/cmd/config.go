/*
Copyright © 2023 Oswaldo Cho oswaldohome007@gmail.com
*/
package cmd

import (
	"fmt"
	sfile "sfile/sfile_body"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "modify config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			sfile.Error()
			fmt.Println("learn about sfile command from here=>https://brotherhoodhk.org/codelabcn/sfile/tutorial")
			return
		}
		sfile.ConfigureSfile(args)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
