package pg

import (
	"context"
	"fmt"

	//"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	DBName   string `json:"dbName"`
	SSLmode  string `json:"sslMode"`
	Debug    string `json:"debug"`
}

func (c *Config) ConnectPostgres(ctx context.Context) (*sqlx.DB, error) {
	//db, err := pgx.Connect(ctx, dbc.genConnStr())
	//"172.17.0.1" c ним из контенера сервера цепляется в базе
	driverName := "postgres"
	db, err := sqlx.Connect(driverName, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName))

	if err != nil {
		return nil, fmt.Errorf("(c *Config)ConnectPostgres #1 \n Error:%s \n", err.Error())
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("(c *Config)ConnectPostgres #2 \n Error:%s \n", err.Error())
	}

	return db, nil
}


