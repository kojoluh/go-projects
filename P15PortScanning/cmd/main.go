package main

import (
	"fmt"

	"github.com/kojoluh/go-secruity-samples/port"
)

func main() {
	fmt.Println("Port Scanner in Go!")
	// open := port.ScanPort("tcp", "localhost", 1313)

	// fmt.Printf("Port Open: %t\n", open)

	// results := port.InitialScan("localhost")
	// fmt.Println(results)

	// for res := range results {
	// 	fmt.Printf("Port-State: %v\n", res)
	// }

	port.WhitePortScan("localhost")

}
