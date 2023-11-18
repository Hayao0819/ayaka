package repo

import (
	"github.com/Hayao0819/ayaka/builder"
	"github.com/Morganamilo/go-srcinfo"
)

type Package struct {
	Path    string
	Srcinfo *srcinfo.Srcinfo
}

func (p *Package) Build(method string, target builder.Target) error {
	builder := builder.GetBuilder(method)
	return builder.Build(p.Path, &target)
}
