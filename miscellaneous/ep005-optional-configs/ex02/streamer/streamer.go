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

func New(db *sql.DB, query string, interval time.Duration, dest string) (*dbStreamConfig, error) {
	conf := &dbStreamConfig{
		db: db, 
		query: query, 
		interval: interval, 
		dest: dest,
	}

	return conf, nil
}

func (conf *dbStreamConfig) Start() error {
	return nil
}

func (conf *dbStreamConfig) Stop() error {
	return nil
}
