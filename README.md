# Redis 工具

## 工具说明
+ cmd/keysize：**获取redis中的大key**，查询 Redis 中，大小超过指定大小的 key
+ cmd/delkeys：**批量删除redis中的key**,使用匹配的模式，如 key1\*  的方式批量删除key
+ cmd/keyttl：**统计redis中过期时间满足条件的key**，用于查询 redis 中无过期时间或过期时间超过多久的 key
+ cmd/keyhasascii：**导出Redis内被AscII序列化过的key**

## keysize
+ 说明:先通过 scan 命令拿到所有的 key，写入文件中，然后遍历文件内的 key，使用 *MEMORY USAGE* 获取 key 占用的字节数，超过指定数值的 key 被记录入文件中。

### 参数说明
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
+ 查询 *127.0.0.1:6379* 中符合 *\*keyWorkd\** 大小超过 20000B 的 key

```bash
./keysize -h 127.0.0.1 -p 6379 -k  '*keyWorkd*' -t 20000
```

## delkeys
+ 支持按照 keyword\* 模式批量删除 redis key
+ -k 参数不允许为空

### 参数说明
```
Usage:
  删除 redis key [flags]

Flags:
      --cluster           集群模式
  -c, --count             计数，仅统计满足删除条件的 key 个数，不实际删除
  -n, --database int      redis database
      --help              显示帮助信息
  -h, --host string       redis ip (default "127.0.0.1")
  -k, --key string        key 检索条件 (default "*")
  -P, --password string   redis 连接密码
  -p, --port string       redis 端口 (default "6379")
```

### 例
+ 仅根据条件匹配符合条件的行数

```bash
./delkeys -h 127.0.0.1 -p 6379 -k "*keyWord*" -c
```

+ 删除

```bash
./delkeys -h 127.0.0.1 -p 6379  -k *keyWord*
INFO[0017] 要删除的key总计: 0
是否开始删除? (y/n): y
INFO[0019] start del key
```

## keyttl
+ 查询 key 过期时间，用于检索无过期时间或过期时间超过多久的key

### 参数说明
```
Flags:
  -c, --cluster           为集群模式
  -n, --database int      redis database
  -e, --eq int            ttl等于(单位为秒) (default -3)
  -g, --gt int            ttl大于(单位为秒) (default -3)
      --help              显示帮助信息
  -h, --host string       redis ip (default "127.0.0.1")
  -k, --key string        key 检索条件 (default "*")
  -l, --lt int            ttl小于(单位为秒) (default -3)
  -P, --password string   redis 连接密码
  -p, --port string       redis 端口 (default "6379")
```

### 例
+ 查看 redis 1 库中无过期时间的 key

```bash
./keyttl -h 127.0.0.1 -p 6379 -k "*" -n  1 -e -1
```
+ 查询 redis 1 库中过期时间超过 1h 的key

```bash
./keyttl -h 127.0.0.1 -p 6379 -k "*" -n 1 -g 3600
```

## keyhasascii
+ 导出 Redis 内被 AscII 序列化过的 key

### 参数说明
```
Flags:
  -c, --cluster           集群模式
  -n, --database int      redis database
      --help              显示帮助信息
  -h, --host string       redis ip (default "127.0.0.1")
  -P, --password string   redis 连接密码
  -p, --port string       redis 端口 (default "6379")
```
