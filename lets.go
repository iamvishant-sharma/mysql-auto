package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute(){
	out, err := exec.Command("/bin/sh", "-c", "ansible-playbook site.yaml -e 'login_user=root login_password=7fpJ3ffm9AHYPOe563U9 db_name=diglet_ka_db username=diglet password=diglet priv=select,insert'").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command executed successfully")
	output := string(out[:])
	fmt.Printf(output)
}

func main(){
	if runtime.GOOS == "windows" {
		fmt.Println("Cannot execute command on windows")
	} else {
		execute()
	}
}

