package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch()  {
	cmd1 := exec.Command("chmod", "a+x ../deploy.sh")
	cmd2 := exec.Command("sh", "../deploy.sh")
	err1 := cmd1.Start()
	err2 := cmd2.Start()
	if err1 != nil || err2 != nil{
		log.Fatal(err1)
	}
	err1 = cmd1.Wait()
	err2 = cmd2.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<h1>Hello, this is supercym deploy page</h1>")
	reLaunch()
}

func main()  {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8001", nil)
}
