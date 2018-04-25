package fileser

import (
    "net/http"
    "log"
)

func InitFileServer(fPath string, address string) {
    fSer := http.NewServeMux()
    fSer.Handle("/", http.FileServer(http.Dir(fPath)))
    log.Fatal(http.ListenAndServe(address, fSer))
}
