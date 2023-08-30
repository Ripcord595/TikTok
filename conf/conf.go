package conf

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Mysql 结构定义了 MySQL 数据库的配置。
type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

// Redis 结构定义了 Redis 数据库的配置。
type Redis struct {
	IP       string
	Port     int
	Database int
}

// Server 结构定义了服务器的配置。
type Server struct {
	IP   string
	Port int
}

// Path 结构定义了文件路径的配置。
type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

// Config 结构定义了整个应用程序的配置。
type Config struct {
	DB     Mysql  `toml:"mysql"`
	RDB    Redis  `toml:"redis"`
	Server Server `toml:"server"`
	Path   Path   `toml:"path"`
}

// NewConfig 返回一个新的配置实例，可以用于初始化配置。
func NewConfig() Config {
	return Config{
		DB: Mysql{
			Host:      "localhost",
			Port:      3306,
			Database:  "douyin",
			Username:  "root",
			Password:  "12345678",
			Charset:   "utf8mb4",
			ParseTime: true,
			Loc:       "Local",
		},
		RDB: Redis{
			IP:       "localhost",
			Port:     6379,
			Database: 0,
		},
		Server: Server{
			IP:   "0.0.0.0",
			Port: 8080,
		},
		Path: Path{
			FfmpegPath:       "ffmpeg",
			StaticSourcePath: "./static_source",
		},
	}
}

// EnsurePathValid 检查配置中的路径是否有效，如果无效则进行修复。
func (c *Config) EnsurePathValid() {
	var err error
	if _, err = os.Stat(c.Path.StaticSourcePath); os.IsNotExist(err) {
		if err = os.Mkdir(c.Path.StaticSourcePath, 0755); err != nil {
			log.Fatalf("mkdir error: path %s", c.Path.StaticSourcePath)
		}
	}
	if _, err = os.Stat(c.Path.FfmpegPath); os.IsNotExist(err) {
		if _, err = exec.Command("ffmpeg", "-version").Output(); err != nil {
			log.Fatalf("ffmpeg not valid %s", c.Path.FfmpegPath)
		} else {
			c.Path.FfmpegPath = "ffmpeg"
		}
	} else {
		c.Path.FfmpegPath, err = filepath.Abs(c.Path.FfmpegPath)
		if err != nil {
			log.Fatalln("get abs path failed:", c.Path.FfmpegPath)
		}
	}
	c.Path.StaticSourcePath, err = filepath.Abs(c.Path.StaticSourcePath)
	if err != nil {
		log.Fatalln("get abs path failed:", c.Path.StaticSourcePath)
	}
}

// DBConnectString 构建并返回 MySQL 数据库连接字符串。
func DBConnectString(db *Mysql) string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		db.Username, db.Password, db.Host, db.Port, db.Database,
		db.Charset, db.ParseTime, db.Loc)
	log.Println(arg)
	return arg
}
