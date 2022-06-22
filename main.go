package main

import (
	"goapi/comm/inits"
	"goapi/comm/log"
	"goapi/routers"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Infof("system begin")
	if err := inits.Init(); err != nil {
		log.Errorf("inits failed, err:%v", err)
		return
	}
	log.Infof("inits.Init Succ")

	var g errgroup.Group

	// 外部服务
	g.Go(func() error {
		r := routers.Init()
		if err := r.Run(":8080"); err != nil {
			log.Error("startup service failed, err:%v", err)
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error(err)
	}
}
