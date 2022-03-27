package egsql

import (
	"os"
	"path/filepath"
)

type config struct {
	homeDir string
}

const (
	// EnvVarHome is environment variable that represents the egsql home directory.
	EnvVarHome = "EGSQL_HOME"
)

func (c *config) home() error {
	home, ok := os.LookupEnv(EnvVarHome)
	if !ok {
		home, err := os.UserHomeDir()
		if err != nil {
			return ErrNoEgSQLHomeDir
		}
		home = filepath.Join(home, ".egsql")
	}
	c.homeDir = home
	return nil
}
