package main

import (
	"os"

	_ "github.com/njuptlzf/servercheck/pkg/check"
	"github.com/njuptlzf/servercheck/pkg/inspector"
	"github.com/njuptlzf/servercheck/pkg/option"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "servercheck",
	Short: "A tool to check server environment",
	Run: func(cmd *cobra.Command, args []string) {
		inspector := inspector.NewInspector()
		// validate
		if err := inspector.Validate(); err != nil {
			panic(err)
		}
		// check
		if err := inspector.Check(); err != nil {
			panic(err)
		}
		// result
		if err := inspector.PrintResult(); err != nil {
			panic(err)
		}
		// fail: return 1; ohters: return 0
		if !inspector.ZeroRc() {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&option.Opt.Strict, "strict", false, "when set to true, the result is only pass return code is passed.")
	rootCmd.PersistentFlags().BoolVar(&option.Opt.CPUCore, "cpu-core", true, "check CPU core")
	rootCmd.PersistentFlags().IntVar(&option.Opt.CPUCoreNum, "cpu-core-num", 4, "expect CPU core number")
	rootCmd.PersistentFlags().BoolVar(&option.Opt.CPUArch, "cpu-arch", true, "check CPU arch")
	rootCmd.PersistentFlags().BoolVar(&option.Opt.DiskAvail, "disk-avail", true, "check disk avail")
	rootCmd.PersistentFlags().StringSliceVar(&option.Opt.DiskForDir, "disk-for-dir", []string{"/;100G;>= 100GiB"},
		"check disk avail for dir. Use the first two semicolons(;) to separate the three elements: directory path; expected size(Default unit (iB)); failure suggestions.\n"+
			"Such as --disk-for-dir=\"/var;150G;greater than or equal to 150GiB\" --disk-for-dir=\"/;100G;>= 100GiB\"")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
