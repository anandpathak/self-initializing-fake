package cmd

import "os"

type config struct {
	setupServerPort string
	fakeServerPort  string
}

func initConfig() config {
	return config{
		setupServerPort: getEnvironmentKeyWithDefault("SETUP_SERVER_PORT", "8112"),
		fakeServerPort:  getEnvironmentKeyWithDefault("FAKE_SERVER_PORT","8113"),
	}

}

func getEnvironmentKeyWithDefault(key, defaultValue string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return defaultValue
}
