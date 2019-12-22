package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Config is the shared config of commands.
type Config struct {
	OutputFolder string
}

// RootCmd is the object of root command settings.
type RootCmd struct {
	Config
	cmd *cobra.Command
}

// NewRootCmd creates a new `RootCmd` object.
func NewRootCmd(cfg Config) *RootCmd {
	ret := &RootCmd{
		Config: cfg,

		cmd: &cobra.Command{
			Use: "purr",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Usage()
			},
		},
	}

	snippetCmd := NewSnippetCmd(cfg)
	ret.cmd.AddCommand(snippetCmd.Cmd())
	return ret
}

// Execute function is the entry point of commands.
func (rc *RootCmd) Execute() {
	if err := rc.cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
