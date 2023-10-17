package helper

import (
	"github.com/yangmungi/go-search-github/pkg/gitimpl"
	"github.com/yangmungi/go-search-github/pkg/gitimpl/gogit"
	"github.com/yangmungi/go-search-github/pkg/gitimpl/shgit"
)

type Implementation string

const (
	Shell Implementation = "shell"
	GoGit                = "go-git"
)

func Get(i Implementation) Git {
	var gitImpl gitimpl.Git

	if i == Shell {
		gitImpl = &gogit.GoGit{}
	} else {
		gitImpl = &shgit.ShellGit{}
	}

	return gitImpl
}
