package cmd

import (
    "errors"
    "fmt"
	"github.com/spf13/cobra"
    "log"
    "os"
)

func catify(params []string) {

    // loop through the params
    for _, filePath := range params {
        log.Print(filePath)
        if _, err := os.Stat(filePath) ; err == nil{
            file, error := os.ReadFile(filePath)
            // if error print error
            if error != nil {
               fmt.Print(error)
            }
            // if not error print the content of the file
            fmt.Println(string(file))

        }else if errors.Is(err, os.ErrNotExist) {
            fmt.Print(err)
            break  // if error quit the request
        }
    }
}

var CatCmd = &cobra.Command{
    Use:   "cat is a standard Unix utility that reads files sequentially, writing them to standard output.",
	Short: "",
    Args: func(cmd *cobra.Command, args []string) error {
        if len(args) > 0 {
            return nil
        }
        return errors.New("requires at least one arg")
    },
    Long:  `cat is a standard Unix utility that reads files sequentially, writing them to standard output. The name is derived from its function to (con)catenate files (from Latin catenare, "to chain"). It has been ported to a number of operating systems.`,
	Run: func(cmd *cobra.Command, args []string) {
        catify(args)
	},
}
