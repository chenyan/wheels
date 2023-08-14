package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

const (
	MySQLMaxLifetime         = 59 * time.Second
	MySQLDefaultMaxConn      = 100
	MySQLDefaultTimeout      = 1
	MySQLDefaultReadTimeout  = 5
	MySQLDefaultWriteTimeout = 5
)

type MySQLConf struct {
	Host         string `toml:"host" yaml:"host"`
	Port         int    `toml:"port" yaml:"port"`
	DB           string `toml:"db" yaml:"db"`
	User         string `toml:"user" yaml:"user"`
	Password     string `toml:"password" yaml:"password"`
	Timeout      int    `toml:"timeout" yaml:"timeout"`
	ReadTimeout  int    `toml:"read_timeout" yaml:"read_timeout"`
	WriteTimeout int    `toml:"write_timeout" yaml:"write_timeout"`
	MaxConn      int    `toml:"max_conn" yaml:"max_conn"`
}

func (conf *MySQLConf) DSN() string {
	if conf.Timeout <= 0 {
		conf.Timeout = MySQLDefaultTimeout
	}
	if conf.ReadTimeout <= 0 {
		conf.ReadTimeout = MySQLDefaultReadTimeout
	}
	if conf.WriteTimeout <= 0 {
		conf.WriteTimeout = MySQLDefaultWriteTimeout
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?timeout=%ds&readTimeout=%ds&writeTimeout=%ds&charset=utf8mb4&collation=utf8mb4_unicode_520_ci&parseTime=true&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DB, conf.Timeout, conf.ReadTimeout, conf.WriteTimeout)
}

func (conf *MySQLConf) Gen() (*sql.DB, error) {
	dsn := conf.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(MySQLMaxLifetime)
	if conf.MaxConn <= 0 {
		conf.MaxConn = MySQLDefaultMaxConn
	}
	db.SetMaxOpenConns(conf.MaxConn)
	return db, nil
}

func (conf *MySQLConf) GenSession() (db.Session, error) {
	dsn, err := mysql.ParseURL(conf.DSN())
	if err != nil {
		return nil, err
	}
	r, err := mysql.Open(dsn)
	if err != nil {
		return nil, err
	}
	r.SetConnMaxLifetime(MySQLMaxLifetime)
	if conf.MaxConn <= 0 {
		conf.MaxConn = MySQLDefaultMaxConn
	}
	r.SetMaxOpenConns(conf.MaxConn)
	return r, nil
}

func (conf *MySQLConf) String() string {
	return fmt.Sprintf("mysql[%s/%s]", conf.Host, conf.DB)
}
