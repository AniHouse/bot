package config

type Log struct {
	File      string `yaml:"file,omitempty"`
	Level     string `yaml:"level,omitempty"`
	Formatter string `yaml:"formatter,omitempty"`
}
