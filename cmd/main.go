package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/v55/github"
	"github.com/yangmungi/go-search-github/pkg/repo"
)

func main() {
	sizeFlag := flag.Int("size", 100, "size >")
	timeoutFlag := flag.Int("timeout", 5, "timeout in seconds")

	bctx := context.TODO()

	do(bctx, *sizeFlag, *timeoutFlag)
}

func do(bctx context.Context, size, timeout int) error {
	ctx, cancel := context.WithTimeout(bctx, time.Duration(timeout)*time.Second)
	defer cancel()

	cli := github.NewClient(nil)
	search := fmt.Sprintf("language:Go size:>%d", size)

	log.Printf("search:%s", search)

	result, _, err := cli.Search.Repositories(ctx, search, &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	})

	if result == nil && err == nil {
		err = fmt.Errorf("missing result; no err")
	}

	if err != nil {
		return err
	}

	log.Printf("results:%d incomplete:%v", result.Total, *result.IncompleteResults)

	r := repo.Repo
	for _, r := range result.Repositories {
		url := r.GitURL
		if url == nil {
			continue
		}
		err := r.CloneAndAnalyze(*url)

		if err != nil {
			log.Println(err)
		}
	}

	return nil
}