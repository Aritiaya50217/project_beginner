package cache

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tarantool/go-tarantool"
)

type TarantoolCache struct {
	Conn *tarantool.Connection
}

func NewTarantoolCache() *TarantoolCache {
	host := os.Getenv("TARANTOOL_HOST")
	port := os.Getenv("TARANTOOL_PORT")

	if host == "" || port == "" {
		log.Fatal("Tarantool environment variables are not set")
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	opts := tarantool.Opts{
		Timeout:        1 * time.Second,
	}

	var conn *tarantool.Connection
	var err error

	// Retry connection 5 ครั้ง
	for i := 0; i < 5; i++ {
		conn, err = tarantool.Connect(addr, opts)
		if err == nil {
			log.Println("Connected to Tarantool!")
			break
		}
		log.Printf("Retrying Tarantool connection... (%d/5)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("failed to connect to Tarantool: %v", err)
	}
	return &TarantoolCache{Conn: conn}
}

func (c *TarantoolCache) Set(key string, value interface{}) error {
	_, err := c.Conn.Call("box.space.cache.insert", []interface{}{key, value})
	return err
}

func (c *TarantoolCache) Get(key string) (interface{}, error) {
	resp, err := c.Conn.Call("box.space.cache:get", []interface{}{key})
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, nil
	}

	// try assertion : resp.Data[0] เป็น []interface{}
	row, ok := resp.Data[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type for resp.Data[0]: %T", resp.Data[0])
	}

	if len(row) < 2 {
		return nil, fmt.Errorf("unexpected row length: %d", len(row))
	}

	return row[1], nil
}
