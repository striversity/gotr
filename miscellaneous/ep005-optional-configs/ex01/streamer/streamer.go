package streamer

import (
	"database/sql"
	"time"
)

type (
	DBStreamConfig struct {
		db       *sql.DB       // source from which to read data
		interval time.Duration // how often to read from source
		query    string        // query to run against source
		dest     string        // destination to store data
		Field    string
	}
)

func (conf *DBStreamConfig) Start() error {
	return nil
}

func (conf *DBStreamConfig) Stop() error {
	return nil
}
