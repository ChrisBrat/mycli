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
	"mycli/common"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var documentationCmd = &cobra.Command{
	Use:   "documentation",
	Short: "Retrieve documentation from the knowledge base",
	Long:  `Retrieve documentation from the knowledge base.`,
	Run: func(cmd *cobra.Command, args []string) {
		viperMethod := viper.GetString("documentation.method")

		switch viperMethod {
		case "mount":
			mountDocumentation(args)
		case "git":
			gitDocumentation(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(documentationCmd)

	documentationCmd.Flags().StringP("method", "m", "git", "Method to be used for documentation copy")
	viper.BindPFlag("documentation.method", documentationCmd.Flags().Lookup("method"))
}

func gitDocumentation(args []string) {

	gitDocumentationViper := viper.Sub("actions.documentation.git")
	if gitDocumentationViper == nil {
		panic("gitDocumentationViper not found")
	}

	remoteDirectory := gitDocumentationViper.GetString("remote.directory")
	fmt.Printf("Clone git documentation %v to %v\n", remoteDirectory, common.DocumentationDirectory)

	common.GitClone(remoteDirectory, common.DocumentationDirectory)
}

func mountDocumentation(args []string) {
	/**
	* TODO: Implement mount documentation
	 */
}
