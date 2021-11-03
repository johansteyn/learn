package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Short description of the 'cd' command...",
	Long: `Long description of the 'cd' command 
that spans multiple lines...`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("*** cd anonymous Run function:")
		fmt.Print("    args: ")
		fmt.Println(args)
    dir := args[0]
		fmt.Println("Changing directory to: " + dir)
	},
}

func init() {
	fmt.Println("*** cd.init")
	rootCmd.AddCommand(cdCmd)
}
