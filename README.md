## gofun包含的功能
gofun是一个集成了常用开发组件的web框架，直接集成了好用且关注度很高的组件，只做了初始化封装，希望能够给开发者节省大量搭建基础服务的工作，快速切入业务开发。主要使用了以下开源组件

* [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
* [github.com/gorilla/websocket](https://github.com/gorilla/websocket)
* [github.com/go-redis/redis](https://github.com/go-redis/redis)
* [github.com/asaskevich/govalidator](https://github.com/asaskevich/govalidator)
* [github.com/aliyun/aliyun-oss-go-sdk/oss](https://github.com/aliyun/aliyun-oss-go-sdk/oss)
* [github.com/minio/minio-go](https://github.com/minio/minio-go)
* [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
* [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

#### 基本功能
- [X] 登录/注册(含参数校验)
- [X] log
- [X] yaml配置文件
- [ ] 发送邮件
- [ ] 优雅启停

#### web请求路由
- [X] gin封装
- [X] mvc+service分层架构

`control`层调用逻辑

[![g3rfoV.png](https://z3.ax1x.com/2021/05/07/g3rfoV.png)](https://imgtu.com/i/g3rfoV)


#### 中间件
- [X] 用户认证(session/jwt)
- [ ] traceId



#### 数据库
- [X] Mysql
- [ ] Redis
- [ ] mongodb
- [ ] elasticsearch

#### OSS对象存储
- [X] 阿里云OSS
- [X] Minio

#### 消息中间件
- [ ] Emqx
- [ ] Rabbitmq

#### 其他
- [ ] swagger 接口文档
- [X] 可并发调用的webSocket
- [ ] grpc
- [ ] restful API 返回规范
- [ ] 常用工具函数(uuid)


#### 高可用
- [ ] 熔断
- [ ] 限流
- [ ] 热升级

#### 应用部署
- [ ] 部署脚本
- [ ] docker
- [ ] kubernetes
