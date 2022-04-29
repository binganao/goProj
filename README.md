Some magic go projects.

## 各个项目介绍
### [dServer](https://github.com/wmillers/kindleWeatherClock)
Bilibili 直播信息整合服务端，支持 gRPC 和 RESTful 通信，使用 gin 作为 api 路由转发，获取和分析直播信息（弹幕、礼物、房间数据等），logrus 记录日志。

### hueSwitch
局域网控制智能家居。

### webStore
基于 gin 框架，使用 Viper 读取参数，JWT 提供认证，结合 gorm 对 Mysql 数据库操作，使用 Redis 缓存用户收藏信息，Nginx 添加 HTTPS 支持，并通过 docker 进行部署。
