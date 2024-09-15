package keyhasascii

import (
	"os"
	"redis_tools/internal/init/sampleinit"
	"redis_tools/internal/server/keyhasasciiserver"
	"redis_tools/internal/server/keyttlserver"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var keyTtlVar struct {
	redisHost     string
	redisPort     string
	redisDatabase int
	password      string
	help          bool
	isCluster     bool
}

var keyHasAscii = &cobra.Command{
	Use:   "查询 redis 序列化过的 key",
	Short: "get redis key with Ascii",
	Long:  "导出 Redis 内被 AscII 序列化过的 key",
	Run: func(cmd *cobra.Command, args []string) {
		if keyTtlVar.help {
			cmd.Help()
			os.Exit(0)
		}

		sampleinit.GlobalVarsInit(
			keyTtlVar.redisHost,
			keyTtlVar.redisPort,
			keyTtlVar.redisDatabase,
		)

		cli, err := sampleinit.RedisClientInit(
			keyTtlVar.redisHost,
			keyTtlVar.redisPort,
			keyTtlVar.password,
			keyTtlVar.redisDatabase,
			1, 1, keyTtlVar.isCluster,
		)

		if err != nil {
			log.Fatalf("%+v", err)
		}
		if err = keyttlserver.GetAllKeys(cli, "*"); err != nil {
			log.Fatalf("%+v", err)
		}

		count := 0
		if count, err = keyhasasciiserver.GetASCIIKeys(cli); err != nil {
			log.Fatalf("%+v", err)
		}

		log.Infof("count:%d", count)

	},
}

func Execute() error {
	return keyHasAscii.Execute()
}

func init() {
	keyHasAscii.Flags().BoolVarP(&keyTtlVar.help, "help", "", false, "显示帮助信息")
	keyHasAscii.Flags().StringVarP(&keyTtlVar.redisHost, "host", "h", "127.0.0.1", "redis ip")
	keyHasAscii.Flags().StringVarP(&keyTtlVar.redisPort, "port", "p", "6379", "redis 端口")
	keyHasAscii.Flags().IntVarP(&keyTtlVar.redisDatabase, "database", "n", 0, "redis database")
	keyHasAscii.Flags().StringVarP(&keyTtlVar.password, "password", "P", "", "redis 连接密码")
	keyHasAscii.Flags().BoolVarP(&keyTtlVar.isCluster, "cluster", "c", false, "集群模式")
}
