package Proxy

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

type Command struct {
	shell   string
	command string
	args    []string
}

func StartProxy() {

}

func CreateTables() error {
	// create new table
	var com Command
	com.shell = "/bin/bash"
	com.command = "iptables"
	table := fmt.Sprintf("-N %s", "PROXY")
	// ignore local host
	log.Print(table)
	return nil
}

func IgnoreLocal() error {
	var com Command
	com.shell = "/bin/bash"
	com.command = "iptables"
	com.args = append(com.args, "")
	return nil
}

func SetRedirect(port int) {
	var com Command
	com.shell = "/bin/bash"
	com.command = "iptables"
	portStr := fmt.Sprintf("-m tcp --to-port %d", port)
	com.args = append(com.args, "-t nat", "-I POSTROUTING", "-p tcp", portStr, "-j REDIRECT")
	_, _ = execute(com)
}

func FilterSelf() {
	var com Command
	com.shell = "/bin/bash"
	com.command = "iptables"
	com.args = append(com.args, "--dest 127.0.0.1/8")
}

func execute(com Command) (string, error) {
	var args bytes.Buffer
	for _, arg := range com.args {
		args.WriteByte(' ')
		args.WriteString(arg)
	}
	out, err := exec.Command(com.shell, "-c", com.command, args.String()).CombinedOutput()
	if err != nil {
		log.Fatalf("execute command failed, out: %s, err: %v", string(out), err)
		return "", err
	}
	return string(out), nil
}

func GetRandomPort() int {
	rand.NewSource(time.Now().Unix())
	for true {
		port := rand.Intn(1000) + 1000
		cmdStr := fmt.Sprintf("netstat -ano -p tcp | findstr %d", port)
		out, err := exec.Command("/bin/bash", "-c", cmdStr).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		if string(out) != "" {
			continue
		}
		return port
	}
	return 1
}
