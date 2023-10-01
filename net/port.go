package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("website", "scanme.nmpa.org", "url to scanner")

func main() {
	// revisando los puertos en este rango, para ver cuales estan abiertos
	// informar que le puerto esta abierto (open, close)
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open \n", i)
		}(i)
	}
	wg.Wait()

}
