package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/sunjiangjun/etherscan-cli/api"
	"github.com/sunjiangjun/etherscan-cli/config"
	"github.com/sunjiangjun/xlog"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "The system file of config")
	flag.Parse()
	if len(configPath) < 1 {
		panic("can not find config file")
	}
	cfg := config.LoadConfig(configPath)
	log.Printf("%+v\n", cfg)

	x := xlog.NewXLogger().BuildOutType(xlog.FILE).BuildLevel(xlog.Level(4)).BuildFormatter(xlog.FORMAT_JSON).BuildFile("./log/err", 24*time.Hour)
	txWriter := xlog.NewXLogger().BuildOutType(xlog.FILE).BuildLevel(xlog.Level(4)).BuildFormatter(xlog.FORMAT_JSON).BuildFile("./data/tx", 24*time.Hour)

	ch := make(chan int64)

	var latestBlock int64

	go func() {
		for {
			var key string
			if l := len(cfg.Key); l == 1 {
				key = cfg.Key[0]
			} else {
				key = cfg.Key[rand.Intn(l)]
			}

			r, err := api.GetBlock(time.Now().Unix(), key)
			if err != nil {
				<-time.After(3 * time.Second)
				continue
			}

			i, err := strconv.ParseInt(r.Result, 0, 64)
			if err != nil {
				<-time.After(2 * time.Second)
				continue
			}

			if i <= latestBlock {
				<-time.After(2 * time.Second)
				continue
			}

			latestBlock = i
			ch <- latestBlock
			<-time.After(5 * time.Second)
		}
	}()

	go func() {
		for {

			blockNumber := <-ch

			var key string
			if l := len(cfg.Key); l == 1 {
				key = cfg.Key[0]
			} else {
				key = cfg.Key[rand.Intn(l)]
			}

			addrList := cfg.Address

			for _, a := range addrList {

				r, err := api.GetTx(a, blockNumber, blockNumber, key)
				if err != nil {
					x.Errorf("address:%v,blockNumber:%v,err:%v", a, blockNumber, err.Error())
					continue
				}

				for _, v := range r.Result {
					bs, _ := json.Marshal(v)
					txWriter.Printf("%v", string(bs))
				}
			}

			log.Printf("scan block complete,blockNumber:%v", blockNumber)

		}

	}()

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	t, c := context.WithTimeout(context.Background(), 2*time.Second)
	defer c()
	<-t.Done()

}
