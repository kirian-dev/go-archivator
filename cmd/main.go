package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Simple archiator",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handlerError((err))
	}
}

func handlerError(err error) {
	_, err = fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
