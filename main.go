package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
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
	var reachablePersistentPeers, reachableSeeds []string
	fmt.Println("********************************************************************")
	fmt.Println("              Checking Reachability of Persistent Peers             ")
	fmt.Println("********************************************************************")
	for _, v := range res.PersistentPeers {
		reachablePersistentPeers = append(reachablePersistentPeers, v.Id+"@"+v.Address)
	}

	persistentPeers = bnet.CheckReachability(reachablePersistentPeers)

	fmt.Println("********************************************************************")
	fmt.Println("                   Checking Reachability of Seeds                   ")
	fmt.Println("********************************************************************")
	for _, v := range res.Seeds {
		reachableSeeds = append(reachableSeeds, v.Id+"@"+v.Address)
	}

	seeds = bnet.CheckReachability(reachableSeeds)

	return persistentPeers, seeds
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
	var portScan, help bool
	flag.StringVar(&chain, "chain", "", "Chain to query for")
	flag.BoolVar(&portScan, "scan", false, "Scan open ports of reachable hosts")
	flag.BoolVar(&help, "help", false, "Show all supported chains")
	flag.Parse()

	if help {
		fmt.Println("List of supported chains:")
		fmt.Println("\n" + strings.Join(cmd.GetAllChains(), " "))
		os.Exit(0)
	}

	cmd.ValidateChain(chain)
	url := "https://raw.githubusercontent.com/cosmos/chain-registry/master/" + chain + "/chain.json"
	result := result{}
	getJson(url, &result)
	fmt.Println("Running for " + chain)
	persistentPeers, seeds := result.getList()
	fmt.Println("********************************************************************")
	fmt.Println("                     Reachable Persistent Peers                     ")
	fmt.Println("********************************************************************")
	fmt.Println(strings.Join(persistentPeers[:], ","))
	fmt.Println("********************************************************************")
	fmt.Println("                          Reachable Seeds                           ")
	fmt.Println("********************************************************************")
	fmt.Println(strings.Join(seeds[:], ","))

	if portScan {
		fmt.Println("********************************************************************")
		fmt.Println("                           Port Scanning                           ")
		fmt.Println("********************************************************************")
		fmt.Println(">> Scanning Reachable Persistent Peers")
		bnet.ScanOpenPorts(persistentPeers)
		fmt.Printf("\n")
		fmt.Println(">> Scanning Reachable Seeds")
		bnet.ScanOpenPorts(seeds)
		fmt.Printf("\n")
		fmt.Println("----------------------------------------------------------")
	}
	// result.getGeoData()
}
