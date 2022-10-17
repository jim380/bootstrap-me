package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ip2location/ip2location-go"
	"github.com/jim380/bootstrap-me/cmd"
	bnet "github.com/jim380/bootstrap-me/net"
	"github.com/jim380/bootstrap-me/util"
)

type result struct {
	peers `json:"peers"`
}

type peers struct {
	Seeds []struct {
		Id      string `json:"id"`
		Address string `json:"address"`
	} `json:"seeds"`
	PersistentPeers []struct {
		Id      string `json:"id"`
		Address string `json:"address"`
	} `json:"persistent_peers"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (res result) getList() ([]string, []string) {
	var persistentPeers, seeds []string

	// TO-DO use go routine to allow parallel execution
	for _, v := range res.PersistentPeers {
		if !isReachable((v.Address)) {
			continue
		}
		persistentPeers = append(persistentPeers, v.Id+"@"+v.Address)

	}

	for _, v := range res.Seeds {
		if !isReachable((v.Address)) {
			continue
		}
		seeds = append(seeds, v.Id+"@"+v.Address)
	}

	return persistentPeers, seeds
}

func isReachable(host string) bool {
	_, err := net.DialTimeout("tcp", host, time.Duration(500)*time.Millisecond)
	if err != nil {
		fmt.Println(host + " is unreachable")
		return false
	}
	return true

}

func (res result) getGeoData() {
	db, err := ip2location.OpenDB("./IP2LOCATION-LITE-DB11.IPV6.BIN")

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, v := range res.PersistentPeers {
		ip := strings.Split(v.Address, ":")[0]
		if !util.ContainsOnlyNumbers(ip) {
			old := ip
			ip = util.DomainToIp(ip)
			fmt.Println(ip + " (" + old + ")")
		} else {
			fmt.Println(ip)
		}
		result, err := db.Get_all(ip)

		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("country_short: %s\n", result.Country_short)
		fmt.Printf("country_long: %s\n", result.Country_long)
		fmt.Printf("region: %s\n", result.Region)
		fmt.Printf("city: %s\n", result.City)
	}
}

func main() {
	var chain string
	var portScan bool
	flag.StringVar(&chain, "chain", "", "Chain to query for")
	flag.BoolVar(&portScan, "scan", false, "Scan open ports of reachable hosts")
	flag.Parse()

	cmd.ValidateChain(chain)
	url := "https://raw.githubusercontent.com/cosmos/chain-registry/master/" + chain + "/chain.json"
	result := result{}
	getJson(url, &result)

	persistentPeers, seeds := result.getList()
	fmt.Println("Persistent Peers: " + strings.Join(persistentPeers[:], ","))
	fmt.Println("")
	fmt.Println("Seeds: " + strings.Join(seeds[:], ","))

	if portScan {
		fmt.Println("Port Scanning for Persistent Peers")
		bnet.ScanOpenPorts(persistentPeers)
		fmt.Println("")
		fmt.Println("Port Scanning for Seeds")
		bnet.ScanOpenPorts(seeds)
		fmt.Println("")
		fmt.Println("----------------------------------------------------------")
	}
	// result.getGeoData()
}
