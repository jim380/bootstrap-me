package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/jim380/bootstrap-me/cmd"
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

func main() {
	var chain string
	flag.StringVar(&chain, "chain", "", "Chain to query for")
	flag.Parse()

	cmd.ValidateChain(chain)
	url := "https://raw.githubusercontent.com/cosmos/chain-registry/master/" + chain + "/chain.json"
	result := result{}
	getJson(url, &result)

	persistentPeers, seeds := result.getList()

	fmt.Println("Persistent Peers: " + strings.Join(persistentPeers[:], ","))
	fmt.Println("")
	fmt.Println("Seeds: " + strings.Join(seeds[:], ","))
}
