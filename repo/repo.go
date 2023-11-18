package repo

import (
	"os"
	"path"

	"github.com/Hayao0819/ayaka/conf"
	"github.com/Morganamilo/go-srcinfo"
	"github.com/spf13/viper"
)

type Repository struct {
	Config *conf.RepoConf
	Pkgs   []*Package
}

type Package struct {
	Path    string
	Srcinfo *srcinfo.Srcinfo
}

func Get() (*Repository, error) {
	repodir := viper.GetString("repodir")
	repoconfig := new(conf.RepoConf)
	repo := new(Repository)
	if err := conf.LoadRepoConfig(repodir, repoconfig); err != nil {
		return nil, err
	}
	repo.Config = repoconfig

	dirs, err := os.ReadDir(repodir)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			info, err := srcinfo.ParseFile(path.Join(repodir, dir.Name(), ".SRCINFO"))
			if err != nil {
				return nil, err
			}

			pkg := new(Package)
			pkg.Path = path.Join(repodir, dir.Name())
			pkg.Srcinfo = info
			repo.Pkgs = append(repo.Pkgs, pkg)
		}
	}

	return repo, nil

}
