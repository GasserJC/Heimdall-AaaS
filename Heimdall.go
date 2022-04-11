package main

import (
	"heimdall.com/app/FileIO"
	"heimdall.com/app/Serializing"
	"heimdall.com/app/User"
	"fmt"
)

func main() {
	usr := user.UC{ Key:"key", Username:"username", Password:"password"	}

	var headerString string = serial.Deserialize(usr.Slice(),"|")
	file.Write("data","output",headerString)
	var lines = file.Read("output")

	//for _,line := range lines {
		//fmt.Println(line)
	//}

	usrFromFile := serial.Serialize(lines[0])
	fmt.Println(usrFromFile.Key)
	fmt.Println(usrFromFile.Username)
	fmt.Println(usrFromFile.Password)
}