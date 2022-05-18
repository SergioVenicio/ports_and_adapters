/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/sergio/go-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productID string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "product CLI",
	Long:  `Create, enable, disable and get products`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(productService, action, productID, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "", "action")
	cliCmd.Flags().StringVarP(&productID, "id", "i", "", "id")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "price")
}
