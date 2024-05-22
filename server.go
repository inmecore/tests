package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"tests/router"
	"time"
)

type server interface {
	ListenAndServe() error
}

type Network struct{}

func (n *Network) initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 30 * time.Minute
	s.WriteTimeout = 30 * time.Minute
	s.MaxHeaderBytes = 1 << 20
	return s
}

func (n *Network) Run() {
	r := router.New()
	s := n.initServer(":3000", r)
	if e := s.ListenAndServe(); e != nil {
		panic(e)
	}
}
