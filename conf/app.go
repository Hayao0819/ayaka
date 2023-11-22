package conf

import (
	"os"

	viper "github.com/spf13/viper"
)

type AppConf struct {
	RepoDir string `mapstructure:"repodir"`
	DistDir string `mapstructure:"distdir"`
	
}

var AppConfigPath string
//var RepoDir string
var AppConfig AppConf

func LoadAppConfig(config *AppConf) error {
	if AppConfigPath != "" {
		viper.SetConfigFile(AppConfigPath)
	} else {

		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		viper.SetConfigName(".ayakarc")
		viper.AddConfigPath(homedir)
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	return nil
}
