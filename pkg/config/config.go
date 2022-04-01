package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cloudcloud/roadie/pkg/types"
	"go.uber.org/zap"
)

// Config is a structure to provide the necessary configuration for
// ensuring roadie will act in the expected fashion.
type Config struct {
	Hostname string `json:"hostname"`
	Listener string `json:"listener"`
	Location string `json:"location"`
}

// New will generate a Config instance, pulling from expected locations
// to determine appropriate values for configuration entries.
func New() types.Configer {
	l, err := strconv.Atoi(os.Getenv("PORT"))
	if l < 1 || err != nil {
		l = 8008
	}

	conf := os.Getenv("CONFIG_FILE")
	if conf == "" {
		conf = "/config.roadie.json"
	}

	h := os.Getenv("HOSTNAME")
	if h == "" {
		h = "http://localhost:8080"
	}

	return &Config{
		Hostname: h,
		Listener: fmt.Sprintf(":%d", l),
		Location: conf,
	}
}

// GetConfigFile will provide the path to load the config data from.
func (c *Config) GetConfigFile() string {
	return c.Location
}

// GetHostname will retrieve the currently configured hostname.
func (c *Config) GetHostname() string {
	return c.Hostname
}

// GetListener will provide the listener entry for the HTTP server.
func (c *Config) GetListener() string {
	return c.Listener
}

// GetLogger will provide the pre-prepared logger for use through roadie.
func (c *Config) GetLogger() types.Logger {
	// Under what conditions would this fail?
	z, _ := zap.NewProduction(
		zap.AddCaller(),
		zap.Fields(
			zap.String("app", "roadie"),
			zap.Time("time", time.Now()),
		),
	)
	return z.Sugar()
}
