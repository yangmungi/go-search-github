package main

import (
	"flag"
	"log"

	"github.com/yangmungi/go-search-github/pkg/gitimpl"
	"github.com/yangmungi/go-search-github/pkg/gitimpl/gogit"
	"github.com/yangmungi/go-search-github/pkg/gitimpl/shgit"
	"github.com/yangmungi/go-search-github/pkg/repo"
)

func main() {
	url := flag.String("url", "https://github.com/git/git.git", "")

	impl := flag.String("impl", "go-git", "go-git or shell, defaults to go-git")
	tmpDir := flag.String("temp-dir", "/tmp/git", "used with impl=shell")
	gitPath := flag.STring("git-path", "/usr/bin/git", "used with impl=shell")

	flag.Parse()

	var gitImpl gitimpl.Git
	if *impl == "shell" {
		gitImpl = &gogit.GoGit{}
	} else {
		gitImpl = &shgit.ShellGit{}
	}

	r := &repo.Repo{
		Git: gitImpl,
	}

	err := r.CloneAndAnalyze(*url)

	if err != nil {
		log.Println(err)
	}
}
