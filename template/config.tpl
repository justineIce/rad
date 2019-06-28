package config

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	global *Config
	once   sync.Once
)

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig(path string) error {
	c, err := ParseConfig(path)
	if err != nil {
		return err
	}
	global = c
	return nil
}

// GetGlobalConfig 获取全局配置
func GetGlobalConfig() *Config {
	if global == nil {
		return &Config{}
	}
	return global
}

// ParseConfig 解析配置文件
func ParseConfig(path string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Config 配置参数
type Config struct {
	Log   Log   `toml:"log"`
	Gorm  Gorm  `toml:"gorm"`
	MySQL MySQL `toml:"mysql"`
	HTTP  HTTP  `toml:"http"`
}

// HTTP http配置参数
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// Log 日志配置参数
type Log struct {
	Level    int    `toml:"level"`
	Path     string `toml:"path"`
	MaxDays  int    `toml:"maxdays"`
	Separate string `toml:"separate"`
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool   `toml:"debug"`
	DBType       string `toml:"db_type"`
	MaxLifetime  int    `toml:"max_lifetime"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	TablePrefix  string `toml:"table_prefix"`
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}
