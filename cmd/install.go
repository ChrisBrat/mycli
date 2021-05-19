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
	"io/ioutil"
	"log"
	"mycli/common"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install local files",
	Long:  `Install local files`,
	Run: func(cmd *cobra.Command, args []string) {

		viperMethod := viper.GetString("install.action")

		switch viperMethod {
		case "list":
			listFiles(args)
		case "execute":
			executeFile(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringP("action", "a", "list", "Action to be performed")
	viper.BindPFlag("install.action", installCmd.Flags().Lookup("action"))
}

func listFiles(args []string) {

	files, err := ioutil.ReadDir(common.InstallsDirectory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available installs : ")
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".zip" {
			fmt.Printf("- %v\n", file.Name())
		}
	}
}

func executeFile(args []string) {

}
