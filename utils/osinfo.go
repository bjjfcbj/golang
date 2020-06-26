package main

import (
    "os"
    "runtime"
)

func main() {
    println(`系统类型：`, runtime.GOOS)

    println(`系统架构：`, runtime.GOARCH)

    println(`CPU 核数：`, runtime.GOMAXPROCS(0))

    name, err := os.Hostname()
    if err != nil {
        panic(err)
    }
    println(`电脑名称：`, name)
}
