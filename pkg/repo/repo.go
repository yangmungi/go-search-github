package repo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Repo struct {
	Progress sideband.Progress
}

type ProgressPrintf struct{}

func (p *ProgressPrintf) Write(b []byte) (int, error) {
	fmt.Printf("%s", string(b))
	return len(b), nil
}

func (r *Repo) CloneAndAnalyze(url string) error {
	log.Printf("clone %s", url)

	fs := memfs.New()
	s := memory.NewStorage()

	_, err := git.Clone(s, fs, &git.CloneOptions{
		URL: url,

		ReferenceName: plumbing.HEAD,
		SingleBranch:  true,
		Depth:         1,

		Progress: r.Progress,
	})

	if err != nil {
		return err
	}

	return Recurse(fs, ".", func(filename string) {
		if !strings.HasSuffix(filename, ".go") {
			return
		}

		fset := token.NewFileSet()
		bf, err := fs.Open(filename)
		if err != nil {
			log.Println(err)
			return
		}

		f, err := parser.ParseFile(fset, filename, bf, 0)
		if err != nil {
			log.Println(err)
			return
		}

		for _, decl := range f.Decls {
			switch af := decl.(type) {
			case *ast.GenDecl:
				for _, spec := range af.Specs {
					switch aspec := spec.(type) {
					case *ast.TypeSpec:
						log.Printf("> %s %T", aspec.Name.String(), aspec.Type)

						// TODO handle nested *ast.StructType
						if aspec.TypeParams == nil {
							continue
						}

						for _, f := range aspec.TypeParams.List {
							log.Printf("> %v %v", f.Names, f.Type)

						}
					case *ast.ImportSpec:
					default:
						log.Printf("> %T", aspec)
					}
				}
			}
		}

	})
}

func Recurse(fs billy.Filesystem, cur string, w func(filename string)) error {
	fis, err := fs.ReadDir(cur)
	for _, fi := range fis {
		filename := fi.Name()
		fullPath := fs.Join(cur, filename)
		log.Printf("%s %v", fullPath, fi.IsDir())

		if fi.IsDir() {
			err = Recurse(fs, fullPath, w)
			if err != nil {
				return err
			}

			continue
		}

		w(filename)
	}

	return nil
}
