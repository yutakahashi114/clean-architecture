package main

import (
	"github.com/yutakahashi114/clean-architecture/controller/rest"
)

func main() {
	rest.Serve()
	// grpc.Serve()
}
