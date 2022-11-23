package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
	"os"
)

// define flags

type Options struct {
	flags  *pflag.FlagSet
	params []string
}

var numberOfLines string // -n define the lines of the Number all output lines.

func catify(options Options) {

	// loop through the params
	for _, filePath := range options.params {
		log.Print(filePath)
		if _, err := os.Stat(filePath); err == nil {
			file, error := os.ReadFile(filePath)
			// if error print error
			if error != nil {
				fmt.Print(error)
			}
			// if not error print the content of the file
			fmt.Println(string(file))

		} else if errors.Is(err, os.ErrNotExist) {
			fmt.Print(err)
			break // if error quit the request
		}
	}
}

var CatCmd = &cobra.Command{
	Use:   "cat is a standard Unix utility that reads files sequentially, writing them to standard output.",
	Short: "Concatenate FILE(s) to standard output.",

	Run: func(cmd *cobra.Command, args []string) {
		var flagsOption = cmd.Flags()
		var options = Options{
			flags:  flagsOption,
			params: args,
		}
		fmt.Print(options.flags.GetString("n"))
		catify(options)
	},
}

// define the cat cmd version with

func DefineCatFlags(command *cobra.Command) {
	command.Flags().StringVarP(&numberOfLines, "number", "n", "", "number all output lines")
}
