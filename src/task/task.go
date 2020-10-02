package task

import (
	"go-web-hook/src/config"
	"os/exec"
)

var isRun = false
var haveMore = false

func Add() {
	println("start")
	start()
}

func check()  {
	if haveMore {
		println("have more")
		haveMore = false
		start()
	}
}

func start()  {
	if isRun {
		haveMore = true
		println("isRun")
		return
	}

	isRun = true
	out, _ := exec.Command("sh", config.Conf.Script).Output()
	println(string(out))
	isRun = false
	check()
}