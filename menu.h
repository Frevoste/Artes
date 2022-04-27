#include <iostream>

#include "goto.h"
//sdl

#include <conio.h>
#include <windows.h>
#include <cstdlib>

using namespace std;

// zmienne używane przez funkcje goto znajdująca się w pliku goto.h 
HANDLE console = GetStdHandle(STD_OUTPUT_HANDLE); 
COORD CursorPosition; 



void menu()
{
    
    system ("CLS");
        char kwadrat=-2;
        int run, x=17;
        bool running = true;
        int zmienna=1;
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
        gotoXY(console,CursorPosition,52,17); for(int i=0;i<13;i++){cout<<("%c",kwadrat);} //printowanie zielonych kwadratów pod pierwszą opcją która jest wybrana odgórnie
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
    while(running)
	{

        //Przekazanie walorów estetycznych, ciekawostka nazwa Artes wywodzi się z połączenia słów "Arłukowicz" + "Test"
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),4);
        gotoXY(console,CursorPosition,33,5); cout <<   "  _|_|                _|                         "; 
        gotoXY(console,CursorPosition,33,6); cout << " _|    _|  _|  _|_|  _|_|_|_|    _|_|      _|_|_|  "; 
        gotoXY(console,CursorPosition,33,7); cout << " _|_|_|_|  _|_|        _|      _|_|_|_|  _|_|      ";
        gotoXY(console,CursorPosition,33,8); cout << " _|    _|  _|          _|      _|            _|_|" ;
        gotoXY(console,CursorPosition,33,9); cout << " _|    _|  _|            _|_|    _|_|_|  _|_|_|    " ;
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),14);
        gotoXY(console,CursorPosition,28,11); cout << "/------------------------[  MENU  ]-----------------------\\";
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
        gotoXY(console,CursorPosition,52,16); cout << "< [1]MROWKI >";
        gotoXY(console,CursorPosition,50,21); cout << "< [2]CZARODZIEJ >";
        gotoXY(console,CursorPosition,52,26); cout << "< [3]XXXXXX >  ";
        gotoXY(console,CursorPosition,53,31); cout << "< [4]EXIT >";
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),11);
        gotoXY(console,CursorPosition,67,36); cout <<"@Mateusz_Redzimski";
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),14);
        gotoXY(console,CursorPosition,28,37); cout << "\\---------------------------------------------------------/";
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),4);
        gotoXY(console,CursorPosition,33,39); cout <<  " _|    _|  _|            _|_|    _|_|_|  _|_|_|    "; 
        gotoXY(console,CursorPosition,33,40); cout << " _|    _|  _|          _|      _|            _|_|"; 
        gotoXY(console,CursorPosition,33,41); cout << " _|_|_|_|  _|_|        _|      _|_|_|_|  _|_|      ";
        gotoXY(console,CursorPosition,33,42); cout << " _|    _|  _|  _|_|  _|_|_|_|    _|_|      _|_|_|  " ;
        gotoXY(console,CursorPosition,33,43); cout << "  _|_|                _|                         " ;
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
        system("pause>nul");

        if(GetAsyncKeyState(VK_DOWN) && x != 32) //Gdy wciśniemy strzałkę w dół przeskakujemy o jedna pozycje w menu niżej
			{
				gotoXY(console,CursorPosition,52,x); cout << "             ";
				x=x+5;
				gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",kwadrat);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15); 
				zmienna++;
				continue;
				
			}
        if(GetAsyncKeyState(VK_UP) && x != 17) //Gdy wciśniemy strzałkę w górę przeskakujemy o jedną pozycję w menu wyżej
			{
				gotoXY(console,CursorPosition,52,x); cout << "             ";
				x=x-5;
				gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",kwadrat);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15); 
				zmienna--;
				continue;
			}
        if(GetAsyncKeyState(VK_UP) && x == 17) //Gdy wciśniemy strzałkę w górę ale jesteśmy na granicy to przeskoczymy do ostatniej opcji
            {
                gotoXY(console,CursorPosition,52,x); cout << "             ";
                x=32;
                gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",kwadrat);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
                zmienna=4; 
                continue;
            }
        if(GetAsyncKeyState(VK_DOWN) && x == 32) //Gdy wciśniemy strzałkę w dół ale jesteśmy na granicy to przeskoczymy do pierwszej opcji
            {
                gotoXY(console,CursorPosition,52,x); cout << "             ";
                x=17;
                gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",kwadrat);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
                zmienna=1; 
                continue;
            }
        if(GetAsyncKeyState(VK_RETURN))
        { // Po wciśnięciu enter następuje przekazanie zmiennej do Switch który wybiera opcję którą wskazaliśmy
            
            switch( zmienna )
            {
                case 1:
                    system("cls");
                    system("mrowki.exe");
                    menu();
                    break;

                case 2:
                    system("cls");
                    system("czarodziej.exe");
                    menu();
                    break;

                    //...
                case 3:
                    system("cls");
                    cout << "To tajne przez poufne nie możesz tutaj zaglądać";
                    menu();
                    break;

                default:
                    system ("CLS");
                    exit(0);
                    break;
            }
        }
    }
}

