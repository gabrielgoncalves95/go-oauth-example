package config

import "fenrir/support"

var current Config

type Config struct {
	Database Postgres
}
type Postgres struct {
	Host string
	Port string
	User string
	Pwd  string
	Db   string
}

func buildPostgres() Postgres {
	host := support.GetEnvStr("DATABASE_HOST", "localhost")
	port := support.GetEnvStr("DATABASE_PORT", "5432")
	user := support.GetEnvStr("DATABASE_USER", "postgres")
	pwd := support.GetEnvStr("DATABASE_PWD", "postgres")
	db := support.GetEnvStr("DATABASE_DB", "oauth-test")
	return Postgres{
		Host: host,
		Port: port,
		User: user,
		Pwd:  pwd,
		Db:   db,
	}
}

func init() {
	current = Config{
		Database: buildPostgres(),
	}
}
func Current() Config {
	return current
}
