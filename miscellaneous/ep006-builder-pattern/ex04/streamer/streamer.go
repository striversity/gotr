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
		Build() (*dbStreamConfig, error)
	}

	builderImpl struct {
		conf *dbStreamConfig
		err  error
	}

	dbStreamConfig struct {
		db       *sql.DB       // source from which to read data
		interval time.Duration // how often to read from source
		query    string        // query to run against source
		dest     string        // destination to store data
	}
)

var (
	defaultInterval = 5 * time.Second
)

func (bi *builderImpl) WithDb(db *sql.DB) builder {
	if db == nil {
		bi.err = fmt.Errorf("db can't be nil: %w", bi.err)
		return bi
	}

	bi.conf.db = db
	return bi
}

func (bi *builderImpl) WithQuery(q string) builder {
	if q == "" {
		bi.err = fmt.Errorf("query string can't be empty: %w", bi.err)
		return bi
	}

	bi.conf.query = q
	return bi
}

func (bi *builderImpl) WithOutputFile(fn string) builder {
	if fn == "" {
		bi.err = fmt.Errorf("output filename can't be empty: %w", bi.err)
		return bi
	}

	bi.conf.dest = fn
	return bi
}

func (bi *builderImpl) WithInterval(interval time.Duration) builder {
	if interval == time.Duration(0) {
		bi.err = fmt.Errorf("interval must be greater than 0: %w", bi.err)
		return bi
	}

	bi.conf.interval = interval
	return bi
}

func (bi *builderImpl) Build() (*dbStreamConfig, error) {
	fmt.Printf("%#v, interval: %v\n", bi.conf, bi.conf.interval)
	return bi.conf, bi.err
}

func NewBuilder() builder {
	bi := &builderImpl{
		conf: &dbStreamConfig{interval: defaultInterval},
	}

	return bi
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
