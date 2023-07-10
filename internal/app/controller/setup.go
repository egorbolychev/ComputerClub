package controller

import (
	"github.com/egorbolychev/internal/app/store"
)

// Configure Store and start the application
func Start(config *Config, taskStr []string) error {
	st := store.New()
	if err := st.ConfigureStore(taskStr, config.MaxTables); err != nil {
		return err
	}

	c := NewController(config, st)
	c.ReadAndServe()
	return nil
}
