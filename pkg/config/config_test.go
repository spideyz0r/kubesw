package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestEnvironment() {
	os.Setenv("HOME", "/tmp")
}

func TestReadConfiguration(t *testing.T) {
	setupTestEnvironment()

	config := ReadConfiguration()
	expectedConfig := Config{
		Shell: "auto",
		PS1:   "",
		Rc_files: map[string][]string{
			"zsh": []string{
				"/.zshrc",
				"/.zprofile",
				"/.zlogin",
				"/.zlogout",
			},
			"bash": []string{
				"/.bashrc",
				"/.bash_profile",
				"/.profile",
				"/.bash_login",
				"/.bash_logout",
			},
		},
	}
	assert.Equal(t, expectedConfig.Shell, config.Shell)
	assert.Equal(t, expectedConfig.PS1, config.PS1)
	assert.Equal(t, expectedConfig.Rc_files, config.Rc_files)
}

func TestReadConfiguration_ConfigFileNotFound(t *testing.T) {
	config := ReadConfiguration()
	expectedConfig := Config{
		Shell: "auto",
		PS1:   "",
		Rc_files: map[string][]string{
			"zsh": []string{
				"/.zshrc",
				"/.zprofile",
				"/.zlogin",
				"/.zlogout",
			},
			"bash": []string{
				"/.bashrc",
				"/.bash_profile",
				"/.profile",
				"/.bash_login",
				"/.bash_logout",
			},
		},
	}
	assert.Equal(t, expectedConfig.Shell, config.Shell)
	assert.Equal(t, expectedConfig.PS1, config.PS1)
	assert.Equal(t, expectedConfig.Rc_files, config.Rc_files)
}
