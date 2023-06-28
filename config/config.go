package config

import (
	"fmt"
	"github.com/nacl2000/clould_storage/constant"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

const (
	M = "cloud_storage"
)

type Config struct {
}

var config = &Config{}

func Init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if !strings.Contains(dir, M) {
		panic(fmt.Sprintf("dir %s error", dir))
	}
	var confDir string
	for _, d := range strings.Split(dir, "/") {
		confDir += d + "/"
		if d == M {
			confDir += "conf/"
			break
		}
	}

	buf := readFile(constant.BaseFile)
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		panic(err)
	}
}

func readFile(fn string) []byte {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	return buf
}

func GetConfig() *Config {
	return config
}
