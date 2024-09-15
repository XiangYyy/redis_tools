package sampleinit

import (
	"fmt"
	"redis_tools/global"
	"redis_tools/pkg/redisclient"
	"time"
)

func GlobalVarsInit(host, port string, db int) {
	global.ToolsRunTime = time.Now().UnixNano()

	global.TempKeysFile = fmt.Sprintf("tmp-scan-keys-%d",
		global.ToolsRunTime)

	global.OutPutFile = fmt.Sprintf("%s:%s-%d-%d.txt",
		host, port, db, global.ToolsRunTime)

}

func RedisClientInit(host, port string, password string,
	db int, poolSize int, minIdleConnect int, isCluster bool) (interface{}, error) {

	var cli interface{}
	var err error
	if isCluster {
		cli, err = redisclient.NewRedisClusterConnect(host,
			port, password, db, poolSize, minIdleConnect)
	} else {
		cli, err = redisclient.NewRedisConnect(host,
			port, password, db, poolSize, minIdleConnect)
	}
	return cli, err
}

// func RedisClientInit(host, port string, password string,
// 	db int, poolSize int, minIdleConnect int, isCluster bool) (*redisclient.RedisClient, error) {

// 	var err error
// 	cli, err := redisclient.NewRedisConnect(host,
// 		port, password, db, poolSize, minIdleConnect)

// 	return cli, err
// }
