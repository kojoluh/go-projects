package port

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

type ScanMessage struct {
	TcpResult []ScanResult
	UdpResult []ScanResult
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}

	defer conn.Close()

	result.State = "Open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult
	ports := 9999

	for i := 1; i <= ports; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}

func BroadPortScan(protocol string, hostname string, ch chan<- *ScanMessage, wg *sync.WaitGroup) {
	var results []ScanResult
	ports := 65535
	for i := 1; i <= ports; i++ {
		result := ScanPort(protocol, hostname, i)
		if result.State == "Open" {
			results = append(results, result)
		}
		// results = append(results, ScanPort(protocol, hostname, i))
	}

	if protocol == "tcp" {
		ch <- &ScanMessage{
			TcpResult: results,
		}
	} else {
		ch <- &ScanMessage{
			UdpResult: results,
		}
	}
	wg.Done()
}

func WhitePortScan(hostname string) {
	wg := &sync.WaitGroup{}
	msgChan := make(chan *ScanMessage, 2)
	wg.Add(2)
	now := time.Now()
	go BroadPortScan("tcp", "localhost", msgChan, wg)
	go BroadPortScan("udp", "localhost", msgChan, wg)

	wg.Wait()
	close(msgChan)

	for msg := range msgChan {
		if len(msg.TcpResult) > 0 {
			log.Printf("Ports: TCP length: %v\n\n", len(msg.TcpResult))
		} else {
			log.Printf("Ports: UDP length: %v\n\n", len(msg.UdpResult))
		}

	}
	log.Println(time.Since(now))

}
