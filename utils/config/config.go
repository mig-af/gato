package config

import (
	//"fmt"
	
	"os/exec"
	"runtime"
	"time"
)



func Systm()string{

	sis := runtime.GOOS
	return sis
}

func ConfigTrmx()bool{
	path := exec.Command("bash", "-c", "echo $PREFIX/bin|grep -w botsinho")
		
	out, _ := path.Output()
	if(string(out) == "botsinho"){
		return true
	}else{
		exec.Command("bash", "-c", "cp ./botsinho $PREFIX/bin/")
		exec.Command("bash", "-c", `echo "echo 'botsinho > /dev/null 2>&1 &'"`)
		time.Sleep(1 * time.Second)
		return ConfigTrmx()
	}
	
}