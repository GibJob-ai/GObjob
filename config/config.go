package config

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"log"
)

// the base of the struct that the data goes into
type Config struct {
	Mode string `hcl:"Mode"`
}

// a global variable to be able to access anywhere
var CONFIG Config

// Load in from the config
func Load() {
	err := hclsimple.DecodeFile("config.hcl", nil, &CONFIG)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
}
