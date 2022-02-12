package util

import (
	"fmt"
	"os"
	"os/exec"
)

// Mkdir mkdir
func Mkdir(path string) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}

	return nil
}

// DelDir 删除文件夹
func DelDir(dir string) error {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return nil
	}

	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("rm -rf %s", dir))

	if err := cmd.Start(); nil != err {
		return err
	}

	if err := cmd.Wait(); nil != err {
		return err
	}

	return nil
}
