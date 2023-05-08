package cmd

import (
  "fmt"
  "strings"

  "github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
  Use:   "add",
  Short: "add a task to the list of tasks",
  Run: func(cmd *cobra.Command, args []string) {
      task := strings.Join(args, " ")
      fmt.Printf("Added \"%s\" to your task list.\n", task)
    },
}

func init() {
  RootCmd.AddCommand(addCmd)
}
