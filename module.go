package bot

import (
	"fmt"
	"path/filepath"

	"github.com/anihouse/bot/config"
	"github.com/sirupsen/logrus"
)

type Module interface {
	Init(string) error
	LoadConfig(string) error
	SetLogger(*logrus.Logger)

	ID() string

	IsEnabled() bool
	Enable()
	Disable()
}

func (m modules) Register(name string, module Module) {
	m[name] = module
}

func (m modules) Init() {
	fmt.Println("\nInit modules...")
	for id, cfg := range *config.Modules {
		fmt.Printf("%-10s", id+":")

		mod, exists := m[id]
		if !exists {
			fmt.Println("Not found.")
			continue
		}

		if cfg.Log != nil {
			l, err := Logger(cfg.Log)
			if err != nil {
				fmt.Println(err)
				continue
			}
			mod.SetLogger(l)
		}

		if cfg.ConfigFile != nil {
			configFile := *cfg.ConfigFile
			if !filepath.IsAbs(configFile) {
				configFile = filepath.Join(config.Path, *cfg.ConfigFile)
			}

			if err := mod.LoadConfig(configFile); err != nil {
				fmt.Println(err)
				continue
			}
		}

		if err := mod.Init(cfg.Prefix); err != nil {
			fmt.Println(err)
			continue
		}

		if cfg.Enabled {
			mod.Enable()
		} else {
			mod.Disable()
		}

		status := "ENABLED"
		if !mod.IsEnabled() {
			status = "DISABLED"
		}
		fmt.Println(status)
	}
}
