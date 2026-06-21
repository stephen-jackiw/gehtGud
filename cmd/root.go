/*
Copyright © 2026 | Stephen Jackiw | interviewaid@stephenjackiw.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "v2",
	Short: "A CLI/TUI tool for tracking your interviews and progress on technical interview skills",
	Long: `GehtGud is an interview aid tool for helping you track your journey while you get better at technical interviews.

It is intended to do a couple of things:
	- Track your attempts on Leetcode and System Design problems
	- Generate mock interviews
	- Provide spaced repetition of previous problems so you can maximize your learning potential

There are potential plans to expand functionality to:
	- Track resumes
	- Track applications
	- Track potential referrals
	- Habit Tracker
	- Day-by-day tracker so you know where your time is going on your job search
`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.v2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


