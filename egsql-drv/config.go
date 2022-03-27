package egsql

import "os"

type config struct{}

const (
	// EnvVarHome is environment variable that represents the egsql home directory.
	EnvVarHome = "EGSQL_HOME"
)

func (c *config) home() {
	home, ok := os.LookupEnv(EnvVarHome)
	if !ok {
		// default
		home = ".bogo/"
		if _, err := os.Stat(home); os.IsNotExist(err) {
			err := os.Mkdir(home, 0777)
			if err != nil {
				panic(err)
			}
		}
	}
}
