package cmd

import (
	"fmt"
	"os"
	"sophon-cli/term"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var helpShowAll bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: `sophon [command] [flags]`,
	// Short: "Sophon: iterative development with AI",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// if no arguments were passed, start the repl
	if len(os.Args) == 1 ||
		(len(os.Args) == 2 && strings.HasPrefix(os.Args[1], "--") && os.Args[1] != "--help") ||
		(len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--") && os.Args[1] != "--help" && strings.HasPrefix(os.Args[2], "--") && os.Args[2] != "--help") {

		// Instead of directly calling replCmd.Run, parse the flags first
		replCmd.ParseFlags(os.Args[1:])
		replCmd.Run(replCmd, []string{})
		return
	}

	if err := RootCmd.Execute(); err != nil {
		// term.OutputErrorAndExit("Error executing root command: %v", err)
		// log.Fatalf("Error executing root command: %v", err)

		// output the error message to stderr
		term.OutputSimpleError("Error: %v", err)

		fmt.Println()

		color.New(color.Bold, color.BgGreen, color.FgHiWhite).Println(" Usage ")
		color.New(color.Bold).Println("  sophon [command] [flags]")
		color.New(color.Bold).Println("  sdx [command] [flags]")
		fmt.Println()

		color.New(color.Bold, color.BgGreen, color.FgHiWhite).Println(" Help ")
		color.New(color.Bold).Println("  sophon help # show basic usage")
		color.New(color.Bold).Println("  sophon help --all # show all commands")
		color.New(color.Bold).Println("  sophon [command] --help")
		fmt.Println()

		color.New(color.Bold, color.BgGreen, color.FgHiWhite).Println(" Common Commands ")
		color.New(color.Bold).Println("  sophon new # create a new plan")
		color.New(color.Bold).Println("  sophon tell # tell the plan what to do")
		color.New(color.Bold).Println("  sophon continue # continue the current plan")
		color.New(color.Bold).Println("  sophon settings # show plan settings")
		color.New(color.Bold).Println("  sophon set # update plan settings")
		fmt.Println()

		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {

}

func init() {
	var helpCmd = &cobra.Command{
		Use:     "help",
		Aliases: []string{"h"},
		Short:   "Display help for Sophon",
		Long:    `Display help for Sophon.`,
		Run: func(cmd *cobra.Command, args []string) {
			term.PrintCustomHelp(helpShowAll)
		},
	}

	RootCmd.AddCommand(helpCmd)
	RootCmd.AddCommand(connectClaudeCmd)
	RootCmd.AddCommand(disconnectClaudeCmd)

	// add an --all/-a flag
	helpCmd.Flags().BoolVarP(&helpShowAll, "all", "a", false, "Show all commands")
}
