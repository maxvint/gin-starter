package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuwenhui/gopkg/logs"
)

func main() {
	defer logs.Flush()

	var (
		flagAddress = flag.String("a", ":9000", "http address")
	)
	flag.Parse()

	ln, err := net.Listen("tcp", *flagAddress)
	if err != nil {
		logs.Fatalf("create listener failed: %s", *flagAddress)
		logs.Flush()
		return
	}

	logs.Infof("listening on %s", ln.Addr())

	r := gin.Default()

	CreateRouter(r)

	// wait for server running
	logs.Errorf("server running With error: %v", http.Serve(ln, r))
}
