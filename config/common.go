package config

import (
	"bytes"
	"io"
	"sync"
	"userservice/model"

	"github.com/pkg/errors"
)

type commonStore struct {
	configLock             sync.RWMutex
	config                 *model.Config
	configWithoutOverrides *model.Config
	enviromentOverrides    map[string]interface{}
}

func (cs *commonStore) Get() *model.Config {
	cs.configLock.RLock()
	defer cs.configLock.RUnlock()

	return cs.config
}

// func (cs *commonStore) GetEnvironmentOverrides() map[string]interface {

// }

func (cs *commonStore) Set() {

}

func (cs *commonStore) load(f io.ReadCloser) error {
	f2 := new(bytes.Buffer)
	tee := io.TeeReader(f, f2)

	loadedConfig, _, err := unmarsharlConfig(tee, true)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config with env overrides")
	}

	cs.config = loadedConfig

	return nil
}
