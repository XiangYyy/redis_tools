package keyttlserver

import (
	"bufio"
	"fmt"
	"os"
	"redis_tools/global"
	"redis_tools/pkg/redisclient"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func GetAllKeys(cli interface{}, key string) error {
	log.Infof("scan all keys to file %s", global.TempKeysFile)

	if vcli, ok := cli.(*redisclient.RedisClient); ok {
		log.Info("one node")
		_, err := vcli.ScanKeysToFile(key, global.TempKeysFile)
		if err != nil {
			return err
		}
	} else if vcli, ok := cli.(*redisclient.RedisClusterClient); ok {
		log.Info("cluster")
		_, err := vcli.ScanKeysToFile(key, global.TempKeysFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTtl(cli interface{}, ttlEqual, ttlLessThen, ttlGreaterThen int) (int, error) {
	log.Infoln("get matched key's ttl to file")
	log.Infof("read tmp key file %s", global.TempKeysFile)

	matchKeyCount := 0

	keyFile, err := os.Open(global.TempKeysFile)
	if err != nil {
		return 0, errors.Wrapf(err, "open file %s failed",
			global.TempKeysFile)
	}
	scanner := bufio.NewScanner(keyFile)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	log.Infof("open out file %s", global.OutPutFile)
	outFile, err := os.OpenFile(global.OutPutFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, errors.Wrapf(err, "open out file %s failed", global.OutPutFile)
	}
	wbuf := bufio.NewWriter(outFile)

	for scanner.Scan() {
		// var ttl int
		ttl := -2
		lineText := scanner.Text()
		// fmt.Println(lineText)

		if vcli, ok := cli.(*redisclient.RedisClient); ok {
			ttl, err = vcli.TtlKey(lineText)
			// log.Println(ttl)
			if err != nil && err != redis.Nil {
				return 0, err
			}
		} else if vcli, ok := cli.(*redisclient.RedisClusterClient); ok {
			ttl, err = vcli.TtlKey(lineText)
			if err != nil && err != redis.Nil {
				return 0, err
			}
		}

		// fmt.Println("check")
		if MatchTtl(ttl, ttlEqual, ttlLessThen, ttlGreaterThen) {
			// fmt.Println(lineText)
			matchKeyCount++
			if _, err = wbuf.WriteString(fmt.Sprintf("%q,%d\n", lineText, ttl)); err != nil {
				return 0, err
			}
		}
	}

	if err = wbuf.Flush(); err != nil {
		return 0, err
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	keyFile.Close()
	err = RemoveTmpFile()
	if err != nil {
		return 0, err
	}

	return matchKeyCount, nil
}

func RemoveTmpFile() error {
	log.Infof("remove key tmp file: %s", global.TempKeysFile)
	keyFileStatus, err := os.Stat(global.TempKeysFile)
	if err != nil {
		return errors.Wrapf(err, "remove tmp file %s faild",
			global.TempKeysFile)
	}

	if !keyFileStatus.IsDir() {
		return os.Remove(global.TempKeysFile)
	}

	return errors.New(
		fmt.Sprintf("%s if dir continur", global.TempKeysFile))
}

func MatchTtl(ttl, eq, lt, gt int) bool {
	// log.Infof("check")
	// log.Infof("ttl:%d", ttl)
	if eq != -3 && ttl == eq {
		// log.Infof("eq")
		return true
	}

	if lt != -3 && ttl < lt {
		if gt == -3 {
			return true
		} else if ttl > gt {
			return true
		}
	}
	if gt != -3 && ttl > gt {
		return true
	}

	return false
}
