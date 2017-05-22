package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/FlyCynomys/cockroachdb/bench"
)

func main() {
	/*for i := 10; i < 200; i = i + 2 {
		bench.Step1(i, i+1)
	}*/

	bench.Step2()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}
