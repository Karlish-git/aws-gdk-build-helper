package build

import (
	"strings"

	confparser "github.com/Karlish-git/aws-gdk-build-helper/internal/conf_parser"
)

// FillRecipe fills the recipe.yaml with the data from gdk-config.`json`
// TO escape, use \{COMPONENT_NAME\}
// it replaces fallowing items:
// {COMPONENT_NAME} -> name from gdk-config.`json`
// {COMPONENT_VERSION} -> version from gdk-config.`json`
// {COMPONENT_AUTHOR} -> author from gdk-config.`json`
// {BUCKET} -> bucket from gdk-config.`json`
// {REGION} -> region from gdk-config.`json`
func FillRecipe(bytes []byte, gdkconf confparser.Config) string {

	recipe := string(bytes)

	recipe = strings.ReplaceAll(recipe, "{COMPONENT_NAME}", gdkconf.Component.Name)
	recipe = strings.ReplaceAll(recipe, "{COMPONENT_VERSION}", gdkconf.Component.Version)
	recipe = strings.ReplaceAll(recipe, "{COMPONENT_AUTHOR}", gdkconf.Component.Author)
	recipe = strings.ReplaceAll(recipe, "{BUCKET}", gdkconf.Component.Publish.Bucket)
	recipe = strings.ReplaceAll(recipe, "{REGION}", gdkconf.Component.Publish.Region)

	return recipe
}
