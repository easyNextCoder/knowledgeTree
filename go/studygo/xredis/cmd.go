package xredis

const (
	set         = "set"
	setex       = "setex"
	get         = "get"
	del         = "del"
	incr        = "incr"
	decr        = "decr"
	incrby      = "incrby"
	decrby      = "decrby"
	setnx       = "setnx"
	mset        = "mset"
	mget        = "mget"
	APPEND      = "APPEND"
	strlen      = "strlen"
	incrbyfloat = "incrbyfloat" //可以处理负值
	getrange    = "getrange"    //getrange key [startIndex, endIndex] //获得对应key的值的指定区间的字符串
	setrange    = "setrange"
	scan        = "scan" //迭代的是当前数据库中的所有数据库键,默认的count是10，可以扫出所有的key

	rpush   = "rpush"
	lpush   = "lpush"
	linsert = "linsert"
	rinsert = "rinsert"
	lrange  = "lrange"
	lpop    = "lpop"
	rpop    = "rpop"
	lrem    = "lrem"
	ltrim   = "ltrim"
	lindex  = "lindex"
	llen    = "llen"
	lset    = "lset"

	hget    = "hget"
	hset    = "hset"
	hgetall = "hgetall" //!谨慎使用,可能阻塞生产环境
	hexists = "hexists"
	hlen    = "hlen"
	hmset   = "hmset"
	hmget   = "hmget"
	hvals   = "hvals"
	hkeys   = "hkeys"
	hincrby = "hincrby"
	hdel    = "hdel"  //hdel 是删除key中的field，del是删除key
	hscan   = "hscan" //用于迭代hash类型中的键值对，第一个参数是一个数据库键

	sadd        = "sadd"        //可以一次add多个
	scard       = "scard"       //计算集合的大小
	sismember   = "sismember"   //某个key是不是集合中的元素
	srandmember = "srandmember" //随机获取一些集合中的元素
	smembers    = "smembers"    //获取集合中的全部元素，会发生阻塞，无序的 !谨慎使用,可能阻塞生产环境
	sdiff       = "sdiff"       //计算两个集合之间的差集
	sinter      = "sinter"      //计算两个集合的交集
	sunion      = "sunion"      //计算两个集合的并集
	spop        = "spop"        //随机弹出一个集合中的元素
	srem        = "srem"        //从集合中移除元素，可以一次操作多个
	sscan       = "sscan"       //用于迭代set集合中的元素，第一个参数是一个数据库键

	zadd             = "zadd"
	zrange           = "zrange"
	zrevrange        = "zrevrange"
	zrangebyscore    = "zrangebyscore"
	zrevrangebyscore = "zrevrangebyscore"
	zscore           = "zscore"
	zrank            = "zrank"
	zrevrank         = "zrevrank"
	zcount           = "zcount"
	zinterstore      = "zinterstore"
	//ZINTERSTORE 时间复杂度: O(N*K)+O(M*log(M)) 这里N表示有序集合中成员数最少的数字，K表示有序集合数量。M表示结果集中重合的数量。
	//交集并集的复杂度很高，如果有bigkey的情况会严重阻塞主线程，建议一般不要使用，可以把两个zset的元素取出来，在内存中进行交并集运算这样不会阻塞redis主线程
	zunionstore      = "zinterstore"
	zincrby          = "zincrby"
	zcard            = "zcard"
	zrem             = "zrem"
	zremrangebyrank  = "zremrangebyrank"
	zremrangebyscore = "zremrangebyscore"
	zpopmax          = "zpopmax" //移除有序集合中分数最大的count个元素
	zpopmin          = "zpopmin" //移除有序组合中分数最小的count个元素
	zscan            = "zscan"   //用于迭代sortset集合中的元素和元素对应的分值，第一个参数是一个数据库键，

	/*
		keys * 打印出所有的key
		keys he*
		keys he[h-l] 第三个字母是h到l的范围
		keys he?  问号表示任意一位
		//key相关命令在生产服严禁使用，因为生产服的key很多会把redis阻塞挂掉，复杂度为O(n)
		dbsize 计算key的总数时间复杂读O(1)
		expire name 3 设置3秒过期
		ttl name  查看name还有多长时间过期
		persist name 去掉name的过期时间
		type name 查看name的类型

		//其他命令
		info 内存，cpu和主从相关
		client list 正在连接的绘画
		client kill ip:端口
		flushall 清空所有
		flushdb 只清空当前库
		select 数字 选择某个库，总共16个库
		monitor 记录操作日志
	*/
)
