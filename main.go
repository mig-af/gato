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
	var axs bool = config.ConfigTrmx()
	if (axs){
		
		Init()
	}
	

}

var content embed.FS
func Init(){
	
	go utils.ServerInit("/")
	go utils.InitTunnel()
	
	time.Sleep(8*time.Second)
	utils.Send("As")
	//fmt.Println("siuuuuu")
	
	
	select {}
	
}



