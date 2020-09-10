* 只是一个简要的过程
  * client端
    * connect(sock, &inet_addr, sizeof(inet_addr))
    * send(sock, &var, sizeof(var))
    * recv(sock, &var, sizeof(var))
  * server端
    * bind(sock, &inet_addr, sizeof(inet_addr))
    * listen(sock, int n)
    * rsock = accept(sock, NULL, NULL)
    * send()
    * recv()