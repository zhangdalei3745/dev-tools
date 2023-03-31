package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	userWorkspace := os.Getenv("USER_WORKSPACE")
	if userWorkspace == "" {
		fmt.Println("环境变量USER_WORKSPACE未设置")
		return
	}

	for {
		files, err := ioutil.ReadDir(userWorkspace)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 显示文件列表
		for i, file := range files {
			if file.IsDir() {
				fmt.Printf("%d. %s/\n", i+1, file.Name())
			} else {
				fmt.Printf("%d. %s/\n", i+1, file.Name())
			}
		}

		// 读取用户输入
		var input string
		fmt.Scanln(&input)

		// 处理用户输入
		index, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || index < 1 || index > len(files) {
			fmt.Println("无效的选择")
			continue
		}
		selectedPath := filepath.Join(userWorkspace, files[index-1].Name())
		if files[index-1].IsDir() {
			// 显示文件夹内容
			userWorkspace = selectedPath
		} else {
			// 执行命令
			arg := fmt.Sprintf("--kubeconfig=%s", selectedPath)
			cmd := exec.Command("k9s", arg)
			_ = cmd.Run()
			return
		}
	}
}
