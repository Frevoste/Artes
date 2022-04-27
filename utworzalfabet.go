package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
	var litery = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	
	for i := range litery{
		
	f, err := os.Create(litery[i]+".txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString("old falcon\n")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("done")
	}
    
}