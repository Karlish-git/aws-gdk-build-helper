package confparser

import (
	"encoding/json"
	"log"
)

type Config struct {
	GdkVersion string `json:"gdk_version"`
	Component  struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Author  string `json:"author"`
		Build   struct {
			BuildSystem string `json:"build_system"`
		} `json:"build"`
		Publish struct {
			Bucket string `json:"bucket"`
			Region string `json:"region"`
		} `json:"publish"`
	}
}

func ParseCoonfig(bytes []byte) Config {

	var rawConfig struct {
		GdkVersion string                     `json:"gdk_version"`
		Component  map[string]json.RawMessage `json:"component"`
	}

	err := json.Unmarshal(bytes, &rawConfig)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	config.GdkVersion = rawConfig.GdkVersion

	if len(rawConfig.Component) != 1 {
		log.Fatal("There should be exactly one component in gdk-config.json")
	}

	for name, rawComponent := range rawConfig.Component {
		err := json.Unmarshal(rawComponent, &config.Component)
		if err != nil {
			log.Fatal(err)
		}
		config.Component.Name = name
	}

	return config
}
