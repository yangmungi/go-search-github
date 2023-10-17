package shgit

import (
	"log"
	"os"
	"os/exec"

	"github.com/go-git/go-billy/v5"
)

type ShellGit struct {
	WorkDir string
	GitPath string
}

func (s *ShellGit) Clone(url string) (billy.Filesystem, error) {
	oldDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	defer func() {
		err := os.Chdir(oldDir)
		if err != nil {
			log.Println(err)
		}
	}()

	cmd := exec.Command(s.GitPath, "clone", "--depth=1", "--single-branch", url)
	if err = cmd.Run(); err != nil {
		return nil, err
	}

	// =
	return nil, nil
}

func (s *ShellGit) Clean(fs billy.Filesystem) {
	// TODO maybe delete the directory
}
