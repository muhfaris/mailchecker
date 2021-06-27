package router

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

// App is wrap app data
type App struct {
	Hostname    string
	Port        int
	Environment string
	Config      ConfigApp
	Connection  ConnectionApp
}

func Init() *App {
	return &App{
		Hostname:    viper.GetString("app.hostname"),
		Port:        viper.GetInt("app.port"),
		Environment: viper.GetString("app.environment"),
		Config: ConfigApp{
			PSQL: PGXApp{
				Hostname: viper.GetString("connections.pgx.hostname"),
				Port:     viper.GetInt("connections.pgx.port"),
				Username: viper.GetString("connections.pgx.username"),
				Password: viper.GetString("connections.pgx.password"),
			},
		},
		Connection: ConnectionApp{
			PSGQL: initDatabase(),
		},
	}
}

// ConfigApp is all connection for app, database, redis, everything
type ConfigApp struct {
	PSQL PGXApp
}

// PGXApp is postgresql database config
type PGXApp struct {
	Hostname string
	Port     int
	Username string
	Password string
}

// ConnectionApp is connection for app
type ConnectionApp struct {
	PSGQL *pgxpool.Pool
}

func initDatabase() *pgxpool.Pool {
	return nil
}
