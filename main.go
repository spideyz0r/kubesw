package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pborman/getopt"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	help := getopt.BoolLong("help", 'h', "display this help")
	debug := getopt.BoolLong("debug", 'd', "debug mode")
	context := getopt.StringLong("context", 'c', "", "set a context. use -l to list all")
	namespace := getopt.StringLong("namespace", 'n', "", "set a namespace. use -l to list all")

	getopt.Parse()

	if *help {
		getopt.Usage()
		os.Exit(0)
	}

	new_kube_config_path, kubeconfig_kubesw_dir := initial_setup()
	if *debug {
		fmt.Printf("KUBECONFIG: %s\n", new_kube_config_path)
	}

	if *context == "-l" {
		list_contexts()
		os.Exit(0)
	}
	if *context == "-u" {
		fmt.Printf("%s\n", get_current_context())
		os.Exit(0)
	}

	if *namespace == "-l" {
		list_namespaces()
		os.Exit(0)
	}
	if *namespace == "-u" {
		fmt.Printf("%s\n", get_current_namespace())
		os.Exit(0)
	}

	if *context == "" {
		*context = get_current_context()
	}

	if *namespace == "" {
		*namespace = "default"
	}

	kube_config := update_context(kubeconfig_kubesw_dir, *context, *namespace)
	update_namespace(kube_config, *namespace)
	spawn_shell(kube_config)
	fmt.Printf("Exited kubesw shell running context %s\n", *context)
}

func update_namespace(kube_config, namespace string) {
	kubeconfigBytes, err := ioutil.ReadFile(kube_config)
	if err != nil {
		fmt.Printf("Failed to read kubeconfig file: %v", err)
		os.Exit(1)
	}

	config, err := clientcmd.Load(kubeconfigBytes)
	if err != nil {
		fmt.Printf("Failed to load kubeconfig: %v", err)
		os.Exit(1)
	}

	config.Contexts[config.CurrentContext].Namespace = namespace

	err = clientcmd.WriteToFile(*config, kube_config)
	if err != nil {
		fmt.Printf("Failed to write kubeconfig file: %v", err)
		os.Exit(1)
	}
	// fmt.Printf("Updated namespace to %s\n", namespace)
}

func get_current_namespace() string {
	cmd := exec.Command("kubectl", "config", "view", "--minify", "--flatten", "--output", "jsonpath={.contexts[?(@.name==\""+get_current_context()+"\")].context.namespace}")
	current_namespace, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(current_namespace))
}

func get_current_context() string {
	cmd := exec.Command("kubectl", "config", "current-context")
	current_context, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(current_context))
}

func update_context(kubeconfig_kubesw_dir, context, namespace string) string {
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
	// fmt.Printf("Wrote %s\n", context_file)
	return context_file
}

func initial_setup() (string, string) {
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

func read_bashrc() string {
	homedir := os.Getenv("HOME")
	var all_rc_files string
	rc_files := []string{
		homedir + "/.bashrc",
		homedir + "/.bash_profile",
		homedir + "/.profile",
		homedir + "/.bash_login",
		homedir + "/.bash_logout",
	}

	for _, rc_file := range rc_files {
		_, err := os.Stat(rc_file)
		if os.IsNotExist(err) {
			// fmt.Printf("File %s does not exist, skipping...\n", rc_file)
			continue
		}

		// fmt.Printf("Reading %s\n", rc_file)
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

func spawn_shell(kube_config string) {
	current_rc := read_bashrc()
	tmp_rc, err := ioutil.TempFile("", "kubesw_rc")
	if err != nil {
		log.Fatal(err)
	}
	tmp_rc_path := tmp_rc.Name()
	defer os.Remove(tmp_rc_path)

	extra_rc_configuration := `
	[[ -f "$HOME/.bash_profile" ]] && source "$HOME/.bash_profile"
	[[ -f "$HOME/.bash_login" ]]  && source "$HOME/.bash_login"
	[[ -f "$HOME/.profile" ]] && source "$HOME/.profile"
	export KUBECONFIG=` + kube_config + `:$KUBECONFIG
	export PS1="[\u@\h \W: $(go run main.go -c -u) @ $(go run main.go -n -u)]\\$ "
	`

	rc := fmt.Sprintf("%s\n%s", current_rc, extra_rc_configuration)
	_, err = tmp_rc.WriteString(rc)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("bash")
	cmd.Args = []string{"bash", "--rcfile", tmp_rc_path}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func list_contexts() {
	cmd := exec.Command("kubectl", "config", "get-contexts", "--no-headers=true", "-o", "name")
	contexts, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.TrimSpace(string(contexts)))
}

func list_namespaces() {
	cmd := exec.Command("kubectl", "get", "namespaces", "--no-headers=true", "-o", "name")
	namespaces, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.TrimSpace(string(namespaces)))
}

// TODO
// Add verbose mode and remove all prints
// Read configuration like rc files and PS1 from configuration file
// Add support for other shells
// Split namespaces list
// Improve parsing
// Improve error messages
// Add tests
// Add release, test, rpmbuilder pipelines