# Gin101

## 项目简介

本项目希望带领你学习Gin框架，快速实现一个简单的Web应用。

## 目录结构

```
Gin101/
├── api/                    # API 接口定义
│   └── v1/                # API 版本控制
│       ├── article.go
│       ├── auth.go
│       ├── exchange_rate.go
│       └── like.go
├── cmd/                    # 主要应用程序入口
│   └── server/            # 服务器应用
│       └── main.go
├── config/                 # 配置文件和配置加载
│   ├── config.go
│   ├── config.yml
│   ├── db.go
│   └── redis.go
├── internal/              # 私有应用程序和库代码
│   ├── middleware/        # 中间件
│   │   └── auth.go
│   ├── model/            # 数据模型
│   │   ├── article.go
│   │   ├── exchange_rate.go
│   │   └── user.go
│   ├── repository/       # 数据库操作层
│   │   └── mysql/
│   ├── service/         # 业务逻辑层
│   │   ├── article.go
│   │   └── auth.go
│   └── pkg/             # 内部共享包
│       ├── constants/
│       ├── errors/
│       └── utils/
├── pkg/                 # 可以被外部应用程序使用的库代码
│   └── util/
├── scripts/            # 构建、安装、分析等操作的脚本
├── test/              # 测试文件
├── web/               # Web 静态资源
│   ├── static/
│   └── template/
├── .gitignore
├── go.mod
├── go.sum
├── Makefile           # 项目管理命令
└── README.md
```

## 贡献指南

欢迎大家的PR，一起学习，一起进步！

## 许可证

MIT