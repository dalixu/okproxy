package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"
)

var config = &struct {
	Addr  string
	Path  string
	Proxy string
	SumDb string
}{}

func init() {
	flag.StringVar(&config.Addr, "addr", "127.0.0.1:8899", "ip:port")
	flag.StringVar(&config.Path, "path", "./mod/cache", "PATH the root of the caches")
	flag.StringVar(&config.Proxy, "proxy", "https://goproxy.cn,direct", "GOPROXY just like what Go is doing right now")
	flag.StringVar(&config.SumDb, "sumdb", "off", "GOSUMDB just like what Go is doing right now")
}

func main() {
	flag.Parse()
	fmt.Println(config)
	proxy := goproxy.New()
	//proxy.ErrorLogger = &log.Logger{}
	proxy.Cacher = &cacher.Disk{Root: config.Path}
	proxy.GoBinEnv = append(proxy.GoBinEnv, "GOPROXY="+config.Proxy)
	proxy.GoBinEnv = append(proxy.GoBinEnv, "GOSUMDB="+config.SumDb)
	err := http.ListenAndServe(config.Addr, proxy)
	fmt.Println(err)
}
