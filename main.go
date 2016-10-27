package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/toolkits/logger"

	"github.com/coraldane/ops-meta/db"
	"github.com/coraldane/ops-meta/g"
	"github.com/coraldane/ops-meta/http"
	"github.com/coraldane/ops-meta/models"
	"github.com/coraldane/ops-meta/store"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if err := g.ParseConfig(*cfg); err != nil {
		log.Fatalln(err)
	}
	logger.SetLevelWithDefault(g.Config().LogLevel, "info")

	models.Init()
	db.Init()

	go store.RefreshDesiredAgent()
	go store.CleanStaleHost()

	addr := g.Config().Listen.Addr
	if "" == addr {
		addr = "127.0.0.1"
	}

	if 0 >= g.Config().Listen.Port {
		return
	}

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = addr
	beego.BConfig.Listen.HTTPPort = g.Config().Listen.Port

	http.Start()
}
