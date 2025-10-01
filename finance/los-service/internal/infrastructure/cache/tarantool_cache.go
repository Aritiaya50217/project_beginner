package cache

import (
	"log"

	"github.com/tarantool/go-tarantool"
)

func Connect() *tarantool.Connection {
	conn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{})
	if err != nil {
		log.Println("failed to connect to tarantool: ", err)
		return nil
	}
	return conn
}
