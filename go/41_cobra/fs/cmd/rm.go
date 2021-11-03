package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var force bool

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Short description of the 'rm' command...",
	Long: `Long description of the 'rm' command 
that spans multiple lines...`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("*** rm anonymous Run function:")
		fmt.Print("    args: ")
		fmt.Println(args)
    file := args[0]
		fmt.Printf("Removing: %s (force=%t)\n", file, force)
	},
}

func init() {
	fmt.Println("*** rm.init")
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal (required)")
	rmCmd.MarkFlagRequired("force")
}
