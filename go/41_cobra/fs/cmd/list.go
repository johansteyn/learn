package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var longlisting bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Short description of the 'list' command...",
	Long: `Long description of the 'list' command 
that spans multiple lines...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("*** list anonymous Run function:")
		fmt.Printf("    --longlisting: %t\n", longlisting)
		fmt.Print("    args: ")
		fmt.Println(args)
	},
}

func init() {
	fmt.Println("*** list.init")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// JS: Added a local flag
	listCmd.Flags().BoolVarP(&longlisting, "long", "l", false, "Long listing")


}
