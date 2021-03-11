package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SetupRoot is to get the root command redy
func SetupRoot() *cobra.Command{
	return &cobra.Command{
		Use:   "emoji",
		Short: "Gives you emojis when you need it.",
		Long: `Dispensary of emoji -

if you need to give someone the finger 
	emoji get finger - 🖕
if uou are delited with someone
	emoji get grin - 😁.`,
	}
}

// WireUPSubCommands adds sub commands to the root command
func WireUPSubCommands(root *cobra.Command,cmds ...*cobra.Command)error{
	if len(cmds) == 0 {
		return fmt.Errorf("There are no subcomands")
	}
	// Add all the commands to the root
	for _,sub := range cmds {
		root.AddCommand(sub)
	}
	return nil
}