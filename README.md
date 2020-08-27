# shortURL
短网址生成网站。golang+mysql。  
使用数据库自增id，通过base62进制转换生成短网址。  
增加ip黑名单，限制单日访问次数，防止爆破。  
LRU缓存机制，保证同一长网址转换成同一短网址。  
## 在线网站
    
## 项目目录
> main.go
> request.go
> transcode.go
> ipBlackList.go
> cache.go  
> shortURL.html

## 后续
* 访问次数统计
* 分布式存储，减轻服务器压力

