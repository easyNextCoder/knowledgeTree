package xredis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sort"
	"strconv"
	"studygo/xstring"
	"time"
)

type Agent struct {
	Proxy  string `json:"proxy,omitempty"`
	UrlPos int    `json:"url_pos,omitempty"`
}

type Player struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Agent *Agent `json:"agent,omitempty"`
}

func redisWork() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	p := &Player{
		Name: "xiaoming",
		Age:  2,
		Agent: &Agent{
			Proxy:  "google.com",
			UrlPos: 15,
		},
	}

	bytes, _ := json.Marshal(&p)
	c1.Do("Set", "player", bytes)
	fmt.Println("marshal bytes ", string(bytes))

	nBytes, _ := redis.Bytes(c1.Do("Get", "player"))
	fmt.Println("marshal nBytes ", string(nBytes))

	np := &Player{}
	json.Unmarshal(nBytes, np)
	fmt.Println(np)

}

func redisType() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	//args := redis.Args{}

	c1.Do("hset", "m", 1, 1)
	c1.Do("HINCRBY", "n", 2, 2)
	do, err := c1.Do("hmget", "n", 2)
	if err != nil {
		return
	}
	fmt.Println(do)
}

func bigKey() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	//args := redis.Args{}

	startTime := time.Now()
	outMax := 1
	max := 1000000
	for i := 0; i < outMax; i++ {
		key := "key" + strconv.Itoa(i)
		for j := 0; j < max; j++ {
			innerKey := "innerKey" + strconv.Itoa(j)
			val := "val" + strconv.Itoa(j)
			c1.Do("hset", key, innerKey, val)
		}
	}
	endTime := time.Now()
	fmt.Println("time cost", endTime.Sub(startTime).Milliseconds())

}

func getBigKey() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	start := time.Now()

	stringMap, err := c1.Do("hgetall", "key0")
	if err != nil {
		return
	}
	end := time.Now()

	fmt.Println("getBigKey time", end.Sub(start).Milliseconds())

	cnt := 0
	for i, v := range stringMap.([]interface{}) {
		fmt.Println(i, string(v.([]byte)))
		cnt++
		if cnt > 10 {
			break
		}
	}

}

func runWork() {
	getBigKey()
}

func sortedSetOperate() {
	conn := GetRedis()
	defer conn.Close()

	conn.Do(del, "set1")
	conn.Do(zadd, "set1", "2", "xyk", "5", "xyq", "3", "xyl", "1", "xyz", "7", "xym", "10", "xyu")
	val, _ := redis.Int(conn.Do(zcard, "set1"))
	val1, _ := redis.Strings(conn.Do(zrange, "set1", 0, -1, "withscores")) //放入的时候会自动排序,withscores可选
	fmt.Println("zadd zcard", val, val1)

	val2, _ := redis.Strings(conn.Do(zrevrange, "set1", 0, -1))
	fmt.Println("zrevrange", val2)

	val3, _ := redis.Strings(conn.Do(zrangebyscore, "set1", 1, 3))
	fmt.Println("zrangebyscore", val3)

	val4, _ := redis.Strings(conn.Do(zrevrangebyscore, "set1", 3, 1))
	fmt.Println("zrevrangebyscore", val4)

	val5, _ := redis.Int(conn.Do(zscore, "set1", "xyk"))
	fmt.Println("zscore", val5)

	val6, _ := redis.Int(conn.Do(zrank, "set1", "xyl")) //不支持同时获取多个值
	val7, _ := redis.Int(conn.Do(zrank, "set1", "xyz"))
	fmt.Println("zrank", val6, val7)

	val8, _ := redis.Int(conn.Do(zrevrank, "set1", "xyz"))
	fmt.Println("zrevrand", val8)

	val9, _ := redis.Int(conn.Do(zcount, "set1", 3, 7))
	fmt.Println("zscore", val9)

	conn.Do(del, "set2")
	conn.Do(del, "set3")
	conn.Do(zadd, "set2", "2", "xyk", "5", "xyq", "3", "xyl", "1", "xyz", "7", "xym", "15", "xyu", "100", "xyy")
	conn.Do(zadd, "set3", "2", "xyk", "5", "xyq", "3", "xyl", "1", "xyz", "7", "xym", "15", "xyu", "100", "xyp")

	conn.Do(del, "des_set")
	val10, _ := redis.Int(conn.Do(zinterstore, "des_set", 2, "set2", "set3", "weights", 0.5, 0.5)) //2代表连个两个集合的交集，weights是可选的，得到dest_set中元素的分数
	fmt.Println("zinterstore", val10)
	val11, _ := redis.Strings(conn.Do(zrange, "des_set", 0, -1, "withscores"))
	fmt.Println("zinterstore", val11)

	conn.Do(del, "des_set2")
	val12, _ := redis.Int(conn.Do(zunionstore, "des_set2", 2, "set2", "set3", "weights", 1, 1)) //跟zinterstore的效果是一样的
	fmt.Println("zunionstore", val12)
	val13, _ := redis.Strings(conn.Do(zrange, "des_set2", 0, -1, "withscores"))
	fmt.Println("zunionstore", val13)

	conn.Do(del, "set3")
	conn.Do(zadd, "set3", "2", "xyk", "5", "xyq", "3", "xyl")
	val14, _ := redis.Float64(conn.Do(zincrby, "set3", -1, "xyk")) //注意几个参数的位置
	val15, _ := redis.String(conn.Do(zscore, "set3", "xyk"))
	fmt.Println("zincrby", val14, val15)

	conn.Do(del, "set3")
	conn.Do(zadd, "set3", "2", "a", "5", "b", "3", "c", 4, "d", 6, "a")
	val16, _ := redis.Strings(conn.Do(zrange, "set3", 0, -1, "withscores"))
	val17, _ := redis.Bool(conn.Do(zrem, "set3", "a"))
	val18, _ := redis.Strings(conn.Do(zrange, "set3", 0, -1))
	fmt.Println("zrem", val16, val17, val18)

	conn.Do(del, "set3")
	conn.Do(zadd, "set3", "2", "a", "5", "b", "3", "c", 4, "d", 6, "a")
	val19, _ := redis.Strings(conn.Do(zrange, "set3", 0, -1, "withscores"))
	val20, _ := redis.Bool(conn.Do(zremrangebyrank, "set3", 0, 2)) //下标从0开始，[0,2]
	val21, _ := redis.Strings(conn.Do(zrange, "set3", 0, -1))
	fmt.Println("zremrangerank", val19, val20, val21)

	conn.Do(del, "set3")
	conn.Do(zadd, "set3", "2", "a", "5", "b", "3", "c", 4, "d", 6, "a")
	val22, _ := redis.Strings(conn.Do(zrange, "set3", 0, -1, "withscores"))
	val23, _ := redis.Bool(conn.Do(zremrangebyscore, "set3", 4, 6)) //下标从0开始，[0,2]
	val24, _ := redis.Strings(conn.Do(zrange, "set3", 0, -1))
	fmt.Println("zremrangescore", val22, val23, val24)

	conn.Do(del, "set4")
	conn.Do(zadd, "set4", "2", "a", 2, "aa", "3", "b", "4", "c", 5, "d", 6, "e", 6, "f", 6, "g")

	val25, _ := redis.Strings(conn.Do(zrange, "set4", 0, -1, "withscores"))
	val26, _ := redis.Bool(conn.Do(zpopmax, "set4", 2)) //会先pop出同分最大的，之后找最大的，直到为空
	val27, _ := redis.Strings(conn.Do(zrange, "set4", 0, -1))
	val28, _ := redis.Bool(conn.Do(zpopmin, "set4", 20))
	val29, _ := redis.Strings(conn.Do(zrange, "set4", 0, -1))
	fmt.Println("zpopmax zpopmin", val25, val26, val27, val28, val29)

	//zscan
	conn.Do("flushall")
	for i := 0; i < 1000; i++ {
		str := xstring.RandString(80)
		k1, _, k2, _ := fmt.Sprintf("key%02d%s", i, str), "val"+strconv.Itoa(i), fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i)
		conn.Do("zadd", "rootKey", i, k1) //这里分数的位置是在前面
		conn.Do("zadd", "rootKey", i+2, k2)
	}
	var pos int = 0
	var res []string
	xscan_(conn, zscan, "rootKey", &pos, 10, 10, "key*", &res)
	sort.StringSlice(res).Sort()
	fmt.Println("zscan res:", len(res), res)

}

func setOperate() {
	conn := GetRedis()
	defer conn.Close()

	conn.Do(sadd, "set1", "xyk", "xyq", "xyl", "xyz", "xym", "xyu")
	val, _ := redis.Int(conn.Do(scard, "set1"))
	fmt.Println("sadd scard", val)

	val2, _ := redis.Bool(conn.Do(sismember, "set1", "xyk"))
	fmt.Println("sismember", val2)

	val3, _ := redis.Strings(conn.Do(srandmember, "set1", 2))
	fmt.Println("srandmember", val3)

	val4, _ := redis.Strings(conn.Do(smembers, "set1"))
	fmt.Println("smembers", val4)

	conn.Do(sadd, "set2", "xyk", "xyq", "xyl", "xyz", "xyv")

	val5, _ := redis.Strings(conn.Do(sdiff, "set1", "set2")) //set1-set2得到的结果
	fmt.Println("sdiff", val5)

	val7, _ := redis.Strings(conn.Do(sinter, "set1", "set2"))
	fmt.Println("sinter", val7)

	val8, _ := redis.Strings(conn.Do(sunion, "set1", "set2"))
	fmt.Println("sunion", val8)

	val9, _ := redis.Strings(conn.Do(spop, "set2"))
	val92, _ := redis.Strings(conn.Do(smembers, "set2"))
	fmt.Println("spop", val9, val92)

	val10, _ := redis.Strings(conn.Do(srem, "set1", "xyk", "xyq"))
	val102, _ := redis.Strings(conn.Do(smembers, "set1"))
	fmt.Println("srem", val10, val102)

	//sscan
	conn.Do("flushall")
	for i := 0; i < 1000; i++ {
		str := xstring.RandString(80)
		k1, _, k2, _ := fmt.Sprintf("key%02d%s", i, str), "val"+strconv.Itoa(i), fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i)
		conn.Do("sadd", "rootKey", k1)
		conn.Do("sadd", "rootKey", k2)
	}
	var pos int = 0
	var res []string
	xscan_(conn, sscan, "rootKey", &pos, 10, 10, "name*", &res)
	sort.StringSlice(res).Sort()
	fmt.Println("sscan res:", len(res), res)

}

func hashOperate() {
	conn := GetRedis()
	defer conn.Close()

	//conn.Do(hset, "user:1:info", "age", 23)
	//conn.Do(hset, "user:1:info", "height", 160)
	//val, _ := redis.Int(conn.Do(hget, "user:1:info", "age"))
	//fmt.Println("hset hget", val)
	//
	////hgetall
	//val2, err := redis.Strings(conn.Do(hgetall, "user:1:info"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("hgetall as []string", val2)
	//
	//val3, err := redis.StringMap(conn.Do(hgetall, "user:1:info"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("hgetall as map[string]string", val3)
	//
	//val4, err := redis.Bool(conn.Do(hexists, "user:1:info", "age"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("hexists", val4)
	//
	//val5, err := redis.Int(conn.Do(hlen, "user:1:info"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("hlen", val5)
	//
	//val6, err := redis.String(conn.Do(hmset, "user:1:info", "weight", 60, "hair", "long_black")) //返回的竟然是"OK"
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("hmset", val6)
	//val62, err := redis.StringMap(conn.Do(hgetall, "user:1:info"))
	//fmt.Println("hmset", val62)

	val7, err := redis.Ints(conn.Do(hmget, "user:1:info", "weight", "hair")) //这里如果用stringMap来获取的话会得到不合逻辑的结果，stringMap只是把每两个字符串转化为key:pair
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hmget", val7)

	val8, err := redis.Strings(conn.Do(hvals, "user:1:info"))
	fmt.Println("hvals", val8)

	val9, err := redis.Strings(conn.Do(hkeys, "user:1:info"))
	fmt.Println("hkeys", val9)

	val10, err := redis.Bool(conn.Do(hincrby, "user:1:info", "weight", -8)) //正负值都支持
	fmt.Println("hincrby", val10)
	val11, err := redis.String(conn.Do(hget, "user:1:info", "weight"))
	fmt.Println("hincrby hget", val11)

	conn.Do(hdel, "user:1:info", "weight")
	val12, err := redis.StringMap(conn.Do(hgetall, "user:1:info"))
	fmt.Println("hdel", val12)

	//hscan
	conn.Do("flushall")
	for i := 0; i < 1000; i++ {
		str := xstring.RandString(80)
		k1, _, k2, _ := fmt.Sprintf("key%02d%s", i, str), "val"+strconv.Itoa(i), fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i)
		conn.Do("hmset", "rootKey", k1, "val1") //scan既可以扫描出字符串，对于集合和哈希等一样能够扫描出来
		conn.Do("hmset", "rootKey", k2, "val2")
	}
	var pos int = 0
	var res []string
	xscan_(conn, hscan, "rootKey", &pos, 10, 10, "name*", &res)
	//当hash对象所保存的键的数量小于512且key和value的长度都小于64个字节的时候hscan的count会失效
	//https://blog.csdn.net/dianxiaoer20111/article/details/120241141
	sort.StringSlice(res).Sort()
	fmt.Println("hscan res:", len(res), res)

}

func listOperate() {
	conn := GetRedis()
	defer conn.Close()

	conn.Do(del, "key")
	conn.Do(rpush, "key", "middle", "right")

	val, _ := redis.Strings(conn.Do(lrange, "key", 0, 5))
	fmt.Println("rpush", val)

	conn.Do(lpush, "key", "left")
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 5)) //lrange 中-1表示倒数第一个
	fmt.Println("lpush", val)

	conn.Do(linsert, "key", "after", "middle", "middle_r")
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 5))
	fmt.Println("linsert", val)

	val, _ = redis.Strings(conn.Do(rpop, "key"))
	fmt.Println("rpop poped", val) //并不会返回值
	val, _ = redis.Strings(conn.Do(lpop, "key"))
	fmt.Println("lpop poped", val) //并不会返回值
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 5))
	fmt.Println("rpop after pop", val)

	//lrem
	conn.Do(del, "key")
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 5))
	fmt.Println("lrem", val)
	conn.Do(rpush, "key", 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 4, 4)
	val, _ = redis.Strings(conn.Do(lrem, "key", 2, "2")) //向右删除两个
	val, _ = redis.Strings(conn.Do(lrem, "key", 0, "3")) //删除所有的
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 25))
	fmt.Println("lrem", val)
	val, _ = redis.Strings(conn.Do(lrem, "key", 1, "4"))
	val, _ = redis.Strings(conn.Do(lrem, "key", -2, "4")) //向左删除2个
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 25))
	fmt.Println("lrem", val)

	//ltrim 类似于slice中的[]
	conn.Do(del, "key")
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 5))
	conn.Do(rpush, "key", 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 4, 4)
	conn.Do(ltrim, "key", 0, 5) //只保留[0,5]下标之间的元素
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 25))
	fmt.Println("ltrim", val)

	conn.Do(del, "key")
	conn.Do(rpush, "key", 1, 1, 2, 2, 3, 3)
	val, _ = redis.Strings(conn.Do(lrange, "key", 0, 5))
	fmt.Println("lindex", val)
	val1, err := redis.String(conn.Do(lindex, "key", 3))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("lindex", val1)

	conn.Do(del, "key")
	conn.Do(rpush, "key", 1, 1, 2, 2, 3, 3)
	val2, _ := redis.Int64(conn.Do(llen, "key"))
	fmt.Println("llen", val2)

	conn.Do(del, "key")
	conn.Do(rpush, "key", 1, 1, 2, 2, 3, 3)
	val3, _ := redis.Int64(conn.Do(lset, "key", 1, 999))
	val4, err := redis.String(conn.Do(lindex, "key", 1))
	fmt.Println("lset", val3, val4)

}

func stringOperate() {
	conn := GetRedis()
	defer conn.Close()

	//set
	conn.Do(set, "key", 1)
	//get
	do, err := conn.Do(get, "key")
	if err != nil {
		return
	}
	fmt.Println("set", do, string(do.([]byte)))

	//setex Redis Setex 命令为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值。
	conn.Do(setex, "setex_key", 3, "setex_val") //3
	do, err = conn.Do(get, "setex_key")
	if err != nil {
		return
	}
	fmt.Println("setex get", do, string(do.([]byte)))
	time.Sleep(time.Second * 3)
	do, err = conn.Do(get, "setex_key")
	if err != nil {
		return
	}
	fmt.Println("setex get after 3s", do, err)

	//del
	conn.Do(del, "key")
	do2, err := conn.Do(get, "key")
	if err == redis.ErrNil {
		fmt.Println("redis.ErrNil")
	}
	if err != nil {
		return
	}
	fmt.Println("del", do2)

	//incr
	conn.Do(set, "key", "hello1")
	conn.Do(incr, "key")
	val, err := redis.String(conn.Do(get, "key"))
	if err != nil {
		return
	}
	fmt.Println("incr 对字符串，浮点不生效不生效", val)

	conn.Do(set, "key", "1")
	conn.Do(incr, "key")
	val2, err := redis.String(conn.Do(get, "key"))
	if err != nil {
		return
	}
	fmt.Println("incr 只对整型的数生效", val2)

	//decrby
	conn.Do(del, "key")
	conn.Do(set, "key", "1")
	conn.Do(decrby, "key", 10)
	val3, err := redis.String(conn.Do(get, "key"))
	if err != nil {
		return
	}
	fmt.Println("decrby 能减到负数", val3)

	//setnx
	conn.Do(del, "key")
	conn.Do(setnx, "key", "1")
	conn.Do(setnx, "key", "2")
	val4, err := redis.String(conn.Do(get, "key"))
	if err != nil {
		return
	}
	val42, _ := conn.Do(get, "key")
	fmt.Println("setnx 不会发生覆盖", val4, string(val42.([]byte)))

	//mset mget
	conn.Do(del, "key")
	conn.Do(mset, "key", "val1", "key2", "val2", "key3", "val3")
	val5, err := redis.Strings(conn.Do(mget, "key", "key2", "key3"))
	if err != nil {
		return
	}
	val52, _ := conn.Do(mget, "key", "key2", "key3")
	fmt.Println("mset mget", val5, string(val52.([]interface{})[2].([]byte)))

	//APPEND
	conn.Do(del, "key")
	conn.Do(set, "key", "1")
	conn.Do(APPEND, "key", "1234")
	val6, err := redis.String(conn.Do(get, "key"))
	if err != nil {
		return
	}
	val62, _ := conn.Do(get, "key")
	fmt.Println("append结果", val6, string(val62.([]byte)))

	//strlen
	conn.Do(del, "key")
	conn.Do(set, "key", "12345")
	val7, err := redis.Int64(conn.Do(strlen, "key"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	val72, _ := conn.Do(get, "key")
	fmt.Println("strlen", val7, string(val72.([]byte))) //strlen

	//incrbyfloat
	conn.Do(del, "key")
	conn.Do(set, "key", "3.5")
	val8, err := redis.Float64(conn.Do(incrbyfloat, "key", 1.2))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	val82, _ := conn.Do(get, "key")
	fmt.Println("incrbyfloat", val8, string(val82.([]byte)))

	//incrbyfloat
	conn.Do(del, "key")
	conn.Do(set, "key", "123456789")
	val9, err := redis.String(conn.Do(getrange, "key", 0, 4))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("getrange", val9)

	//setrange
	conn.Do(del, "key")
	conn.Do(set, "key", "123456789")
	conn.Do(setrange, "key", 3, "hello")
	val10, err := redis.String(conn.Do(get, "key"))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("setrange", val10)

	//scan
	conn.Do("flushall")
	conn.Do("hmset", "key", "field1", "val1") //scan既可以扫描出字符串，对于集合和哈希等一样能够扫描出来
	conn.Do("hmset", "key", "field2", "val2")
	conn.Do("hmset", "key", "field3", "val3")
	conn.Do("sadd", "set", "val1")
	for i := 0; i < 100; i++ {
		conn.Do(set, fmt.Sprintf("key%02d", i), "val"+strconv.Itoa(i))
		conn.Do(set, fmt.Sprintf("name%02d", i), "val"+strconv.Itoa(i))
	}
	var pos int = 0
	var res []string
	scan_(conn, &pos, 10, 30, "name*", &res) //scan的过程中会出现重复，count*n会大于总的key的个数
	sort.StringSlice(res).Sort()
	fmt.Println("scan res:", len(res), res)

}

func xscan_(i interface{}, cmd, key string, pos *int, count int, n int, match string, res *[]string) {
	if n <= 0 {
		return
	}
	conn := i.(redis.Conn)
	//第一次扫描
	val11, err := conn.Do(cmd, key, *pos, "match", match, "count", count)
	if err != nil {
		fmt.Println(err)
	}
	//在这里解析结果
	replies := val11.([]interface{})
	repla1 := replies[0].([]byte)
	repla2 := replies[1].([]interface{})
	repla1final, _ := strconv.Atoi(string(repla1))
	repla2final := make([]string, len(repla2))
	for idx, v := range repla2 {
		repla2final[idx] = string(v.([]byte))
		*res = append(*res, repla2final[idx])
	}

	*pos = repla1final
	fmt.Println(repla1final, repla2final)
	xscan_(conn, cmd, key, pos, count, n-1, match, res)
}

func scan_(i interface{}, pos *int, count int, n int, match string, res *[]string) {
	if n <= 0 {
		return
	}
	conn := i.(redis.Conn)
	//第一次扫描
	val11, err := conn.Do(scan, *pos, "match", match, "count", count)
	if err != nil {
		fmt.Println(err)
	}
	//在这里解析结果
	replies := val11.([]interface{})
	repla1 := replies[0].([]byte)
	repla2 := replies[1].([]interface{})
	repla1final, _ := strconv.Atoi(string(repla1))
	repla2final := make([]string, len(repla2))
	for idx, v := range repla2 {
		repla2final[idx] = string(v.([]byte))
		*res = append(*res, repla2final[idx])
	}

	*pos = repla1final

	fmt.Println(repla1final, repla2final)
	scan_(conn, pos, count, n-1, match, res)

}
