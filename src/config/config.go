package config

import (
	"errors"
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

// implements YAML config
// follows the example from https://github.com/koddr/example-go-config-yaml

type ConfigItems struct {
	common struct {
		buildVersion string `yaml:"buildVersion"` // will not implement legacy features such as 10540/10541 ports
	}
	sentry struct {
		sentryKey	string `yaml:"sentryKey"`
		sentryProject	string	`yaml:"sentryProject"`
		sentryEnvironment	string	`yaml:"sentryEnvironment"`
	}
}

type Values interface {
	getValue() string
	setValue(key string, value string) ()
}

func initConfigPath() error {
	// inits config path on /etc/OpenMOS/mosConfig.yaml

	folderPath := "/etc/datameshquantum"

	// look for the OpenMOS folder and create if missing

	if _, err := os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// look for the mosConfig.yaml and create empty if missing

	f1, err := os.Create(folderPath, "/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	f1.Close()

}

func initConfig(configPath string) (*Config, error) {
	config := &Config{}

	// open file
	file, err := os.Open(configPath)

	if err != nil {

		// if missing, attempt to recreate once (expected in first run)

		err := initConfigPath()

		file, err := os.Open(configPath)

		if err != nil {
			log.Fatal(err)
		}

		return nil, err

	}

	
	}

	// defer file.Close()

	// init YAML decode and start decoding

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

}
