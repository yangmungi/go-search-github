package main

import (
	"flag"
	"log"

	"github.com/yangmungi/go-search-github/pkg/repo"
)

func main() {
	url := flag.String("url", "https://github.com/git/git.git", "")
	flag.Parse()

	//timeoutFlag := flag.Int("timeout", 5, "timeout in seconds")

	//bctx := context.TODO()

	//ctx, cancel := context.WithTimeout(bctx, time.Duration(timeout)*time.Second)
	//defer cancel()

	r := new(repo.Repo)
	err := r.CloneAndAnalyze(*url)
	if err != nil {
		log.Println(err)
	}
}
