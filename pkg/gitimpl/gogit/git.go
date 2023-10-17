package gogit

import (
	"fmt"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/storage/memory"
)

type ProgressPrintf struct{}

func (p *ProgressPrintf) Write(b []byte) (int, error) {
	fmt.Printf("%s", string(b))
	return len(b), nil
}

type GoGit struct {
	Progress sideband.Progress
}

func (g *GoGit) Clone(url string) (billy.Filesystem, error) {
	fs := memfs.New()
	s := memory.NewStorage()
	_, err := git.Clone(s, fs, &git.CloneOptions{
		URL: url,

		ReferenceName: plumbing.HEAD,
		SingleBranch:  true,
		Depth:         1,

		Progress: g.Progress,
	})

	if err != nil {
		return nil, err
	}

	return fs, nil
}
