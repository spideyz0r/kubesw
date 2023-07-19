package common

import (
	"os"
	"testing"
)

func TestDetectShell(t *testing.T) {
	testCases := []struct {
		shell    string
		expected string
	}{
		{
			shell:    "bash",
			expected: "/bin/bash",
		},
		{
			shell:    "zsh",
			expected: "/bin/zsh",
		},
		{
			shell:    "",
			expected: "Failed to detect shell",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.shell, func(t *testing.T) {
			os.Setenv("SHELL", tc.expected)
			actualShell, _ := detect_shell()
			if actualShell != tc.expected {
				t.Errorf("Expected shell to be %s, but got %s", tc.shell, actualShell)
			}
			os.Unsetenv("SHELL")
		})
	}
}

func TestReadRc(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		shell    string
		expected string
	}{
		{
			name:     "/.bashrc",
			input:    "Test content",
			shell:    "bash",
			expected: "\nTest content",
		},
		{
			name:     "/.zshrc",
			input:    "Test content",
			shell:    "zsh",
			expected: "\nTest content",
		},
		{
			name:     "/.bash_logout",
			input:    "Test content",
			shell:    "bash",
			expected: "\nTest content",
		},
		{
			name:     "/.nonexistent",
			input:    "Anything",
			shell:    "bash",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			homedir := "/tmp"
			os.Setenv("HOME", "/tmp")
			file_path := homedir + tc.name
			file, err := os.Create(file_path)
			if err != nil {
				t.Fatal(err)
			}
			file.WriteString(tc.input)
			file.Close()

			actual_rc := read_rc(tc.shell)
			if actual_rc != tc.expected {
				t.Errorf("Expected shell to be >>%s<<, but got >>%s<<", tc.expected, actual_rc)
			}
			os.Remove(file_path)
		})
	}
}

func TestInitialSetup(t *testing.T) {
	testCases := []struct {
		homeEnvVal       string
		kubeconfigEnvVal string
		expectedFiles    string
		expectedDir      string
	}{
		{
			homeEnvVal:       "/tmp",
			kubeconfigEnvVal: "/tmp/.kube",
			expectedFiles:    "/tmp/.kube",
			expectedDir:      "/tmp/.kube/config.kubesw.d",
		},
	}

	for _, tc := range testCases {
		os.Setenv("HOME", tc.homeEnvVal)
		os.Setenv("KUBECONFIG", tc.kubeconfigEnvVal)
		os.MkdirAll(tc.kubeconfigEnvVal, os.FileMode(0755))

		actualFiles, actualDir := InitialSetup()
		os.Remove(tc.kubeconfigEnvVal)
		os.Remove(tc.expectedDir)

		// Check if the actual results match the expected results
		if actualFiles != tc.expectedFiles {
			t.Errorf("Expected files: %s, but got: %s", tc.expectedFiles, actualFiles)
		}
		if actualDir != tc.expectedDir {
			t.Errorf("Expected dir: %s, but got: %s", tc.expectedDir, actualDir)
		}
	}
}
