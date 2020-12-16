package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var cfg config

var (
	// absolute path to te config directory
	Path string
)

// Config fields
var (
	Session = &cfg.Session
	Bot     = &cfg.Bot
	DB      = &cfg.DB
	Modules = &cfg.Modules
)

type config struct {
	Session session           `yaml:"session,omitempty"`
	Bot     bot               `yaml:"bot,omitempty"`
	DB      db                `yaml:"database,omitempty"`
	Modules map[string]module `yaml:"modules,omitempty"`
}

func Load(path string) {
	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	Path = filepath.Dir(path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	Session.Token = os.Getenv("TOKEN")

	fmt.Println("Config loaded:", path)
}
