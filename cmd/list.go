package cmd

import (
	"os"
	"path"

	"github.com/Hayao0819/ayaka/conf"
	srcinfo "github.com/Morganamilo/go-srcinfo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func listCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "List packages in repository",
		Long:  "List packages in repository",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			repodir := viper.GetString("repodir")
			repoconfig := new(conf.RepoConf)
			if err := conf.LoadRepoConfig(repodir, repoconfig); err != nil {
				return err
			}

			dirs, err := os.ReadDir(repodir)
			if err != nil {
				return err
			}
			for _, dir := range dirs {
				if dir.IsDir() {
					info, _ := srcinfo.ParseFile(path.Join(repodir, dir.Name(), ".SRCINFO"))
					cmd.Println(info.Pkgbase)
				}
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	subCmds = append(subCmds, listCmd())
}
