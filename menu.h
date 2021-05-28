#include <iostream>

#include "goto.h"
//sdl

#include <conio.h>
#include <windows.h>
#include <cstdlib>

using namespace std;
HANDLE console = GetStdHandle(STD_OUTPUT_HANDLE); // used for goto
COORD CursorPosition; // used for goto



void menu()
{
    
    system ("CLS");
        char t=-2;
        int run, x=17;
        bool running = true;
        int zmienna=1;
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
        gotoXY(console,CursorPosition,52,17); for(int i=0;i<13;i++){cout<<("%c",t);}
        SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
    while(running)
	{
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

        if(GetAsyncKeyState(VK_DOWN) && x != 32) //down button pressed
			{
				gotoXY(console,CursorPosition,52,x); cout << "             ";
				x=x+5;
				gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",t);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15); 
				zmienna++;
				continue;
				
			}
        if(GetAsyncKeyState(VK_UP) && x != 17) //up button pressed
			{
				gotoXY(console,CursorPosition,52,x); cout << "             ";
				x=x-5;
				gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",t);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15); 
				zmienna--;
				continue;
			}
        if(GetAsyncKeyState(VK_UP) && x == 17) //up button pressed
            {
                gotoXY(console,CursorPosition,52,x); cout << "             ";
                x=32;
                gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",t);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
                zmienna=4; 
                continue;
            }
        if(GetAsyncKeyState(VK_DOWN) && x == 32) //up button pressed
            {
                gotoXY(console,CursorPosition,52,x); cout << "             ";
                x=17;
                gotoXY(console,CursorPosition,52,x);
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),10);
                for(int i=0;i<13;i++){cout<<("%c",t);}
                SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);
                zmienna=1; 
                continue;
            }
        if(GetAsyncKeyState(VK_RETURN))
        { // Enter key pressed
            
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
                    
                    break;

                default:
                    system ("CLS");
                    exit(0);
                    break;
            }
        }
    }
}

