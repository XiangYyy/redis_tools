# Redis 工具

## 工具说明
+ cmd/keysize：**获取redis中的大key**，查询 Redis 中，大小超过指定大小的 key
+ cmd/delkeys：**批量删除redis中的key**,使用匹配的模式，如 key1\*  的方式批量删除key

## keysize
### 参数说明
+ 说明:先通过 scan 命令拿到所有的 key，写入文件中，然后遍历文件内的 key，使用 *MEMORY USAGE* 获取 key 占用的字节数，超过指定数值的 key 被记录入文件中。
+ 参数说明

```
--help : 显示帮助信息
-h, --host string：指定 redis 的 IP，如 127.0.0.1
-p, --port string：指定 redis 的端口
-P, --password string：指定 redis 的访问密码
-n, --database int：指定 redis 查询哪个 database 下 key 的大小
-k, --key string：指定检索条件，默认为 *，即检索 keys * 所匹配的key
-t, --thresholdSize int：指定要被记录的key的大小的阈值(单位:B)，默认为 0 ，即记录所有 key 的大小
--samples int：嵌套类型采样嵌套值的数量(default 5)
-c, --cluster：指定链接的集群为 redis cluster
```

### 例
+ 查询 *127.0.0.1:6379* 中符合 *\*sellCardBatchKey\** 大小超过 20000B 的 key

```bash
./keysize -h 127.0.0.1 -p 6379 -k  '*sellCardBatchKey*' -t 20000
```
