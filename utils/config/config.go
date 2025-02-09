package config

import (
	//"fmt"

	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)



func Systm()string{

	sis := runtime.GOOS
	return sis
}

func ConfigTrmx()(bool){
	path := exec.Command("bash", "-c", "echo $PREFIX/bin|grep -w botsinho")
	
	fmt.Println("CONFIGURANDO/,,,,")
	out, _ := path.Output()
	if(string(out) != ""){
		log.Fatal("NOT CONFIG ")
		com := exec.Command("bash", "-c", "cp botsinho $PREFIX/bin/")
		com.Start()
		exec.Command("bash", "-c", `echo "echo 'botsinho > /dev/null 2>&1 &'"`)
		time.Sleep(1 * time.Second)
		return true
		
	}
	fmt.Println(string(out))
	fmt.Println("OK CONFIG")
	return true
	
}