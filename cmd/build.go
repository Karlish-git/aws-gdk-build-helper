/*
Copyright © 2024 Karlish

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"log"

	"os"

	"github.com/Karlish-git/aws-gdk-build-helper/internal/build"
	confparser "github.com/Karlish-git/aws-gdk-build-helper/internal/conf_parser"
	"github.com/spf13/cobra"
)

var buildSimple bool
var buildZip bool

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Help you build GGv2 component reading gdk-config.json and generating the recipe and artifacts",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: buildProject,
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().BoolVarP(&buildSimple, "simple", "s", false, "Build the project without zipping")
	buildCmd.Flags().BoolVarP(&buildZip, "zip", "z", false, "Build the project and zip the artifacts")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func buildProject(cmd *cobra.Command, args []string) {
	if buildSimple && buildZip {
		log.Fatal("You can't use both --simple and --zip flags")
	}
	// if there isn't recipie.yaml and gdk-config.json in cwd, then exit
	// if there is recipie.yaml and gdk-config.json in cwd, then Execute
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	recipePath := wd + "/recipe.yaml"
	gdkConfigPath := wd + "/gdk-config.json"

	byteRecipie, err := os.ReadFile(recipePath)
	if err != nil {
		log.Fatal("Error reading recipe.yaml: ", err)
	}
	byteConfig, err := os.ReadFile(gdkConfigPath)
	if err != nil {
		log.Fatal("Error reading gdk-config.json: ", err)
	}

	conf := confparser.ParseCoonfig(byteConfig)
	recipe := build.FillRecipe(byteRecipie, conf)

	print(recipe)

}
