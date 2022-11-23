package main

import (
    "github.com/phareal/unix-utilities/cmd"
    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{Use: "app"}
    rootCmd.AddCommand(cmd.CatCmd)
    cmd.DefineCatFlags(cmd.CatCmd)
    rootCmd.Execute()
}