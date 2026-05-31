/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var filepath string = "tasktracker.json"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Allows you to maintain a task tracker in json.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("Hello! Your app is up and running. How can I help?")
	},
}

var addTaskCmd = &cobra.Command{
	Use: "add",
	Short: "Adds tasks to the file.",
	Long: "Adds tasks to the file.",
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		result, ok := addTaskToFile(filepath, args)
		if ok != nil {
			fmt.Println("Error: ", ok)
		} else {
			fmt.Println("Result: ", result)
		}
	},
}

var deleteTaskCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete tasks from the file based on id.",
	Long: "Delete tasks from the file based on id.",
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		result, ok := deleteTaskFromFile(filepath, args)
		if ok != nil {
			fmt.Println("Error: ", ok)
		} else {
			fmt.Println("Result: ", result)
		}
	},
}

var listAllTasksCmd = &cobra.Command{
	Use: "list [status]",
	Short: "Lists all tasks from the file.",
	Long: "Lists all tasks from the file.",
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			args = append(args, "all")
		}
		if args[0] != "all" && args[0] != "done" && args[0] != "in-progress" && args[0] != "todo" {
			fmt.Println("Error: Invalid status. Please use 'all', 'done', 'in-progress', or 'todo'.")
			return
		}

		_, ok := listAllTasksFromFile(filepath, args)
		if ok != nil {
			fmt.Println("Error: ", ok)
		}
	},
}

var updateTaskDescriptionCmd = &cobra.Command{
	Use: "update",
	Short: "Update tasks from the file based on id.",
	Long: "Update tasks from the file based on id.",
	Args: cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		result, ok := updateTaskDescription(filepath, args)
		if ok != nil {
			fmt.Println("Error:", ok)
		} else {
			fmt.Println("Result:", result)
		}
	},
}

var updateStatusDoneCmd = &cobra.Command{
	Use: "updateDone",
	Short: "Update task status from the file based on id.",
	Long: "Update task status from the file based on id.",
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		args = append(args, "done")
		result, ok := updateTaskStatus(filepath, args)
		if ok != nil {
			fmt.Println("Error:", ok)
		} else {
			fmt.Println("Result:", result)
		}
	},
}

var updateStatusInProgressCmd = &cobra.Command{
	Use: "updateInProgress",
	Short: "Update task status from the file based on id.",
	Long: "Update task status from the file based on id.",
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		args = append(args, "in-progress")
		result, ok := updateTaskStatus(filepath, args)
		if ok != nil {
			fmt.Println("Error:", ok)
		} else {
			fmt.Println("Result:", result)
		}
	},
}

var updateStatusToDoCmd = &cobra.Command{
	Use: "updateToDo",
	Short: "Update task status from the file based on id.",
	Long: "Update task status from the file based on id.",
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		args = append(args, "todo")
		result, ok := updateTaskStatus(filepath, args)
		if ok != nil {
			fmt.Println("Error:", ok)
		} else {
			fmt.Println("Result:", result)
		}
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task_tracker_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
	rootCmd.AddCommand(listAllTasksCmd)
	rootCmd.AddCommand(updateTaskDescriptionCmd)
	rootCmd.AddCommand(updateStatusDoneCmd)
	rootCmd.AddCommand(updateStatusInProgressCmd)
	rootCmd.AddCommand(updateStatusToDoCmd)
}


