package main

import (
	"fmt"
	"gorpc/common/blog"
	"gorpc/scene_server/tmserver/service"
	"net/http"

	"encoding/json"
	restful "github.com/emicklei/go-restful"
	"gorpc/common/rpc"
	"gorpc/common/types"
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
	<-sync
	cli, err := rpc.DialHTTPPath("tcp", "127.0.0.1:7878", "/rpc")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := &types.OPFindOperation{}
	msg.OPCode = types.OPCountCode
	reply := types.OPReply{}
	err = cli.Call(service.CommandRDBOperation, msg, &reply)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", reply)
	d, _ := json.Marshal(reply)
	fmt.Println(string(d))
}
