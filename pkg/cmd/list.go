package cmd

import (
	"fmt"

	"github.com/cli-spf13/pkg/config"
	"github.com/spf13/cobra"
)

type list struct{
	conf *config.Config
}
// BuildList to create the get command
func BuildList(conf *config.Config) *cobra.Command{
	l := &list{conf: conf}

	return &cobra.Command{
		Use: "list",
		Short: "Lists all the available emojis",
		Run: func(cmd *cobra.Command, args []string) {
			l.ListALL()
		},
	}
	
}

// ListALL prints all emojis
func (l *list) ListALL(){
	fmt.Println("==============🎊THE AVAILABLE EMOJIS🎊============")
	// Printing all of the emojis
	for k,v := range l.conf.Emojis{
		fmt.Printf("==== The name %s => %s \n",k,v)
	}
}