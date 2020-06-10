package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	execCommand("whoami")
	execCommand("ls")
}

func execCommand(commandName string, params ...string) bool {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}

	//执行命令
	cmd.Start()

	//实时循环读取输出流中的一行内容
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Printf(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
}
