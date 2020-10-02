package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var Conf Config

type Config struct {
	//RequestURL string `json:"request_url"`
	//Section    string
	Host       string
	Port       string
	Script     string
}

func LoadConfig() {
	flag.Usage = func() {
		fmt.Println("Usage: go-web-hook [-h] [-config path]")
		flag.PrintDefaults()
	}
	path := flag.String("config", "./config.json", "配置路径")
	flag.Parse()

	result, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := Config{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println("文件错误")
		fmt.Println(err)
		os.Exit(1)
	}

	Conf = data
}
