package config

type db struct {
	Driver            string `yaml:"driver,omitempty"`
	Connection        string `yaml:"connection,omitempty"`
	MaxIdleConnection int    `yaml:"max_idle_connection,omitempty"`
	MaxOpenConnection int    `yaml:"max_open_connection,omitempty"`
	Log               *Log   `yaml:"log,omitempty"`
}
