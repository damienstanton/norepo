package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/damienstanton/norepo/interfaces/intf"
)

func main() {
	intf.Decorate(http.DefaultClient,
		intf.Authorization("blahblah"),
		intf.Logging(log.New(os.Stdout, "client:", log.LstdFlags)),
		intf.FaultTolerance(5, time.Second))
	log.Println("🤖")
}
