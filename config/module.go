package config

type module struct {
	Enabled    bool    `yaml:"enabled,omitempty"`
	Prefix     string  `yaml:"prefix,omitempty"`
	ConfigFile *string `yaml:"config_file,omitempty"`
	Log        *Log    `yaml:"log,omitempty"`
}
