package streamer

import (
	"database/sql"
	"time"
)

type (
	dbStreamConfig struct {
		db       *sql.DB       // source from which to read data
		interval time.Duration // how often to read from source
		query    string        // query to run against source
		dest     string        // destination to store data
	}
)

func New(options ...interface{}) (*dbStreamConfig, error) {
	conf := &dbStreamConfig{}

	for _, opt := range options {
		if v, ok := opt.(string); ok {
			// what should we store, is it the destionation or the query
			conf.query = v
		}
	}
	// db: db,
	// query: query,
	// interval: interval,
	// dest: dest,

	return conf, nil
}

func (conf *dbStreamConfig) Start() error {
	return nil
}

func (conf *dbStreamConfig) Stop() error {
	return nil
}
