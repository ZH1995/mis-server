// 文件配置
package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server        server        `yaml:"server"`
	Db            db            `yaml:"db"`
	Redis         redis         `yaml:"redis"`
	ImageSettings imageSettings `yaml:"imageSettings"`
	Log           log           `yaml:"log"`
}

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

// 图片上传配置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

// 日志配置
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

var Config *config

func init() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
