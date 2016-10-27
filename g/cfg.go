package g

import (
	"encoding/json"
	"fmt"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type ListenConfig struct {
	Protocol string `json:"protocol"`
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
}

type DatabaseConfig struct {
	Dsn     string `json:"dsn"`
	MaxIdle int    `json:"max_idle"`
}

type GlobalConfig struct {
	LogLevel        string          `json:"log_level"`
	RefreshInterval int             `json:"refresh_interval"`
	Listen          *ListenConfig   `json:"listen"`
	Database        *DatabaseConfig `json:"database"`
	TarballDir      string          `json:"tarballDir"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("config file %s is nonexistent", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read config file %s fail %s", cfg, err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse config file %s fail %s", cfg, err)
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
	return nil
}
