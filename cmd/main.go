package main

import (
	"fmt"
	"log"
	"time"

	"github.com/k-vanio/protobuf/basic"
)

type debug struct{}

func (d *debug) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(debug))

	basic.WriterSample()
	basic.ReadSample()
}
