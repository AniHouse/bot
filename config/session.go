package config

type session struct {
	Log   *Log   `yaml:"log,omitempty"`
	Token string `yaml:"-"`
}
