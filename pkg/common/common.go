package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"k8s.io/client-go/tools/clientcmd"
)

var debug bool

func SetDebug(d bool) {
	debug = d
}

func UpdateNamespace(kube_config, namespace string) {
	kubeconfigBytes, err := ioutil.ReadFile(kube_config)
	if err != nil {
		fmt.Printf("Failed to read kubeconfig file: %v", err)
		os.Exit(1)
	}

	config, err := clientcmd.Load(kubeconfigBytes)
	if err != nil {
		fmt.Printf("Failed to load kubeconfig %s: %v", kube_config, err)
		os.Exit(1)
	}

	config.Contexts[config.CurrentContext].Namespace = namespace

	err = clientcmd.WriteToFile(*config, kube_config)
	if err != nil {
		fmt.Printf("Failed to write kubeconfig file %s: %v", kube_config, err)
		os.Exit(1)
	}
	if debug {
		fmt.Printf("Updated namespace to %s\n", namespace)
	}
}

func GetCurrent(resource string) string {
	var cmd *exec.Cmd
	if debug {
		fmt.Printf("Getting current %s\n", resource)
	}
	if resource == "context" {
		cmd = exec.Command("kubectl", "config", "current-context")
	} else {
		cmd = exec.Command("kubectl", "config", "view", "--minify", "--flatten", "--output", "jsonpath={.contexts[?(@.name==\""+GetCurrent("context")+"\")].context.namespace}")
	}
	current, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(current))
}

func UpdateContext(kubeconfig_kubesw_dir, context, namespace string) string {
	if debug {
		fmt.Printf("Creating new kubeconfig file for context %s\n", context)
	}
	context_file := fmt.Sprintf("%s/%s-%s.yaml", kubeconfig_kubesw_dir, context, namespace)
	fileMode := os.FileMode(0600)
	file, err := os.OpenFile(context_file, os.O_CREATE|os.O_WRONLY, fileMode)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = file.Chmod(fileMode)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("kubectl", "config", "view", "--minify", "--flatten", "--context", context)
	new_ctx_config, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(new_ctx_config)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		fmt.Printf("Wrote %s\n", context_file)
	}
	return context_file
}

func InitialSetup() (string, string) {
	if debug {
		fmt.Printf("Performing initial setup\n")
	}
	homedir := os.Getenv("HOME")
	kubeconfig_orig := os.Getenv("KUBECONFIG")
	kubeconfig_kubesw_dir := fmt.Sprintf("%s/.kube/config.kubesw.d", homedir)

	initial_directories := []string{kubeconfig_kubesw_dir}
	init_dir_permission := os.FileMode(0755)
	for _, dir := range initial_directories {
		err := os.MkdirAll(dir, init_dir_permission)
		if err != nil {
			log.Fatal(err)
		}
	}

	dir, err := os.Open(kubeconfig_kubesw_dir)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	var existing_files []string
	for _, file := range files {
		existing_files = append(existing_files, fmt.Sprintf("%s/%s", kubeconfig_kubesw_dir, file.Name()))
	}
	existing_files = append(existing_files, kubeconfig_orig)
	return strings.Join(existing_files, ":"), kubeconfig_kubesw_dir
}

func read_rc(shell string) string {
	if debug {
		fmt.Printf("Reading %s rc files\n", shell)
	}
	zdotdir := os.Getenv("ZDOTDIR")
	homedir := os.Getenv("HOME")
	if zdotdir == "" && shell == "zsh" {
		zdotdir = homedir
	}

	rc_shell := make(map[string][]string)
	rc_shell["zsh"] = []string{
		zdotdir + "/.zshrc",
		zdotdir + "/.zprofile",
		zdotdir + "/.zlogin",
		zdotdir + "/.zlogout",
	}
	rc_shell["bash"] = []string{
		homedir + "/.bashrc",
		homedir + "/.bash_profile",
		homedir + "/.profile",
		homedir + "/.bash_login",
		homedir + "/.bash_logout",
	}

	var all_rc_files string
	for _, rc_file := range rc_shell[shell] {
		_, err := os.Stat(rc_file)
		if os.IsNotExist(err) {
			if debug {
				fmt.Printf("File %s does not exist, skipping...\n", rc_file)
			}
			continue
		}

		if debug {
			fmt.Printf("Reading %s\n", rc_file)
		}
		file, err := os.Open(rc_file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		all_rc_files = fmt.Sprintf("%s\n%s", all_rc_files, string(content))
	}
	return all_rc_files
}

func detect_shell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", fmt.Errorf("Failed to detect shell")
	}
	return shell, nil
}

func SpawnShell(kube_config, history string) {
	shell, err := detect_shell()
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		fmt.Printf("Spawning shell: %s\n", shell)
	}
	switch shell {
	case "/bin/bash":
		spawn_generic_shell(kube_config, history, "bash")
	case "/bin/zsh":
		spawn_generic_shell(kube_config, history, "zsh")
	default:
		log.Fatal("Unsupported shell")
	}
}

func InjectShellHistory(option, value string) string {
	return fmt.Sprintf("kubesw set %s %s", option, value)
}

func spawn_generic_shell(kube_config, history, shell string) {
	// to keep history in bash the user needs to put the following to their bashrc
	// shopt -s histappend
	// PROMPT_COMMAND="history -a"

	// to keep history in zsh the user needs to enable setopt inc_append_history
	// https://zsh.sourceforge.io/Doc/Release/Options.html#index-SHARE_005fHISTORY

	current_rc := read_rc(shell)

	var extra_rc_configuration string
	var file *os.File
	var tmp_rc_path string
	var err error
	if shell == "bash" {
		file, err = ioutil.TempFile("", "kubesw_rc")
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			fmt.Printf("Created temporary rc file: %s\n", file.Name())
		}
		tmp_rc_path = file.Name()
		defer os.Remove(tmp_rc_path)
		extra_rc_configuration = `
		[[ -f "$HOME/.bash_profile" ]] && source "$HOME/.bash_profile"
		[[ -f "$HOME/.bash_login" ]]  && source "$HOME/.bash_login"
		[[ -f "$HOME/.profile" ]] && source "$HOME/.profile"
		export KUBECONFIG=` + kube_config + `:$KUBECONFIG
		# shopt -s histappend
		# PROMPT_COMMAND="history -a; history -n"
		history -n
		history -s ` + history + `
		# export PS1="[\u@\h \W: $(go run main.go get context) @ $(go run main.go get namespace)]\\$ "
		`
	} else {
		extra_rc_configuration = `
		[[ -f /etc/zshenv ]] && source "/etc/zshenv"
		[[ -f /etc/zsh/zshenv ]] && source "/etc/zsh/zshenv"
		[[ -f "$HOME/.zshenv" ]] && source "$HOME/.zshenv"
		export KUBECONFIG=` + kube_config + `:$KUBECONFIG
		# export PS1="[\u@\h \W: $(go run main.go get context) @ $(go run main.go get namespace)]\\$ "
		`
		tmp_rc_path, err = ioutil.TempDir("", "zdir")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(tmp_rc_path)
		file, err = os.OpenFile(tmp_rc_path+"/.zshrc", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	rc := fmt.Sprintf("%s\n%s", current_rc, extra_rc_configuration)
	if debug {
		fmt.Printf("Writing to rc file: %s:\n", tmp_rc_path)
		fmt.Printf("%s\n", rc)
	}
	_, err = file.WriteString(rc)
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		fmt.Printf("Spawning %s shell with rcfile: %s\n", shell, tmp_rc_path)
	}

	var cmd *exec.Cmd
	var env []string
	if shell == "bash" {
		cmd = exec.Command("/bin/bash")
		cmd.Args = []string{"/bin/bash", "--rcfile", tmp_rc_path}
	} else {
		cmd = exec.Command("/bin/zsh")
		env = os.Environ()
		env = append(env, "ZDOTDIR="+tmp_rc_path)
		cmd.Env = env
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start bash shell: %v", err)
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Failed to wait for bash shell: %v", err)
		log.Fatal(err)
	}
	if debug {
		fmt.Printf("Shell session closed.")
		fmt.Printf("%s", rc)
	}
}

func ListContexts() {
	if debug {
		fmt.Printf("Listing contexts\n")
	}
	cmd := exec.Command("kubectl", "config", "get-contexts", "--no-headers=true", "-o", "name")
	contexts, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.TrimSpace(string(contexts)))
}

func ListNamespaces() {
	if debug {
		fmt.Printf("Listing namespaces\n")
	}
	cmd := exec.Command("kubectl", "get", "namespaces", "--no-headers=true", "-o", "name")
	raw_namespaces, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	var namespaces []string
	lines := strings.Split(strings.TrimSpace(string(raw_namespaces)), "\n")
	for _, line := range lines {
		namespaces = append(namespaces, strings.Split(line, "/")[1])
	}
	fmt.Println(strings.Join(namespaces, "\n"))
}
