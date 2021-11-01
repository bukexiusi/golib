package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

const cacheExpiration = 7 * 24 * 60 * 60 // 单位秒

func initLuaScriptMap(rdb *Redis) {
	rdb.luaScriptMap = LuaScriptMap{}
	if err := getDelScript(rdb); err != nil {
		panic(err)
	}
	if err := hsetExScript(rdb); err != nil {
		panic(err)
	}
}

func getDelScript(rdb *Redis) error {
	script := redis.NewScript(`
			local result = redis.call('get', KEYS[1]);
			if (result) then
				redis.call('del', KEYS[1]);
			end
			return result;
		`)
	sha, err := script.Load(rdb.cli.Context(), rdb.cli).Result()
	rdb.luaScriptMap["getdel"] = sha
	return err
}

func hsetExScript(rdb *Redis) error {
	script := redis.NewScript(fmt.Sprintf(`
			local result = 0;
			for i = 1, #ARGV, 2 do
				result = result + redis.call('hset', KEYS[1], ARGV[i], ARGV[i+1]);
			end
			redis.call('expire', KEYS[1], '%d');
			return result;
		`, cacheExpiration))
	sha, err := script.Load(rdb.cli.Context(), rdb.cli).Result()
	rdb.luaScriptMap["hsetex"] = sha
	return err
}
