package cmd

import (
	"sophon-cli/auth"
	"sophon-cli/term"

	"github.com/spf13/cobra"
)

var pin string

var signInCmd = &cobra.Command{
	Use:   "sign-in",
	Short: "Sign in to a Sophon server",
	Args:  cobra.NoArgs,
	Run:   signIn,
}

func init() {
	RootCmd.AddCommand(signInCmd)

	signInCmd.Flags().StringVar(&pin, "pin", "", "Sign in with a pin from the Sophon web UI")
}

func signIn(cmd *cobra.Command, args []string) {
	if pin != "" {
		err := auth.SignInWithCode(pin, "")

		if err != nil {
			term.OutputErrorAndExit("Error signing in: %v", err)
		}

		return
	}

	err := auth.SelectOrSignInOrCreate()

	if err != nil {
		term.OutputErrorAndExit("Error signing in: %v", err)
	}
}
