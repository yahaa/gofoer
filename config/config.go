package config

import (
    "flag"
    "log"
    "encoding/json"
    "io/ioutil"
)

var G Config

type Config struct {
    FilePath          string
    FileServerAddress string

}

func InitConfig() {
    log.Println("Start to read config ...")
    var cfgPath string
    flag.StringVar(&cfgPath, "c", "none", "This is the config path for this application")
    flag.Parse()
    if cfgPath == "" {
        log.Fatalf("Param error .")
        return
    }

    var global Config

    data, err := ReadFile(cfgPath)
    if err != nil {
        log.Fatalf("Init config error %v", err)
        return
    }

    err = json.Unmarshal(data, &global)
    if err != nil {
        log.Fatalf("Json data not available.")
        return
    }
    G = global
    log.Printf("Init global config success.")

}

func InitLog(prefix string) {

    log.SetPrefix(prefix + " ")
}

// ReadFile 读取配置文件
func ReadFile(filename string) ([]byte, error) {
    data, err := ioutil.ReadFile(filename)
    return data, err
}
