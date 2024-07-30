package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github/1in9/wjvs/internal"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "list",
	Short: "It will list the Java Development Kit versions installed on the current device.",
	Long: `It will list the Java Development Kit versions installed on the current device. For example:

> wjvm list

    1.8.0_281
    11.0.19
`,
	Run: func(cmd *cobra.Command, args []string) {
		currentUseVersion := internal.CurrentUseJdkVersion()
		fmt.Println()
		jdkInstallInfos := internal.InstalledJDKInfo()
		for _, info := range jdkInstallInfos {
			if currentUseVersion == info.Version {
				using := color.New(color.FgHiGreen).Sprintf("* %s (Currently using)", info.Version)
				fmt.Println(fmt.Sprintf(`  %s`, using))
			} else {
				fmt.Println(fmt.Sprintf(`    %s`, info.Version))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
