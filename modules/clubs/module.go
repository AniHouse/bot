package clubs

import (
	"os"

	"github.com/anihouse/bot/app"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	log  *logrus.Logger
	conf cfg
)

type module struct {
	app     *app.Module
	enabled bool
}

func (module) ID() string {
	return "clubs"
}

func (m module) IsEnabled() bool {
	return m.enabled
}

func (module) LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	err = yaml.NewDecoder(file).Decode(&conf)
	if err != nil {
		return err
	}

	return nil
}

func (module) SetLogger(logger *logrus.Logger) {
	log = logger
}

func (m *module) Init(prefix string) error {
	m.app = app.NewModule(_module.ID(), prefix)

	m.app.On("clubcreate").Handle(onClubCreeate)
	m.app.On("clubapply").Handle(onClubApply)

	return nil
}

func (m *module) Enable() {
	m.enabled = true
	m.app.Enable()
}

func (m *module) Disable() {
	m.enabled = false
	m.app.Disable()
}
