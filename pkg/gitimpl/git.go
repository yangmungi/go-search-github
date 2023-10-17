package gitimpl

import "github.com/go-git/go-billy/v5"

type Git interface {
	Clone(url string) (billy.Filesystem, error)

	Clean(fs billy.Filesystem)
}
