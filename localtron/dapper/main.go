/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	dapper "github.com/singulatron/singulatron/localtron/dapper/app"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "dapper"}
	var folder string
	var anon bool

	params, remainingArgs := extractParams(os.Args)
	os.Args = remainingArgs

	var runCmd = &cobra.Command{
		Use:   "run [config file]",
		Short: "Run app",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appFilePath := args[0]
			run(appFilePath, params, anon)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&folder, "folder", "f", ".", "directory containing configuration files")

	runCmd.Flags().BoolVar(&anon, "anon", false, "Run in anonymous mode")

	rootCmd.AddCommand(runCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func extractParams(args []string) (map[string]string, []string) {
	params := make(map[string]string)
	var newArgs []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--var-") {
			split := strings.SplitN(arg[2:], "=", 2)
			if len(split) == 2 {
				key := strings.TrimPrefix(split[0], "var-")
				value := split[1]
				params[key] = value
			}
		} else {
			newArgs = append(newArgs, arg)
		}
	}
	return params, newArgs
}

func run(appFilePath string, params map[string]string, anon bool) {
	cm := dapper.NewConfigurationManagerFromSource()
	app, err := cm.LoadAppConfiguration(appFilePath)
	if err != nil {
		log.Fatalf("Failed to load app file: %v", err)
	}

	fmt.Println("Parameters:")
	if len(params) == 0 {
		fmt.Println("   None")
	}
	for key, value := range params {
		fmt.Printf("   %v=%v\n", key, value)
	}

	cont, err := cm.Run(app, params, anon)
	if err != nil {
		fmt.Printf("Failed to resolve feature dependencies: %v\n", err)
		if cont != nil && cont.RebootRequired {
			fmt.Printf("A restart is required to fix this!")
		}
		os.Exit(1)
	}

}
