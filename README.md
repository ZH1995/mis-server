# MyGin

MyGin 是一个基于 Gin 框架的示例 API 项目，旨在展示如何构建一个简单的 RESTful API，同时集成 Swagger 文档生成。

## 特性

- 使用 Gin 框架构建高性能的 HTTP API
- 集成 Swagger 生成 API 文档
- 支持用户登录功能
- 使用 GORM 进行数据库操作

## 目录结构

```
MyGin/
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

## 安装

1. 克隆项目到本地：
   ```bash
   git clone https://github.com/yourusername/MyGin.git
   cd MyGin
    ```

2. 安装依赖包:
   ```bash
   go mod tidy
   ```

## 使用

1. 启动服务器：
    ```bash
    go run cmd/server/main.go
    ```

2. 访问 API 文档： 打开浏览器并访问 
    http://localhost:8080/swagger/index.html


## 贡献

欢迎任何形式的贡献！请提交问题或拉取请求。

## 许可证

本项目采用 MIT 许可证，详情请参见 LICENSE 文件