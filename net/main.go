package net

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

// adjust this as you like but
// the higher it is, the more unreliable it becomes
const workers = 100

// the number of ports to scan
// 5 digits = 10^5 = 100,000
const workUnits = 100000

func worker(address string, ports, results chan int) {
	for p := range ports {
		conn, err := net.DialTimeout("tcp", address+":"+strconv.Itoa(p), time.Duration(500)*time.Millisecond)
		if err != nil {
			// send 0 if port is closed
			results <- 0
			continue
		}
		conn.Close()
		results <- p
		fmt.Println("Port " + strconv.Itoa(p) + " is open on " + address)
	}
}

func reachabilityWorker(host string, results chan string) {
	address := strings.Split(host, "@")[1]
	conn, err := net.DialTimeout("tcp", address, time.Duration(500)*time.Millisecond)
	if err != nil {
		fmt.Println(host + " is unreachable")
		// send "unreachable" if port is closed
		results <- "unreachable"
		return
	}
	conn.Close()
	results <- host
}

func CheckReachability(hosts []string) []string {
	// a dedicated thread to pass results back to main
	results := make(chan string, len(hosts))

	var reachableHosts []string

	// send work to the workers
	for _, v := range hosts {
		go reachabilityWorker(v, results)
	}

	// gather results
	for i := 0; i < len(hosts); i++ {
		host := <-results
		if host != "unreachable" {
			reachableHosts = append(reachableHosts, host)
		}
	}

	// close the thread
	close(results)

	return reachableHosts
}

func ScanOpenPorts(addressList []string) {
	var addresses []string
	for _, v := range addressList {
		addresses = append(addresses, strings.Split(strings.Split(v, "@")[1], ":")[0])
	}

	// one thread to work units to the worker
	ports := make(chan int, workers)
	// a dedicated thread to pass results back to main
	results := make(chan int)

	var openports []int

	// send work to the workers
	for _, v := range addresses {
		for j := 0; j < cap(ports); j++ {
			go worker(v, ports, results)
		}
	}

	// send to the workers in a separate goroutine
	go func() {
		for i := 1; i <= workUnits; i++ {
			ports <- i
		}
	}()

	// this loop starts before more work (than workUnits) can continue
	for i := 0; i < workUnits; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	// close the threads
	close(ports)
	close(results)

	// sort the open ports
	sort.Ints(openports)

	if len(openports) == 0 {
		fmt.Println("no open ports found")
	}
}
