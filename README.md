Some magic go projects.

## 各个项目介绍
### [dServer](https://github.com/wmillers/kindleWeatherClock)
Bilibili 直播弹幕格式化服务端，本项目使用 gin 作为 api 路由转发，通过 Websocket 连接B站服务器，获取直播信息（弹幕、礼物等），使用 grpc 与接收端通信。

### hueSwitch
局域网控制智能家居。

### webShop
基于 gin 框架，使用 Viper 读取参数，JWT 提供认证，结合 gorm 对 Mysql 数据库操作，使用 Redis 缓存用户购物车信息，kafka 管理消息队列，Nginx 添加 HTTPS 支持 ，并通过 docker 进行部署。