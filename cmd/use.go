package cmd

import (
	"github/1in9/wjvs/internal"
	"log"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Used to switch the current Java Development Kit version, must be the installed version.",
	Long: `Used to switch the current Java Development Kit version.
It must be the installed version number listed in the wjvm list instruction.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Please enter version parameters.")
		}
		jdkInstalledInfos := internal.InstalledJDKInfo()
		version := args[0]
		versionExist := false
		javaHomePath := ""
		for _, info := range jdkInstalledInfos {
			if version == info.Version {
				versionExist = true
				javaHomePath = info.JavaHome
			}
		}
		if versionExist {
			err := internal.SetEnv("JAVA_HOME", javaHomePath)
			if err != nil {
				log.Fatal("Failed to set JAVA_SOME system variable.Error: ", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
