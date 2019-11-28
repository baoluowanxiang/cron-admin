package main

import (
	"cron-admin/base"
	"cron-admin/option"
	"cron-admin/services/http"
	"cron-admin/services/tcp"
	"log"
)

func main() {
	opts := option.Options{}
	httpSrv := &http.HttpServer{}
	tcpSrv := &tcp.TcpServer{}
	startService(opts, httpSrv, tcpSrv)
	select{}
}

func startService(opts option.Options, services ...base.Server) {
	for _, srv := range services {
		go func() {
			err := srv.Start(opts)
			if err != nil {
				log.Print(err)
			}
		}()
	}
}