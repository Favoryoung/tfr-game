package conf

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"log"
	"os"
	"strings"
)

var Conf *koanf.Koanf

func init() {
	configPath := getEnvJson()
	Conf = koanf.New(".")
	f := file.Provider(configPath)
	err := Conf.Load(f, json.Parser())
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	hotUpdateAllConf()

	//配置热更新
	if HotUpdate {
		go func() {
			f.Watch(func(event interface{}, err error) {
				if err != nil {
					log.Printf("配置文件【%s】 监听异常: %v", configPath, err)
					return
				}

				log.Printf("****************** 配置文件【%s】 热更新 ******************", configPath)
				newC := koanf.New(".")
				newC.Load(f, json.Parser())
				Conf = newC

				hotUpdateAllConf()
			})
		}()
	}
}

func hotUpdateAllConf() {
	for _, confFunc := range setConfFunc {
		confFunc()
	}
}

func StringOr(path string, def string) string {
	v := Conf.String(path)
	if v == "" {
		return def
	}
	return v
}

func BoolOrFalse(path string) bool {
	return Conf.Bool(path)
}

func IntOr(path string, def int) int {
	v := Conf.Int(path)
	if v == 0 {
		return def
	}
	return v
}

func getEnvJson() string {
	return GetOsArgs("env_json=", "conf.json")
}

func GetOsArgs(key string, def string) string {
	if key == "" {
		return ""
	}
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, key) {
			return strings.TrimPrefix(arg, key)
		}
	}

	return def
}
