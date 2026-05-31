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
	Long: `Allows you to maintain a task tracker in json.
Commands that you can use:
add - Adds tasks to the file.
delete - Delete tasks from the file based on id.
list - Lists tasks from the file.
update - Update the description of the task from the file based on id.
updateDone - Update task status to done for a particular task from the file based on id.
updateInProgress - Update task status to in-progress for a particular task from the file based on id.
updateToDo - Update task status to todo for a particular task from the file based on id.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("Hello! Your app is up and running. How can I help?")
	},
}

var addTaskCmd = &cobra.Command{
	Use: "add",
	Short: "Adds tasks to the file.",
	Long: `Adds tasks to the file. 
Usage: task-cli add "task description"`,
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
	Long: `Delete tasks from the file based on id. 
Usage: task-cli delete <id>`,
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

var listTasksCmd = &cobra.Command{
	Use: "list [status]",
	Short: "Lists tasks from the file.",
	Long: `Lists tasks from the file.  
Usage: task-cli list [status]
Pass nothing or all to list all tasks.
Pass done, in-progress, or todo to list tasks with that status.`,
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
	Short: "Update the description of the task from the file based on id.",
	Long: `Update the description of the task from the file based on id. 
Usage: task-cli update <id> "new description"`,
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
	Short: "Updates task status to done for a particular task from the file based on id.",
	Long: `Updates task status to done for a particular task from the file based on id. 
Usage: task-cli updateDone <id>`,
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
	Short: "Updates task status to in-progress for a particular task from the file based on id.",
	Long: `Updates task status to in-progress for a particular task from the file based on id. 
Usage: task-cli updateInProgress <id>`,
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
	Short: "Updates task status to todo for a particular task from the file based on id.",
	Long: `Updates task status to todo for a particular task from the file based on id. 
Usage: task-cli updateToDo <id>`,
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
	rootCmd.AddCommand(listTasksCmd)
	rootCmd.AddCommand(updateTaskDescriptionCmd)
	rootCmd.AddCommand(updateStatusDoneCmd)
	rootCmd.AddCommand(updateStatusInProgressCmd)
	rootCmd.AddCommand(updateStatusToDoCmd)
}


