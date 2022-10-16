package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var chainList []string

func ValidateChain(chain string) {
	chainList = GetAllChains()
	if chain == "" {
		log.Fatal("Chain was not provided.")
	} else {
		valid := false
		for _, v := range chainList {
			if v == chain {
				valid = true
			}
		}
		if !valid {
			log.Fatal(fmt.Sprintf("\n%s is not a supported chain.\n", chain) + fmt.Sprint("\nList of supported chains: ", chainList))
		}
	}
}

func GetAllChains() []string {
	resp, err := http.Get("https://github.com/cosmos/chain-registry")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "flex-auto min-width-0 col-md-2 mr-3"
		}
		return false
	}
	// grab all folders and print them
	foldersRaw := scrape.FindAll(root, matcher)
	var folders []string
	for _, folder := range foldersRaw {
		if strings.Contains(scrape.Text(folder), ".") || strings.Contains(scrape.Text(folder), "_") {
			continue
		}
		folders = append(folders, scrape.Text(folder))
	}
	return folders
}
