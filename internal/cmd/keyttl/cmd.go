package keyttl

import (
	"os"
	"redis_tools/internal/init/sampleinit"
	"redis_tools/internal/server/keyttlserver"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var keyTtlVar struct {
	redisHost      string
	redisPort      string
	redisDatabase  int
	keyStr         string
	password       string
	help           bool
	isCluster      bool
	ttlEqual       int
	ttlLessThen    int
	ttlGreaterThen int
}

var keyTtl = &cobra.Command{
	Use:   "查询 redis key 过期时间",
	Short: "get redis key ttl",
	Long:  "统计 redis ttl 匹配某段范围的 key",
	Run: func(cmd *cobra.Command, args []string) {
		if keyTtlVar.help {
			cmd.Help()
			os.Exit(0)
		}

		if keyTtlVar.ttlEqual == -3 && keyTtlVar.ttlGreaterThen == -3 && keyTtlVar.ttlLessThen == -3 {
			cmd.Help()
			log.Fatalf("请指定 ttl 匹配条件")
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

		if err = keyttlserver.GetAllKeys(cli, keyTtlVar.keyStr); err != nil {
			log.Fatalf("%+v", err)
		}

		count := 0
		if count, err = keyttlserver.GetTtl(cli, keyTtlVar.ttlEqual, keyTtlVar.ttlLessThen, keyTtlVar.ttlGreaterThen); err != nil {
			log.Fatalf("%+v", err)
		}
		log.Infof("count:%d", count)

	},
}

func Execute() error {
	return keyTtl.Execute()
}

func init() {
	keyTtl.Flags().BoolVarP(&keyTtlVar.help, "help", "", false, "显示帮助信息")
	keyTtl.Flags().StringVarP(&keyTtlVar.redisHost, "host", "h", "127.0.0.1", "redis ip")
	keyTtl.Flags().StringVarP(&keyTtlVar.redisPort, "port", "p", "6379", "redis 端口")
	keyTtl.Flags().IntVarP(&keyTtlVar.redisDatabase, "database", "n", 0, "redis database")
	keyTtl.Flags().StringVarP(&keyTtlVar.keyStr, "key", "k", "*", "key 检索条件")
	keyTtl.Flags().StringVarP(&keyTtlVar.password, "password", "P", "", "redis 连接密码")
	keyTtl.Flags().BoolVarP(&keyTtlVar.isCluster, "cluster", "c", false, "集群模式")
	keyTtl.Flags().IntVarP(&keyTtlVar.ttlEqual, "eq", "e", -3, "ttl等于(单位为秒)")
	keyTtl.Flags().IntVarP(&keyTtlVar.ttlLessThen, "lt", "l", -3, "ttl小于(单位为秒)")
	keyTtl.Flags().IntVarP(&keyTtlVar.ttlGreaterThen, "gt", "g", -3, "ttl大于(单位为秒)")

}
