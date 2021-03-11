package main

import (
	"log"

	"github.com/anonsachin/cli-spf13/pkg/config"
	"github.com/spf13/cobra"

	"github.com/anonsachin/cli-spf13/pkg/cmd"
)

func main() {
	// Getting the configuration
	conf, err := config.NewEmojiConfig()

	if err != nil{
		log.Panic(err.Error())
	}
	// Setting the commands
	root := cmd.SetupRoot()
	get := cmd.BuildGet(conf)
	list := cmd.BuildList(conf)
	// Linking them together
	cmd.WireUPSubCommands(root,get,list)
	// Running it
	cobra.CheckErr(root.Execute())
}
