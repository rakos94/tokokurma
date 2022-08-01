package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"tokokurma/config/db"

	"github.com/spf13/viper"
)

// Config is struct all of configuration
type Config struct {
	APP struct {
		ENV string `mapstructure:"env"`
	} `mapstructure:"app"`
	Database db.DatabaseList
}

var configuration Config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	var err error

	ex, err := os.Getwd()
	log.Println(ex)
	log.Println(basepath)

	fmt.Printf("Reading Config %s\n", basepath+"/db")
	viper.AddConfigPath(basepath + "/db")
	viper.SetConfigType("yml")
	viper.SetConfigName("database.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	viper.Unmarshal(&configuration)

	fmt.Println("============================")
	data, _ := json.Marshal(configuration)
	fmt.Println(string(data))
	fmt.Println("============================")

}

// GetConfig get config
func GetConfig() *Config {
	return &configuration
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
