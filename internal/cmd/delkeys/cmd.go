package delkeys

import (
	"fmt"
	"os"
	"redis_tools/internal/init/sampleinit"
	"redis_tools/internal/server/delkeyserver"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var delKeyVar struct {
	redisHost     string
	redisPort     string
	redisDatabase int
	keyStr        string
	// useKeysCmd    bool
	// outToFile     string
	password  string
	count     bool
	isCluster bool
	help      bool
}

var delKey = &cobra.Command{
	Use:   "删除 redis key",
	Short: "del redis key size",
	Long:  "循环删除 redis key",
	Run: func(cmd *cobra.Command, args []string) {
		if delKeyVar.help {
			cmd.Help()
			os.Exit(0)
		}

		sampleinit.GlobalVarsInit(
			delKeyVar.redisHost,
			delKeyVar.redisPort,
			delKeyVar.redisDatabase,
		)

		cli, err := sampleinit.RedisClientInit(
			delKeyVar.redisHost,
			delKeyVar.redisPort,
			delKeyVar.password,
			delKeyVar.redisDatabase,
			1, 1, delKeyVar.isCluster,
		)

		if err != nil {
			log.Fatalf("%+v", err)
		}
		if delKeyVar.keyStr == "*" {
			log.Fatal("请使用 -k 或 --key 指定 scan 的条件")
		}

		kCount, err := delkeyserver.GetAllKeys(cli, delKeyVar.keyStr)
		if err != nil {
			log.Fatalf("%+v", err)
		}

		log.Infof("要删除的key总计: %v", kCount)
		if delKeyVar.count {
			os.Exit(0)
		}

		var input string
		fmt.Print("是否开始删除? (y/n): ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			log.Fatal("Error reading input:", err)
		}

		if input == "y" || input == "Y" {
			if err = delkeyserver.DoDelKeys(cli); err != nil {
				log.Fatalf("%+v", err)
			}
		} else {
			log.Info("exit")
			os.Exit(0)
		}

	},
}

func Execute() error {
	return delKey.Execute()
}

func init() {
	delKey.Flags().BoolVarP(&delKeyVar.help, "help", "", false, "显示帮助信息")
	delKey.Flags().BoolVarP(&delKeyVar.count, "count", "c", false, "计数")
	delKey.Flags().StringVarP(&delKeyVar.redisHost, "host", "h", "127.0.0.1", "redis ip")
	delKey.Flags().StringVarP(&delKeyVar.redisPort, "port", "p", "6379", "redis 端口")
	delKey.Flags().IntVarP(&delKeyVar.redisDatabase, "database", "n", 0, "redis database")
	delKey.Flags().StringVarP(&delKeyVar.keyStr, "key", "k", "*", "key 检索条件")
	// delKey.Flags().BoolVarP(&delKeyVar.useKeysCmd, "use-keys-cmd", "", false, "使用 keys * 获取所有 key,推荐不配置使用 scan")
	// delKey.Flags().StringVarP(&delKeyVar.outToFile, "outToFile", "o", "./redis_search_out.txt", "结果写入文件中")
	delKey.Flags().StringVarP(&delKeyVar.password, "password", "P", "", "redis 连接密码")
	delKey.Flags().BoolVarP(&delKeyVar.isCluster, "cluster", "", false, "为集群模式")
}
