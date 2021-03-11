/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"

	"github.com/cli-spf13/pkg/config"
	"github.com/spf13/cobra"

	"github.com/cli-spf13/pkg/cmd"
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
