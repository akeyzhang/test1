package main

import "log"

func  init()  {
	log.SetFlags(log.Ldate|log.Lshortfile)
}
func main() {
    log.Println("aaaaaa")
}
