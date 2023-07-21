package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Shell    string
	PS1      string
	Rc_files map[string][]string
}

func ReadConfiguration() Config {
	viper.SetDefault("shell", "auto")
	viper.SetDefault("PS1", "")

	rc_shell := make(map[string][]string)

	rc_shell["zsh"] = []string{
		"/.zshrc",
		"/.zprofile",
		"/.zlogin",
		"/.zlogout",
	}
	rc_shell["bash"] = []string{
		"/.bashrc",
		"/.bash_profile",
		"/.profile",
		"/.bash_login",
		"/.bash_logout",
	}
	viper.SetDefault("default_rc", rc_shell)

	homedir := os.Getenv("HOME")
	viper.AddConfigPath(homedir + "/.kubesw")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; using defaults
		} else {
			panic(err)
		}
	}

	return Config{
		Shell:    viper.GetString("shell"),
		PS1:      viper.GetString("PS1"),
		Rc_files: viper.GetStringMapStringSlice("default_rc"),
	}
}
