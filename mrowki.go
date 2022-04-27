//IMIE: Mateusz
//Nazwisko: Redzimski
//Nr. Studenta: 266601
//Grupa: 3
//Tytuł: Artes


package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"os/exec"
	"os"
)

const maxX = 119
const maxY = 47

// pozycja jednego organizmu
type Position struct {
	x, y int
	back int
}

// tablica 1d współrzędnych wszystkich organizmów
type Coords []Position

type Setup struct {
	popul int // populacja
	iters int // liczba przerysowań ekranu
	dream int // opóźnienie = czas spania
	food int //liczba jedzonka
	hill int //wielkość mrowiska
}
//Liczenie wielkosci mrowiska przy uzyciu generatora funkcji
func genSequence (start,iteration int) func() int {
	start = start
	return func() int {
		start = start + iteration
		return start
	}
}
//Rysowanie obramówki terrarium
func borders(maxX,maxY int) {
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++{
			if(i==0||i==maxY){
				fmt.Printf("\u001b[36m■\u001b[0m")
			}else{
				if(j==0||j==maxX){
					fmt.Printf("\u001b[36m■\u001b[0m")
				}else{
					fmt.Printf(" ")
				}
			}
		}
        println()
    }
}
func setup(CFG *Setup) { //Ustawienie zasad symulacji
	flag.IntVar(&CFG.popul, "population", 100, "Population size or the ant")
	flag.IntVar(&CFG.iters, "iterations", 1000, "Number of generations to live")
	flag.IntVar(&CFG.dream, "slowness", 0, "More slowness - more lazy redraw")
	flag.IntVar(&CFG.food, "amount",300, "More food - less hungry")
	flag.IntVar(&CFG.hill, "size",10, "Bigger hill - less travel")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func makeant(mcount int) Coords {//Stworzenie mrówek
	col := make(Coords, mcount)
	for o := range col {
		col[o].x = rand.Intn(maxX-1) + 2
		col[o].y = rand.Intn(maxY-1) + 2
		col[o].back = 0

	}
		return col
}
func makeFood(fcount int) Coords {//Stworzenie jedzenia

	food := make(Coords,fcount)
	for o := range food {
		food[o].x = rand.Intn(maxX-1) + 2
		food[o].y = rand.Intn(maxY-1) + 2
		food[o].back = 1
	}
	return food
}
func makeAnthill(acount,maxacount int) Coords {//Stworzenie mrowiska
	ant := make(Coords, maxacount)
	//Podstawowe mrowisko
	for o := 0; o < acount; o++ {
		if(o==0){
			ant[o].x = rand.Intn(maxX-1) + 2
			ant[o].y = rand.Intn(maxY-1) + 2
			ant[o].back = 1
		}else{
			ant[o].x = ant[o-1].x + 1 - rand.Intn(3)
			ant[o].y = ant[o-1].y + 1 - rand.Intn(3)
			ant[o].back = 1

			for i:= 1 ; i<=o ; i++{
				if(ant[o].x==ant[o-i].x && ant[o].y==ant[o-i].y || ant[o].x <= 1 || ant[o].y <= 1 || ant[o].x > maxX-1 || ant[o].y > maxY -1 ){
					ant[o].back=0
					o=o-1
					i=o
				}
			}
		}
	}


	return ant
	}

//Czyszczenie ekranu dla windowsa
func czyszczenie() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {

	czyszczenie()
	borders(maxX,maxY)
	var config Setup
	setup(&config)
	gen := genSequence(config.hill,1)
	tmp := gen()
	ant := makeant(config.popul)
	food := makeFood(config.food)
	anth := makeAnthill(config.hill,config.hill+config.food)

	//  DODANIE JEDZENIA NA PLANSZE
	for o := range food {
		fmt.Printf("\u001b[32m\x1B[%d;%df%s\u001b[0m", food[o].y, food[o].x, "q")
	}

	// GŁÓWNA PĘTLA
	for iter := 0; iter < config.iters; iter++ {

		// RYSOWANIE
				for o := range ant {
			if ( ant[o].back == 0){
			fmt.Printf("\u001b[33m\x1B[%d;%df%s\u001b[0m", ant[o].y, ant[o].x, "o")
			}else{
				fmt.Printf("\u001b[31m\x1B[%d;%df%s\u001b[0m", ant[o].y, ant[o].x, "ó")
			}
		}
		for o := range anth {
			if(anth[o].back == 1 && anth[o].x != 0 && anth[o].y !=0){
				fmt.Printf("\u001b[35m\x1B[%d;%df%s\u001b[0m", anth[o].y, anth[o].x, "H")
			}
		}
		time.Sleep(time.Duration(config.dream))


		// CZYSZCZENIE
		for o := range ant {
				fmt.Printf("\x1B[%d;%df%s", ant[o].y, ant[o].x, " ")
			}
		// Sprawdzenie czy mrówka mogła zjesc jedzenie jeżeli nie to pojawia sie tam spowrotem
		for o := range ant {
			for i:= range food{
				if( ant[o].back == 1 && (ant[o].x == food[i].x && ant[o].y == food[i].y && food[i].back == 1)){
					fmt.Printf("\u001b[32m\x1B[%d;%df%s\u001b[0m", ant[o].y, ant[o].x, "q")
				}
			}
		}
		//Sprawdzenie czy mrowki cos podniosły
		for o := range ant {
			if(ant[o].back==1){
				o=o+1
			}else{
				for i := range food{

					if(ant[o].x == food[i].x && ant[o].y == food[i].y && food[i].back==1){
						ant[o].back = 1
						food[i].back = 0
					}
				}
			}

		}


		// PRZESUWANIE(GENEROWANIE nowych położeń)
		for o := range ant {
			ant[o].x += 1 - rand.Intn(3)
			if ant[o].x <= 1 {
				ant[o].x = 2
			}
			if ant[o].x > maxX-1 {
				ant[o].x = maxX - 1
			}

			ant[o].y += 1 - rand.Intn(3)
			if ant[o].y <= 1 {
				ant[o].y = 2
			}
			if ant[o].y > maxY-1 {
				ant[o].y = maxY - 1
			}
		}
		//Upuszczenie jedzenia przy mrowisku
		for o := range ant {
			//Sprawdzenie czy mrówka jest przy mrowisku
			if(ant[o].back == 1){
				if(anth[tmp].back == 1){
					tmp = gen()
				}
				for i := range anth {
					if(ant[o].x - 1 == anth[i].x && ant[o].y == anth[i].y){//czy mrowisko z lewej?

						anth[tmp].x = ant[o].x
						anth[tmp].y = ant[o].y
						anth[tmp].back = 1
						ant[o].back=0
						break
					}else if(ant[o].x  == anth[i].x && ant[o].y - 1 == anth[i].y){ //czy mrowisko z gory?

						anth[tmp].x = ant[o].x
						anth[tmp].y = ant[o].y
						anth[tmp].back = 1
						ant[o].back=0
						break
					}else if(ant[o].x + 1 == anth[i].x && ant[o].y  == anth[i].y){ //czy mrowisko z prawej?

						anth[tmp].x = ant[o].x
						anth[tmp].y = ant[o].y
						anth[tmp].back = 1
						ant[o].back=0
						break
					}else if(ant[o].x  == anth[i].x && ant[o].y + 1 == anth[i].y){ //czy mrowisko z dołu?

						anth[tmp].x = ant[o].x
						anth[tmp].y = ant[o].y
						anth[tmp].back = 1
						ant[o].back=0
						break
					}
				}


			}
		}
	}

	for o := range anth {
		fmt.Printf("\u001b[35m\x1B[%d;%df%s\u001b[0m", anth[o].y, anth[o].x, "H")
	}
	fmt.Printf("\x1B[%d;%df%s", 0, 0, "\u001b[36m■\u001b[0m")
	fmt.Printf("\x1B[%d;%df%s", maxY+2, 10, "Twoje piekne mrowisko :)     [KLIKNIJ ENTER ABY WROCIC DO MENU]")
	fmt.Scanf("%d")
}
