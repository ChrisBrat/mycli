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
		copyFile(args)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//copyCmd.PersistentFlags().StringP("remote.directory", "r", viper.GetString("remote.directory"), "Remote copy target url")
	//viper.BindPFlag("remote.directory", copyCmd.PersistentFlags().Lookup("remote.directory"))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func copyFile(args []string) {
	remoteDirectory := viper.GetString("remote.directory")
	filename := "file1.txt"

	fmt.Printf("COPY %v/%v to %v - ", remoteDirectory, filename, ".data/installs")
	err := os.MkdirAll(".data/installs", 0755)
	check(err)

	from, err := os.Open(remoteDirectory + "/" + filename)
	check(err)
	defer from.Close()

	to, err := os.Create(".data/installs/" + filename)
	check(err)
	defer to.Close()

	_, err = io.Copy(to, from)
	check(err)
	fmt.Printf("DONE\n")
}
