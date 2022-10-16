package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
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

func (res result) getList() (string, string) {
	var firstP, restP, firstS, restS string

	for i, v := range res.PersistentPeers {
		if i == 0 {
			firstP = v.Id + "@" + v.Address
		}
		restP = restP + "," + v.Id + "@" + v.Address
	}

	for i, v := range res.Seeds {
		if i == 0 {
			firstS = v.Id + "@" + v.Address
		}
		restS = restS + "," + v.Id + "@" + v.Address
	}
	return (firstP + restP), (firstS + restS)
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

	fmt.Println("Seeds: " + seeds)
	fmt.Println("")
	fmt.Println("Persistent Peers: " + persistentPeers)
}
