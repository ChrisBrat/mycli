/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"mycli/common"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a remote file to a local directory",
	Long:  `Copy a remote file to a local directory.`,
	Run: func(cmd *cobra.Command, args []string) {

		viperMethod := viper.GetString("copy.method")

		switch viperMethod {
		case "mount":
			mountCopyFile(args)
		case "git":
			gitCopyFile(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	copyCmd.Flags().StringP("method", "m", "mount", "Method to be used for copy")
	viper.BindPFlag("copy.method", copyCmd.Flags().Lookup("method"))
	//copyCmd.MarkFlagRequired("method")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gitCopyFile(args []string) {
	gitCopyViper := viper.Sub("actions.copy.git")
	if gitCopyViper == nil {
		panic("gitCopyViper not found")
	}

	remoteDirectory := gitCopyViper.GetString("remote.directory")
	fmt.Printf("Git copy %v", remoteDirectory)
}

func mountCopyFile(args []string) {
	mountCopyViper := viper.Sub("actions.copy.mount")
	if mountCopyViper == nil {
		panic("mountCopyViper not found")
	}

	remoteDirectory := mountCopyViper.GetString("remote.directory")
	filename := "file1.txt"

	fmt.Printf("COPY %v/%v to %v - ", remoteDirectory, filename, common.InstallsDirectory)
	err := os.MkdirAll(common.InstallsDirectory, 0755)
	check(err)

	from, err := os.Open(remoteDirectory + "/" + filename)
	check(err)
	defer from.Close()

	to, err := os.Create(common.InstallsDirectory + "/" + filename)
	check(err)
	defer to.Close()

	_, err = io.Copy(to, from)
	check(err)
	fmt.Printf("DONE\n")
}
