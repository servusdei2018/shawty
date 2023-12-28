package conf

import (
	"os"

	envparse "github.com/hashicorp/go-envparse"
)

type Config struct {
	Auth string
	Port string
}

// Parse parses .env into config.
func Parse() *Config {
	conf := Config{
		Auth: os.Getenv("SHAWTYAUTH"),
		Port: os.Getenv("SHAWTYPORT"),
	}

	if _, err := os.Stat(".env"); err == nil {
		f, err := os.Open(".env")
		if err != nil {
			return &conf
		}
		defer f.Close()
		data, err := envparse.Parse(f)
		if err != nil {
			return &conf
		}
		if data["SHAWTYAUTH"] != "" {
			conf.Auth = data["SHAWTYAUTH"]
		}
		if data["SHAWTYPORT"] != "" {
			conf.Port = data["SHAWTYPORT"]
		}
	}

	return &conf
}
