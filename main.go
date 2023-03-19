package main

import (
	"log"
	"os"

	"github.com/nathankjer/gin-proxy/db"
	"github.com/nathankjer/gin-proxy/routers"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	if os.Getenv("PROXY_HOST") != "" {
		err := db.Connect()
		if err == nil {
			g.Go(func() error {
				r := routers.InitRouterRequests()
				return r.Run(":3000")
			})
			g.Go(func() error {
				r := routers.InitRouterGateway()
				return r.Run(":3001")
			})
			if err := g.Wait(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal("PROXY_HOST environment variable is not set. Run 'export PROXY_HOST=example.com'.")
	}
}
