package test

import (
	"testing"
	"time"

	"github.com/bukexiusi/golib/redis"
)

var rdb *redis.Redis

func TestMain(m *testing.M) {
	{
		rdb, _ = redis.Dial(&redis.Configuration{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		defer rdb.Close()
	}
	m.Run()
	{
	}
}

func TestGet(t *testing.T) {
	_ = rdb.Set("key", "val", time.Minute)
	val, _ := rdb.Get("key")
	t.Logf("%q", val)
	val, _ = rdb.GetDel("key")
	t.Logf("%q", val)
	val, _ = rdb.Get("key")
	t.Logf("%q", val)
}
