package config

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"log"
)

// the base of the struct that the data goes into
type Config struct {
	Mode string   `hcl:"Mode"`
	Port string   `hcl:"Port"`
	DB   DBConfig `hcl:"DB,block"`
}

type DBConfig struct {
	Host     string `hcl:"Host"`
	Port     string `hcl:"Port"`
	User     string `hcl:"User"`
	Password string `hcl:"Password"`
	DBName   string `hcl:"DBName"`
	SSLMode  string `hcl:"SSLMode"`
}

// a global variable to be able to access anywhere
var CONFIG Config

// Load in from the config
func Load() {
	err := hclsimple.DecodeFile("config.hcl", nil, &CONFIG)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	// if its in dev mode
	// print out entire config
	if CONFIG.Mode == "dev" {
		log.Printf("CONFIG: %#v", CONFIG)
	}
}
