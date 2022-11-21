package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func catify(filePath string) {
	// check if the file  provided in the path is exist
	if _, err := os.Stat(filePath); err == nil {
		// open the file to disply it content
		file, error := os.ReadFile(filePath)

		if error != nil {
            // display the error
            log.Fatal(error)
		}
        fmt.Println(string(file))

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
        log.Fatalln(err)
	} else {

	}
}

var catCmd = &cobra.Command{
    Use:   "cat is a standard Unix utility that reads files sequentially, writing them to standard output.",
	Short: "",
    Long:  `cat is a standard Unix utility that reads files sequentially, writing them to standard output. The name is derived from its function to (con)catenate files (from Latin catenare, "to chain"). It has been ported to a number of operating systems.`,
	Run: func(cmd *cobra.Command, args []string) {
        if len(args)> 0 {
            catify(args[1])
        }
	},
}

func Execute() {
	if err := catCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
