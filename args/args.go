package main

import (
	"bufio"
	"flag"
	"fmt"
	"os/exec"
)

var command = flag.String("e", "whoami", "执行命令")
var dir = flag.String("d", "", "执行目录")

func main() {
	flag.Parse()

	// *command = "mvn clean"
	*dir = "/home/hxp/source-code/stvd-dms/stvd-dms-service-api"
	cmds := []string{"/bin/bash", "-c", *command}
	execCommand(*dir, cmds...)
	// execCommand("mvn clean")
	// execCommand("df -lh")
	// execCommand("whoami")
	// execCommand("ls")
}

func execCommand(dir string, commands ...string) bool {
	commandName := commands[0]
	args := commands[1:]
	cmd := exec.Command(commandName, args...)
	cmd.Dir = dir
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return false
	}

	//实时循环读取输出流中的一行内容
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && "EOF" == err.Error() {
			return true
		}
		fmt.Printf(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	// cmd.Wait()
	// return true
}
