package keysizeserver

import (
	"bufio"
	"fmt"
	"os"
	"redis_tools/global"
	"redis_tools/pkg/redisclient"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

// func GetAllKeys(cli *redisclient.RedisClient, key string) error {
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

// func GetMemoryToFile(cli *redisclient.RedisClient, samples int, sizeKeyword int) error {
func GetMemoryToFile(cli interface{}, samples int, sizeKeyword int) error {
	log.Infoln("get key's memory size to file")
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
		var size int
		lineText := scanner.Text()

		fmt.Println(len(lineText))

		// if vcli, ok := cli.(*redisclient.RedisClient); ok {
		// 	size, err = vcli.MemoryUsage(lineText, samples)
		// 	// size, err = vcli.MemoryUsage("yanxiangaaaaaaaaaaaa", samples)
		// 	if err != nil {
		// 		return err
		// 	}
		// } else if vcli, ok := cli.(*redisclient.RedisClusterClient); ok {
		// 	size, err = vcli.MemoryUsage(lineText, samples)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		if size >= sizeKeyword {
			if _, err = wbuf.WriteString(fmt.Sprintf("%q,%d\n", lineText, size)); err != nil {
				return err
			}
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
