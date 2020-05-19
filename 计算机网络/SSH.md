# 加密

## 对称加密

client与server端都有一个相同的公有秘钥，安全性高，但是如果client有很多个，公有秘钥安全性不高.

## 非对称加密

* 非对称加密的过程：
	* server端有**公钥**和**私钥**
	* server将公钥发送给client端
	* client端利用公钥加密密码发送给server端
	* server端利用私钥解密得到密码。


SSH就是使用这种方式来进行认证的。但是仍然容易出现中间人攻击即**一个人冒充server端的整个验证过程**。CA认证可以有效避免中间人攻击。

> https://www.cnblogs.com/diffx/p/9553587.html

* SSH登录的过程
	1. Client将自己的公钥存放在Server上，追加在文件authorized_keys中。
	2. Server端接收到Client的连接请求后，会在authorized_keys中匹配到Client的公钥pubKey，并生成随机数R，用Client的公钥对该随机数进行加密得到pubKey(R)
	，然后将加密后信息发送给Client。
	3. Client端通过私钥进行解密得到随机数R，然后对随机数R和本次会话的SessionKey利用MD5生成摘要Digest1，发送给Server端。
	4. Server端会也会对R和SessionKey利用同样摘要算法生成Digest2。
	5. Server端会最后比较Digest1和Digest2是否相同，完成认证过程。