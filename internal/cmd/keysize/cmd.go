package keysize

import (
	"os"
	"redis_tools/internal/init/sampleinit"
	"redis_tools/internal/server/keysizeserver"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var keySizeVar struct {
	redisHost     string
	redisPort     string
	redisDatabase int
	keyStr        string
	// useKeysCmd    bool
	thresholdSize int
	// outToFile     string
	password  string
	help      bool
	isCluster bool
	samples   int
}

var keySize = &cobra.Command{
	Use:   "查询 redis key 大小",
	Short: "get redis key size",
	Long:  "获取 redis 超过指定大小的 key",
	Run: func(cmd *cobra.Command, args []string) {
		if keySizeVar.help {
			cmd.Help()
			os.Exit(0)
		}

		sampleinit.GlobalVarsInit(
			keySizeVar.redisHost,
			keySizeVar.redisPort,
			keySizeVar.redisDatabase,
		)

		cli, err := sampleinit.RedisClientInit(
			keySizeVar.redisHost,
			keySizeVar.redisPort,
			keySizeVar.password,
			keySizeVar.redisDatabase,
			1, 1, keySizeVar.isCluster,
		)

		if err != nil {
			log.Fatalf("%+v", err)
		}

		// if vcli, ok := cli.(*redisclient.RedisClient); ok {
		if err = keysizeserver.GetAllKeys(cli, keySizeVar.keyStr); err != nil {
			log.Fatalf("%+v", err)
		}
		if err = keysizeserver.GetMemoryToFile(cli, keySizeVar.samples, keySizeVar.thresholdSize); err != nil {
			log.Fatalf("%+v", err)
		}

	},
}

func Execute() error {
	return keySize.Execute()
}

func init() {
	keySize.Flags().BoolVarP(&keySizeVar.help, "help", "", false, "显示帮助信息")
	keySize.Flags().StringVarP(&keySizeVar.redisHost, "host", "h", "127.0.0.1", "redis ip")
	keySize.Flags().StringVarP(&keySizeVar.redisPort, "port", "p", "6379", "redis 端口")
	keySize.Flags().IntVarP(&keySizeVar.redisDatabase, "database", "n", 0, "redis database")
	keySize.Flags().StringVarP(&keySizeVar.keyStr, "key", "k", "*", "key 检索条件")
	// keySize.Flags().BoolVarP(&keySizeVar.useKeysCmd, "use-keys-cmd", "", false, "使用 keys * 获取所有 key,推荐不配置使用 scan")
	keySize.Flags().IntVarP(&keySizeVar.thresholdSize, "thresholdSize", "t", 0, "记录阈值")
	// keySize.Flags().StringVarP(&keySizeVar.outToFile, "outToFile", "o", "./redis_search_out.txt", "结果写入文件中")
	keySize.Flags().StringVarP(&keySizeVar.password, "password", "P", "", "redis 连接密码")
	keySize.Flags().IntVarP(&keySizeVar.samples, "samples", "", 5, "嵌套类型采样嵌套值的数量")
	keySize.Flags().BoolVarP(&keySizeVar.isCluster, "cluster", "c", false, "集群模式")
}
