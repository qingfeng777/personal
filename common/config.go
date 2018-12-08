package common

import (
	"github.com/BurntSushi/toml"
)

type Configs struct {
	App *App
	Log *Log
	DB  *DB
}

// LogConfig
type Log struct {
	Level string
}

// StatsdConfig Statsdç
type App struct {
	GoMaxPro    int
	Port        uint
	TimeOut     uint8
	ReleaseMode bool
	Prefix      string
}

//db
type DB struct {
	Url     string
	MaxIdle uint
	MaxOpen uint
}

// 全局配置信息
var ConfigRef *Configs

// InitConfig 实例化配置
func InitConfig(path string) {
	config, err := loadConfig(path)
	if err != nil {
		panic(err)
	}
	ConfigRef = config
}

func loadConfig(path string) (*Configs, error) {
	config := new(Configs)
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, err
	}

	return config, nil
}
