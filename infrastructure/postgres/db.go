package postgres

import (
	"fmt"

	"xorm.io/xorm"
)

// DB .
type DB struct {
	engine *xorm.Engine
}

// Close .
func (db *DB) Close() {
	db.engine.Close()
}

// DBConfig .
type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
}

// NewDB .
func NewDB(cfg DBConfig) (DB, error) {
	engine, err := xorm.NewEngine("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name,
	))

	if err != nil {
		return DB{}, err
	}

	return DB{
		engine: engine,
	}, nil
}
