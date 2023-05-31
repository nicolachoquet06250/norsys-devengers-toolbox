package files

import (
	"fmt"
	"norsys-devengers-toolbox/arrays"
	"norsys-devengers-toolbox/environement"
	"norsys-devengers-toolbox/errors"
	"os"
	"strings"
)

type (
	FileSystemDirectory interface {
		Exists() (exists bool, err error)
		Create() error
		Remove() error
		Is() bool
	}

	Dir struct {
		Path string
	}
)

func (d Dir) Exists() (exists bool, err error) {
	d.Path = strings.ReplaceAll(d.Path, environement.Slash()+environement.Slash(), environement.Slash())

	if d.Is() {
		_, err := os.ReadDir(d.Path)
		errors.MaybeError(err, func(err error) *bool {
			r := err == nil
			return &r
		})
		if err != nil {
			return false, nil
		}
		return true, nil
	}

	return false, fmt.Errorf("le répertoire demandé n'existe pas")
}

func (d Dir) Create() error {
	d.Path = strings.ReplaceAll(d.Path, environement.Slash()+environement.Slash(), environement.Slash())

	var splitPath = strings.Split(d.Path, environement.Slash())

	_, _ = arrays.ArrayPop(&splitPath)

	p := strings.Join(splitPath, environement.Slash())
	err := os.MkdirAll(p, 0777)
	if err != nil {
		return fmt.Errorf("%s directory create error", err.Error())
	}

	return nil
}

func (d Dir) Remove() error {
	d.Path = strings.ReplaceAll(d.Path, environement.Slash()+environement.Slash(), environement.Slash())

	err := os.RemoveAll(d.Path)

	return err
}

func (d Dir) Is() bool {
	if !strings.Contains(d.Path, ".") {
		return true
	}
	return false
}

func NewDir(path string) Dir {
	return Dir{
		Path: path,
	}
}
