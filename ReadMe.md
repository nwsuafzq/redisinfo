
### 镜像上传API和版本

本程序为前端页面提供镜像上传RESTful API。
镜像上传API版本v1。

### API接口

####  GET /metrics
##### 获取redis全部信息

#### GET /metrics/:info
##### 获取redis指定信息
###### 参数包含
**Server** 获取Server信息

**Clients** 获取Clients信息

**Memory** 获取Memory信息

**Persistence** 获取Persistence信息

**Stats** 获取Stats信息

**Replication** 获取Replication信息

**CPU** 获取CPU信息

**Cluster** 获取Cluster信息

**Keyspace** 获取Keyspace信息


**docker run -p <宿主端口号>:10012 <镜像名称>:<版本Tag> --server "10012" --auth "" ip1,ip2...ipn**

**注：--server ":10012" 为服务的端口号，不填默认为10012**

**--auth "" 为redis密码**

**ip1 为redis得IP地址，可填多个,传入IP时应带端口号如"10.1.234.36:30323"，若不填择端口号为redis默认的6443端口**



