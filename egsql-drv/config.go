package egsql

import (
	"os"
	"path/filepath"

	"github.com/nao1215/egsql/misc/errfmt"
	"github.com/nao1215/egsql/misc/file"
)

type config struct {
	homeDir string
}

const (
	// EnvVarHome is environment variable that represents
	// the egsql home directory.
	EnvVarHome = "EGSQL_HOME"
)

// setHomeDirPath sets the egsql home directory path
// to the field in the config struct.
func (c *config) setHomeDirPath() error {
	home, ok := os.LookupEnv(EnvVarHome)
	if !ok {
		home, err := os.UserHomeDir()
		if err != nil {
			return errfmt.Wrap(ErrNotGetEgSQLHomeDir, err.Error())
		}
		home = filepath.Join(home, ".egsql")
	}
	c.homeDir = home
	return nil
}

// 	CreateHomeDirIfNeeded creates the egsql home directory if needed.
func (c *config) createHomeDirIfNeeded() error {
	if c.homeDir == "" {
		if err := c.setHomeDirPath(); err != nil {
			return err
		}
	}
	if !file.IsDir(c.homeDir) {
		err := os.Mkdir(c.homeDir, 0644)
		if err != nil {
			return errfmt.Wrap(ErrNotCreateEgSQLHomeDir, err.Error())
		}
	}
	return nil
}
