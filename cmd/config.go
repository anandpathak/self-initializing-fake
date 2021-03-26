package cmd

import (
	"os"
	"strconv"
)

type config struct {
	setupServerPort string
	fakeServerPort  string
	fakeServerTimeout int
}

func initConfig() config {
	timeout, _ := strconv.Atoi(getEnvironmentKeyWithDefault("FAKE_SERVER_TIMEOUT","30"))
	return config{
		setupServerPort: getEnvironmentKeyWithDefault("SETUP_SERVER_PORT", "8112"),
		fakeServerPort:  getEnvironmentKeyWithDefault("FAKE_SERVER_PORT","8113"),
		fakeServerTimeout: timeout,
	}

}

func getEnvironmentKeyWithDefault(key, defaultValue string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return defaultValue
}
