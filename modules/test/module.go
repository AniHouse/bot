package test

import (
	"github.com/anihouse/bot/app"
	"github.com/sirupsen/logrus"
)

type module struct {
	app     *app.Module
	enabled bool
}

func (module) ID() string {
	return "test"
}

func (m module) IsEnabled() bool {
	return m.enabled
}

func (module) LoadConfig(path string) error {
	return nil
}

func (module) SetLogger(logger *logrus.Logger) {

}

func (m *module) Init(prefix string) error {
	m.app = app.NewModule(_module.ID(), prefix)

	m.app.On("ping").Handle(m.onPing)
	
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
