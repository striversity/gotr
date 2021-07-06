package streamer

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	dbStreamConfig struct {
		db       *sql.DB       // source from which to read data
		interval time.Duration // how often to read from source
		query    string        // query to run against source
		dest     string        // destination to store data
	}

	config func(c dbStreamConfig) dbStreamConfig
)

var (
	defaultInterval = 5 * time.Second
)

func WithDb(db *sql.DB) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.db = db
		return c
	}
}

func WithQuery(q string) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.query = q
		return c
	}
}

func WithOutputFile(fn string) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.dest = fn
		return c
	}
}

func WithInterval(interval time.Duration) config {
	return func(c dbStreamConfig) dbStreamConfig {
		c.interval = interval
		return c
	}
}

func New(options ...config) (*dbStreamConfig, error) {
	conf := dbStreamConfig{interval: defaultInterval}

	for _, opt := range options {
		conf = opt(conf)
	}

	fmt.Printf("%#v, interval: %v\n", conf, conf.interval)
	return &conf, nil
}

func (conf *dbStreamConfig) Start() error {
	go func(c *dbStreamConfig) {
		for {
			logrus.Infof("fetching data at %v", time.Now())
			time.Sleep(c.interval)
		}
	}(conf)
	return nil
}

func (conf *dbStreamConfig) Stop() error {
	return nil
}
