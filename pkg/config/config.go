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
	Listener string `json:"listener"`
	Location string `json:"location"`
}

// New will generate a Config instance, pulling from expected locations
// to determine appropriate values for configuration entries.
func New() (types.Configer, error) {
	l, err := strconv.Atoi(os.Getenv("PORT"))
	if l < 1 || err != nil {
		l = 8008
	}

	conf := os.Getenv("CONFIG_FILE")
	if conf == "" {
		conf = "/config.roadie.json"
	}

	return &Config{
		Listener: fmt.Sprintf(":%d", l),
		Location: conf,
	}, nil
}

// GetConfigFile will provide the path to load the config data from.
func (c *Config) GetConfigFile() string {
	return c.Location
}

// GetListener will provide the listener entry for the HTTP server.
func (c *Config) GetListener() string {
	return c.Listener
}

// GetLogger will provide the pre-prepared logger for use through roadie.
func (c *Config) GetLogger() *zap.SugaredLogger {
	z, err := zap.NewProduction(
		zap.AddCaller(),
		zap.Fields(
			zap.String("app", "roadie"),
			zap.Time("time", time.Now()),
		),
	)
	if err != nil {
		return nil
	}
	sugar := z.Sugar()

	return sugar
}
