package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
	"math/rand"
)
const maxX = 21 //49
const maxY = 11 //18

func RandomString() string {
    var letters = []rune("aaaaaaaabbcccddddeeeeeeeeeeeeffgghhhhhhiiiiiiijkllllmmnnnnnnmmmmmmmooooooooppqrrrrrrssssssttttttttttuuuvwwwxyyz")
	//https://www3.nd.edu/~busiforc/handouts/cryptography/letterfrequencies.html
	//nie blokujacy zczyt z klawiatury informacja o nacisnieciu klawisza z sdl
	// curses biblioteka jak wczytywac hasla w terminalu
	// ncurses
	//https://stackoverflow.com/questions/49150316/how-to-detect-key-press-event
    s := make([]rune, 1)
	s[0] = letters[rand.Intn(len(letters))]
    
    return string(s)
}

func borders(maxX,maxY int) {
	
	
	fmt.Printf("\u001b[36m\x1B[%d;%df%s\u001b", 10, 75, "/\\")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 11, 74, "/  \\")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 12, 73, "|    |")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 13, 71, "--:'''':--")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 14, 73, ":'_' :")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 15, 73, "_:\"\":\\___")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 16, 58, "' '      ____.' :::     '._")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 17, 57, ". *=====<<=)           \\    :")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 18, 58, ".  '      '-'-'\\_      /'._.'")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 19, 75, "\\====|  \"\"")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 20, 74, ".'     \\")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 21, 73, ":       :")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 22, 72, "/   :    \\")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 23, 71, ":   .      '.")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 24, 71, ":  : :      :")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 25, 71, ":__:-:__.;--'")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 26, 71, "'-'   '-'")


	fmt.Printf("\u001b[33m\x1B[%d;%df%s\u001b", 24, 44, "/\\")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 25, 43, "/  \\")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 26, 42, "|    |")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 27, 40, "--:'''':--")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 28, 42, ":___.:")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 29, 39, "___/   \\_")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 30, 36, "_.'         '.____      ' '")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 31, 34, ":              _____|>=====* .")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 32, 34, "'._.'\\      _/'          '  .")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 33, 36, "\"\"  |====/")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 34, 39, "/     '.")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 35, 39, ":       :")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 36, 39, "/        \\")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 37, 37, ".'          :")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 38, 37, ":           :")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 39, 37, "'--:._______:")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b[0m", 40, 42, "'-'  '-'")
	
	
	for i := 1; i <= maxY; i++ {
			for j := 1; j <= maxX; j++{
			if(i==1||j==maxX && i!=maxY){
				fmt.Printf("\u001b\x1B[%d;%df%s\u001b", i+18, j+49, "\u001b[36m■\u001b[0m")
			}else if(j==1||i==maxY){
				fmt.Printf("\u001b\x1B[%d;%df%s\u001b", i+18, j+49, "\u001b[33m■\u001b[0m")
			}
		}
    } 
}

func czyszczenie() {
	cmd := exec.Command("cmd", "/c", "cls") 
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	
	rand.Seed(time.Now().Unix())
	
	var mapa [21][11]string
	
	
	borders(maxX,maxY)
	for z := 0; z < 100 ; z++ {
		
		for i :=  maxY-2; i > 0; i-- {
			for j :=  maxX-2; j > 0; j--{
				if(mapa[j][i]!=" "){
					if(i==maxY-2){
						mapa[j][i]=" "
					}else{
						mapa[j][i+1]=mapa[j][i]
						mapa[j][i]=" "
					}
				}
						
					
			}

			
		}
		for i:=0; i< rand.Intn(2)+1 ; i++{
			mapa[rand.Intn(maxX-2)+1][1]= RandomString()
		}
		
		for i := 0; i < maxY-1; i++ {
			for j := 0; j < maxX-1; j++{
				fmt.Printf("\u001b[32m\x1B[%d;%df%s\u001b[0m", i+19, j+50, mapa[j][i])
			}
		}
		
		
		time.Sleep(1000000000)
	}
	
	
	
	
	czyszczenie()
}