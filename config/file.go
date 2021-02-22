package config

import (
	"io"
	"os"
	osPath "path"
	"userservice/model"

	"github.com/pkg/errors"
)

type FileStore struct {
	commonStore

	path string
}

func NewFileStore(path string) (fs *FileStore, err error) {
	pwd, _ := os.Getwd()

	fs = &FileStore{
		path: osPath.Join(pwd, "config", path),
	}

	if err = fs.Load(); err != nil {
		return nil, errors.Wrap(err, "failed to load")
	}

	return fs, nil
}

func (fs *FileStore) Load() (err error) {
	var f io.ReadCloser
	f, err = os.Open(fs.path)

	if os.IsNotExist(err) {
		if err != nil {
			return errors.Wrap(err, "failed to serialize default")
		}
	} else if err != nil {
		return errors.Wrapf(err, "failed to open %s for reading", fs.path)
	}
	defer func() {
		closeErr := f.Close()
		if err == nil && closeErr != nil {
			err = errors.Wrap(closeErr, "failed to close")
		}
	}()

	return fs.commonStore.load(f)
}

func (fs *FileStore) Set(newConfig *model.Config) (*model.Config, error) {
	return newConfig, nil
}
