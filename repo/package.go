package repo

import (
	"bytes"
	"os/exec"
	"strings"

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

func (p *Package) GetPkgFilePath() string {
	stdout := new(bytes.Buffer)
	cmd := exec.Command("makepkg", "--packagelist")
	cmd.Dir = p.Path
	cmd.Stdout = stdout
	err := cmd.Run()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(stdout.String())
}
