package main

import (
    "github.com/yahaa/gofoer/config"
    "github.com/yahaa/gofoer/fileser"
    "os"
    "os/signal"
    "log"
)

func main() {
    config.InitLog("[Gofoer]")
    config.InitConfig()
    go fileser.InitFileServer(config.G.FilePath, config.G.FileServerAddress)

    c := make(chan os.Signal)
    signal.Notify(c)
    s := <-c
    log.Printf("Exist signal %v", s)
}
