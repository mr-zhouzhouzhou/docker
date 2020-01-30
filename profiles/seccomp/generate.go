// +build ignore

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/docker/docker/profiles/seccomp"
)

// saves the default seccomp profile as a json file so people can use it as a
// base for their own custom profiles
func main() {
	// 获取当前文件所在的路径
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f := filepath.Join(wd, "default.json")

	// write the default profile to the file
	b, err := json.MarshalIndent(seccomp.DefaultProfile(), "", "\t")
	if err != nil {
		panic(err)
	}
	// 这里是覆盖掉default.json 文件，可能是不同的版本，需要的参数不一致吧
	if err := ioutil.WriteFile(f, b, 0644); err != nil {
		panic(err)
	}
}
