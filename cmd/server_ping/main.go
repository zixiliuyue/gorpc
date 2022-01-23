package main

import (
	"fmt"
	"gorpc/common/blog"
	"gorpc/scene_server/tmserver/service"
	"net/http"

	// "time"
	restful "github.com/emicklei/go-restful"
	"gorpc/common/rpc"
	_ "gorpc/scene_server/tmserver/core/command"
)

func main() {
	sync := make(chan int)
	go func(addrCha chan int) {
		coreService := service.New("127.0.0.1", 7878)
		coreService.SetConfig()
		c := restful.NewContainer().Add(coreService.WebService())
		server := &http.Server{
			Addr:    "0.0.0.0:7878",
			Handler: c,
		}
		addrCha <- 1
		if err := server.ListenAndServe(); err != nil {
			blog.Fatalf("listen and serve failed, err: %v", err)
		}
	}(sync)
	a := <-sync
	fmt.Println(a)
	// time.Sleep(10 * time.Second)
	for i := 0; i < 3; i++ {
		cli, err := rpc.DialHTTPPath("tcp", "127.0.0.1:7878", "/rpc")
		if err != nil {
			fmt.Println(err)
			return
		}
		m, err := cli.Ping()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(m.String(), 1111)
	}

}
