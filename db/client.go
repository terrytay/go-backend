package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
)

type IClient interface {
	Initialise() error
}

type Client struct {
	db *pg.DB
}

func (d *Client) Initialise(connectionConfig Config) error {
	if connectionConfig == (Config{}) {
		return errors.New("no config defined for database connection")
	}

	d.db = pg.Connect(&pg.Options{
		Addr:     connectionConfig.Addr,
		User:     connectionConfig.User,
		Password: connectionConfig.Password,
		Database: connectionConfig.Database,
	})

	ctx := context.Background()

	if err := d.db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection to database established")

	return nil
}
