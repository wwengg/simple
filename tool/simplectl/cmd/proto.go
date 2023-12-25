/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// protoCmd represents the proto command
var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "proto",
	Long:  `proto`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proto called")
	},
}

var protoNewCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "eg:simplectl proto new user",
	Long:  `generate *.proto file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proto called")
	},
}

func init() {
	rootCmd.AddCommand(protoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// protoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// protoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
