//IMIE: Mateusz
//Nazwisko: Redzimski
//Nr. Studenta: 266601
//Grupa: 3
//Tytuł: Artes

package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
	"bufio"
	"math/rand"
	"log"
	"github.com/go-vgo/robotgo" //Biblioteka która pozwala na Automatyczny Input bez oczekiwania na użytkownika
	"io/ioutil"
)
const maxX = 21 //49
const maxY = 11 //18
//Rysowanie wizerunku czarodzieja którym steruje gracz
func gracz(stan,hpgracz int){
	if(stan==0){ //Ten warunek służy określeniu czy rysujemy gracza normalnie czy też dostał on obrażenia w tym wypadku jest kolorowany na czerwono
		fmt.Printf("\u001b[33m\x1B[%d;%df%s\u001b", 24, 44, "/\\")
	}else if(stan==1){
		fmt.Printf("\u001b[31m\x1B[%d;%df%s\u001b", 24, 44, "/\\")
	}

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
	//Tutaj przy okazji rysoawnia wizerunku gracza rysujemy również jego stan życia
	fmt.Printf("\u001b[0m\x1B[%d;%df HP GRACZA:     \u001b", 41, 38)
	fmt.Printf("\u001b[0m\x1B[%d;%df HP GRACZA: \u001b", 41, 38)
	for i:= 0; i<=hpgracz;i++ {
		fmt.Printf("X")
	}
}
// Działa tak samo jak gracz z tą różnicą że znajduje się gdzie indziej
func przeciwnik(stan,hpprzeciwnik int){
	fmt.Printf("\u001b[0m\x1B[%d;%df HP PRZECIWNIKA:    \u001b", 9, 68)
	fmt.Printf("\u001b[0m\x1B[%d;%df HP PRZECIWNIKA: \u001b", 9, 68)
	for i:= 0; i<=hpprzeciwnik;i++ {
		fmt.Printf("X")
	}
	if(stan==0){
		fmt.Printf("\u001b[36m\x1B[%d;%df%s\u001b", 10, 75, "/\\")
	}else if(stan==1){
		fmt.Printf("\u001b[31m\x1B[%d;%df%s\u001b", 10, 75, "/\\")
	}
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
}
// funkcja służąca do rysowania pięknej kolorowej obramówki naszej planszy
func borders(maxX,maxY int) {

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
//Funkcja która pozwala na czyszczenie terminala przy użycia języka GOLANG
func czyszczenie() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()
}
//Funkcja która pozwala na użycie programu odczyt.pl który służy do sprawdzania czy słowo które podajemy niżej w programie znajduje się w słowniku anglojęzycznym
func perl() {
	perl := exec.Command("cmd", "/c", "odczyt.pl")
	if err := perl.Run(); err!=nil{
		fmt.Println("Error:",err)
	}

}
//Funkcja która tworzy nam []rune ze wszystkimi możliwymi literami w języku Angielskim z uwzględnieinem ich częstotliwości występowania i ztej właśnie runy wyciąga jedną losową literę którą wykorzystamy później przy wrzucaniu liter na planszę
func RandomString() string {
    var letters = []rune("aaaaaaaabbcccddddeeeeeeeeeeeeffgghhhhhhiiiiiiijkllllmmnnnnnnmmmmmmmooooooooppqrrrrrrssssssttttttttttuuuvwwwxyyz")
	//https://www3.nd.edu/~busiforc/handouts/cryptography/letterfrequencies.html
	s := make([]rune, 1)
	s[0] = letters[rand.Intn(len(letters))]

    return string(s)
}
func main() {

	hpgracz := 3
	hpprzeciwnik := 3
	var wynik string
	var sukces int
	var slowo string
	var tmps string
	var mapa [21][11]string
	//Go routine który pozwala na oczekiwanie na INPUT użytkownika bez zwieszania programu
	go func() {
		fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 40, 73, "Podaj swoje magiczne zaklęcie: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			slowo = scanner.Text()
		}
	}()

	rand.Seed(time.Now().Unix())
	borders(maxX,maxY)
	gracz(0,hpgracz)
	przeciwnik(0,hpprzeciwnik)
	//Główna pętla w której rozgrywa się gra
	for hpgracz>0 {
		slowo = ""
		//Tutaj właśnie wykorzystujemy wyżej napisaną goroutine
		<-time.After(time.Second * 2)
		if(len(slowo)>2){
			tmp := []byte(slowo) // podane wyzej slowo zmieniamy na zapisy bajtowy by uzyskac osobne literki

			for n :=0; n < len(slowo); n++{ // w tej petli sprawdzamy czy literki ze slowa ktore podalismy znajduja sie na planszy
				tmps = string(tmp[n])
				for i :=  maxY-2; i > 0; i-- {
					for j :=  maxX-2; j > 0; j--{
						if(tmps==mapa[j][i]){
							mapa[j][i]=" "
							j = 0
							i = 0
							if(n == len(slowo)-1){//jezeli wszystkie literki ze slowka znajduja sie na planszy uzyskujemy sukces i mozemy przejsc dalej
								sukces=1
							}
						}else if(j==1&&i==1&&tmps!=mapa[j][i]){ //jezeli doszlismy do konca tabeli i nie znalezlismy szukanej literki w tym momencie gracz traci zdrowie
							n=len(slowo)
							gracz(1,hpgracz)
							hpgracz--
							time.Sleep(1*time.Second)
							gracz(0,hpgracz)
						}
					}
				}
			}
		}
		if(sukces==1){ // jezeli udalo sie stworzyc slowo z literek znajdujacych sie na planszy wtedy przechodzimy do sprawdzenia czy znajduje sie ono w slowniku jezyka angielskiego
			file, err := os.Create("slowo.txt") // do wymiany danych miedzy programami zdecydowalem sie na pliki textowe poniewaz ulatwiaja mi one wyszukiwanie bledow
			if err != nil {                     // w tym momencie zapisujemy slowo w pliku zeby pozniej skrypt Perlowy mogl je odczytac i sprawdzic
				log.Fatal("Mamy tutaj blad",err)
			}
			defer file.Close()
			slowo = slowo + "\n"
			fmt.Fprintf(file,slowo)
			perl() // uruchamiamy zewnetrzny skrypt perla
			content, err := ioutil.ReadFile("wynik.txt") // tutaj pobieramy wynik ktory uzyskalismy dzieki uzyciu zewnetrznego skryptu
     		if err != nil {
         		 log.Fatal(err)
     		}
			wynik = string(content)
			if(wynik=="0"){ //jezeli skrypt Perlowy dal wynik negatywny to gracz traci zycie
				hpgracz--
				gracz(1,hpgracz)
				time.Sleep(1*time.Second)
				gracz(0,hpgracz)
			}
    		if(wynik=="1"){ //jezeli skrypt Perlowy dal wynik pozytywny to przeciwnik traci zycie
				przeciwnik(1,hpprzeciwnik)
				hpprzeciwnik--
				time.Sleep(1*time.Second)
				przeciwnik(0,hpprzeciwnik)
			}
			sukces = 0
		}
		//Pętla która przesuwa literki w dół ekranu zapisane są w tablicy dwuwymiarowej którą zadekrowaliśmy na początku programu
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

		// W tym miejscu korzystamy z funckji ktora pozwala na wylosowanie liter i umieszczamy te litery w pierwszym wierszu planszy
		for i:=0; i< rand.Intn(4)+1 ; i++{
			mapa[rand.Intn(maxX-2)+1][1]= RandomString()
		}

		// W tym miejscu rysujemy wszystkie literki w odpowiednich miejscach
		for i := 0; i < maxY-1; i++ {
			for j := 0; j < maxX-1; j++{
				fmt.Printf("\u001b[32m\x1B[%d;%df%s\u001b[0m", i+19, j+50, mapa[j][i])
			}
		}
		//W przypadku gdyby uzytkownik cos napisal a nie kliknal enter nie wiedzialby z jakich liter juz korzystal a dodaly by sie do slowa dlatego wykorzystuje zewnetrzna biblioteke robotgo ktora w wyniku braku aktywnosci gracza klika enter
		robotgo.KeyTap("enter")
		//tutaj przygotowanie miejsca pod wpisywanie naszych magicznych zaklęć
		fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 41, 73, "                                        ")
		fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 0, 74, "                                        ")
		fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 41, 73, "")
		//Wyzerowanie hp gracza w celu zakonczenia petli jezeli najpierw pokanamy przeciwnika
		if(hpprzeciwnik==0){
			hpgracz=0
		}
	}

	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 1, 55, "KONIEC GRY")
	fmt.Printf("\u001b\x1B[%d;%df%s\u001b", 2, 55, "")
	fmt.Scanf("%d")
	time.Sleep(5*time.Second) // czas na nacieszenie sie wynikiem gry :)


}
