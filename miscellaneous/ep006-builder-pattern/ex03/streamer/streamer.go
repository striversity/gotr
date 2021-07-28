package streamer

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	builder interface {
		WithDb(db *sql.DB) builder
		WithQuery(q string) builder
		WithOutputFile(fn string) builder
		WithInterval(interval time.Duration) builder
		Build() *dbStreamConfig
	}

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

func (conf *dbStreamConfig) WithDb(db *sql.DB) builder {
	conf.db = db
	return conf
}

func (conf *dbStreamConfig) WithQuery(q string) builder {
	conf.query = q
	return conf
}

func (conf *dbStreamConfig) WithOutputFile(fn string) builder {
	conf.dest = fn
	return conf
}

func (conf *dbStreamConfig) WithInterval(interval time.Duration) builder {
	conf.interval = interval
	return conf
}

func (conf *dbStreamConfig) Build() *dbStreamConfig {
	fmt.Printf("%#v, interval: %v\n", conf, conf.interval)
	return conf
}

func NewBuilder() builder {
	conf := dbStreamConfig{interval: defaultInterval}

	return &conf
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
