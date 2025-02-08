package security

import (
	
	"strings"
)




// func Encrypt(text string)string{
// 	tex := strings.Split(text, "")
// 	//var t string = ""
// 	for i := 0; i < len(tex); i++ {
// 		if(tex[i] == "a"){
// 			tex[i] = "="
// 		}else if(tex[i] == "e"){
// 			tex[i] = "-"
// 		}else if(tex[i] == "i"){
// 			tex[i] = "%"
// 		}
// 	}
// 	return strings.Join(tex, "")

// }
func D(text string)string{

	txt := strings.Split(text, "")

	for i := 0; i < len(txt); i++ {
		if(txt[i] == "="){
			txt[i] = "a"
		}else if(txt[i] == "-"){
			txt[i] = "e"
		}else if(txt[i] == "%"){
			txt[i] = "i"
		}
	}
	return strings.Join(txt, "")

}



