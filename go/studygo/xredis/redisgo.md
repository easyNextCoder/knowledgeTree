
>参考地址
> 
>https://www.cnblogs.com/l-uz/p/16290572.html
>https://blog.csdn.net/Eric_aixiansen/article/details/127777322
>https://pkg.go.dev/github.com/gomodule/redigo/redis
>https://www.cnblogs.com/throwable/p/11644790.html

## redis的通信协议RESP支持的数据类型

RESP支持的数据类型#
RESP在Redis 1.2中引入，在Redis 2.0，RESP正式成为与Redis服务端通信的标准方案。也就是如果需要编写Redis客户端，你就必须在客户端中实现此协议。

RESP本质上是一种序列化协议，它支持的数据类型如下：单行字符串、错误消息、整型数字、定长字符串和RESP数组。

RESP在Redis中用作请求-响应协议的方式如下：

Redis客户端将命令封装为RESP的数组类型（数组元素都是定长字符串类型，注意这一点，很重要）发送到Redis服务器。
Redis服务端根据命令实现选择对应的RESP数据类型之一进行回复。
在RESP中，数据类型取决于数据报的第一个字节：

单行字符串的第一个字节为+。
错误消息的第一个字节为-。
整型数字的第一个字节为:。
定长字符串的第一个字节为$。
RESP数组的第一个字节为*。

## redisgo函数

Executing Commands ¶
The Conn interface has a generic method for executing Redis commands:

```go
    Do(commandName string, args ...interface{}) (reply interface{}, err error)
```
The Redis command reference (http://redis.io/commands) lists the available commands. An example of using the Redis APPEND command is:
```go
    n, err := conn.Do("APPEND", "key", "value")
```

The Do method converts command arguments to bulk strings for transmission to the server as follows:
```go
    Go Type                 Conversion
    []byte                  Sent as is
    string                  Sent as is
    int, int64              strconv.FormatInt(v)
    float64                 strconv.FormatFloat(v, 'g', -1, 64)
    bool                    true -> "1", false -> "0"
    nil                     ""
    all other types         fmt.Fprint(w, v)
```
Redis command reply types are represented using the following Go types:
```go
    Redis type              Go type
    error                   redis.Error
    integer                 int64
    simple string           string
    bulk string             []byte or nil if value not present.
    array                   []interface{} or nil if value not present.
```
Use type assertions or the reply helper functions to convert from interface{} to the specific Go type for the command result.

## redis中的hash的底层数据结构

hash中使用了两种数据结构ziplist和hashtable
当哈希对象可以同时满足以下两个条件时， 哈希对象使用 ziplist 编码：

哈希对象保存的所有键值对的键和值的字符串长度都小于 64 字节；
* 哈希对象保存的键值对数量小于 512 个；
* 不能满足这两个条件的哈希对象需要使用 hashtable 编码。
>http://redisbook.com/preview/object/hash.html

## redis进程配置信息

#### redis 操作

* 查看redis的运行状态
  **systemctl status redis**

  ```linux
  ● redis-server.service - Advanced key-value store
     Loaded: loaded (/lib/systemd/system/redis-server.service; enabled; vendor preset: enabled)
     Active: active (running) since Thu 2022-07-28 10:44:37 CST; 1 months 24 days ago
       Docs: http://redis.io/documentation,
             man:redis-server(1)
   Main PID: 1247 (redis-server)
      Tasks: 3
     Memory: 219.6M
        CPU: 23h 38min 52.213s
     CGroup: /system.slice/redis-server.service
             └─1247 /usr/bin/redis-server 0.0.0.0:6379         
  
  Warning: Journal has been rotated since unit was started. Log output is incomplete or unavailable.
  
  ```

  通过loaded(xxx)这一行发现redis运行systemctl的配置文件在/lib/systemd/system/redis-server.service，然后

  **cat /lib/systemd/system/redis-server.service**

  ```
  [Unit]
  Description=Advanced key-value store
  After=network.target
  Documentation=http://redis.io/documentation, man:redis-server(1)
  
  [Service]
  Type=forking
  ExecStart=/usr/bin/redis-server /etc/redis/redis.conf
  PIDFile=/var/run/redis/redis-server.pid
  TimeoutStopSec=0
  Restart=always
  User=redis
  Group=redis
  
  xxx
  ```

  最终发现
  * 日志文件在：/var/log/redis/redis-server.log 
  * 配置文件在/etc/redis/redis.conf
    redis.conf中的一个配置项：

      * stop-writes-on-bgsave-error yes
