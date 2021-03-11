package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/anonsachin/cli-spf13/pkg/config"
	"github.com/spf13/cobra"
)

type get struct{
	conf *config.Config
}
// BuildGet to create the get command
func BuildGet(conf *config.Config) *cobra.Command{
	g := &get{conf: conf}

	return &cobra.Command{
		Use: "get",
		Short: "Used to get a particular emoji from the list",
		RunE: func(cmd *cobra.Command, args []string) error{
			emojiName := args[0]
			emoji, err := g.GetEmoji(emojiName)
			if err != nil{
				return err
			}
			if !clipboard.Unsupported{
				err = clipboard.WriteAll(*emoji)
				}
				fmt.Println(*emoji)
			return nil
		},
		Args: cobra.ExactArgs(1),
	}
	
}

// GetEmoji prints the emoji and adds to clipboard
func (g *get) GetEmoji(name string) (*string, error){
	emoji, ok := g.conf.Emojis[name]

	if !ok {
		return nil, fmt.Errorf("Emoji not present try something else 😞")
	}
	return &emoji, nil
}