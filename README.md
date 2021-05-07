## gofun包含的功能
gofun是一个集成了常用开发组件的web框架，直接集成了好用且关注度很高的组件，只做了初始化封装，希望能够给开发者节省大量搭建基础服务的工作，快速切入业务开发

#### 基本功能
- [X] 登录/注册(含参数校验)
- [X] log
- [X] yaml配置文件
- [ ] 发送邮件

#### web请求路由
- [X] gin封装
- [X] mvc+service分层架构

`control`层调用逻辑
[![g3rfoV.png](https://z3.ax1x.com/2021/05/07/g3rfoV.png)](https://imgtu.com/i/g3rfoV)


#### 用户认证
- [X] session功能
- [X] jwt功能



#### 数据库
- [ ] Mysql
- [ ] Redis
- [ ] mongodb

#### OSS对象存储
- [X] 阿里云OSS
- [X] Minio

#### 消息中间件
- [ ] Emqx
- [ ] Rabbitmq

#### 文档
- [ ] swagger

#### 其他
- [X] 可并发调用的webSocket
- [ ] grpc
- [ ] elasticsearch
- [ ] restful API 返回规范
- [ ] 加解密


#### 高可用
- [ ] 熔断
- [ ] 限流
- [ ] 热升级

#### 应用部署
- [ ] docker
- [ ] kubernetes