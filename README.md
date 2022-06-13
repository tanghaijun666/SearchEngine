# simple-demo

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./SimpleTikTok
```

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

### 5.26 谢腾
将功能抽象了出来，分为

common:用于数据库的连接以及放置 约定的Response结构体

controller:负责每一个接口的请求转发

dao：实现与数据库的交互

model :放置gorm与数据的映射结构体

public ：暂时放置静态资源

route ：组成接口的group

service ：将dao封装成服务


增加了注册与登录功能，原来的表因为id都为string，现在将其改为了int，并且取消了外键
尽量在应用层面实现数据关联，并且没有加上comment，做这块接口的同学自己对照创建一下

后续需要开发新的接口的时候，直接在controller上写自己的handle即可，然后对应增加dao与service

## 6.5 谢腾

在dao层增加了jwt颁发token的功能（本来想放到common但是会出现import cycle not allowed的问题）

后续可以使用JwtAuth去将获取的token解析为userid和username

##   akun
点赞和点赞列表和评论的数据库文件