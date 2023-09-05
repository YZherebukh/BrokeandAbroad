package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/BrokeandAbroad/one/client/web"
	"github.com/BrokeandAbroad/one/internal"
)

func main() {
	c := web.New(http.Client{})

	fmt.Printf("number of bytes %d \n", internal.Run(c, rand.Intn(100)))
}
