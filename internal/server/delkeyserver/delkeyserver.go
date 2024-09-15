package delkeyserver

import (
	"bufio"
	"fmt"
	"os"
	"redis_tools/global"
	"redis_tools/pkg/redisclient"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func GetAllKeys(cli interface{}, key string) (int, error) {
	log.Infof("scan all keys to file %s", global.TempKeysFile)
	var count int = 0
	var err error

	if vcli, ok := cli.(*redisclient.RedisClient); ok {
		log.Info("one node")
		count, err = vcli.ScanKeysToFile(key, global.TempKeysFile)
		if err != nil {
			return 0, err
		}
	} else if vcli, ok := cli.(*redisclient.RedisClusterClient); ok {
		log.Info("cluster")
		count, err = vcli.ScanKeysToFile(key, global.TempKeysFile)
		if err != nil {
			return 0, err
		}
	}
	// count, err := cli.ScanKeysToFile(key, global.TempKeysFile)
	// if err != nil {
	// 	return 0, err
	// }
	return count, nil
}

func DoDelKeys(cli interface{}) error {
	log.Infoln("start del key")
	log.Infof("read tmp key file %s", global.TempKeysFile)
	keyFile, err := os.Open(global.TempKeysFile)
	if err != nil {
		return errors.Wrapf(err, "open file %s failed",
			global.TempKeysFile)
	}
	scanner := bufio.NewScanner(keyFile)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	log.Infof("open out file %s", global.OutPutFile)
	outFile, err := os.OpenFile(global.OutPutFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return errors.Wrapf(err, "open out file %s failed", global.OutPutFile)
	}
	wbuf := bufio.NewWriter(outFile)

	for scanner.Scan() {
		lineText := scanner.Text()
		var delCount int
		var err error

		if vcli, ok := cli.(*redisclient.RedisClient); ok {
			delCount, err = vcli.DelKey(lineText)
			log.Infof("del %q success: %v", lineText, delCount)
			if err != nil {
				return err
			}
		} else if vcli, ok := cli.(*redisclient.RedisClusterClient); ok {
			delCount, err = vcli.DelKey(lineText)
			log.Infof("del %q success: %v", lineText, delCount)
			if err != nil {
				return err
			}
		}

		if _, err = wbuf.WriteString(fmt.Sprintf("%q,%d\n", lineText, delCount)); err != nil {
			return err
		}
	}

	if err = wbuf.Flush(); err != nil {
		return err
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	keyFile.Close()
	err = RemoveTmpFile()
	if err != nil {
		return err
	}
	return nil
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
