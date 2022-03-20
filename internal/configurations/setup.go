//Package configurations provides functions to load the parameters to start the app
package configurations

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// Application contains data related to application configuration parameters.
type Application struct {
	Port string `yaml:"port" env:"PORT" env-default:"8080"`
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
}

// Load load application configuration
func Load() (Application, error) {
	cfg := Application{}
	if err := cleanenv.ReadConfig("../../config/config.yml", &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
