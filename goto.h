#include <iostream>
#include <conio.h>
#include <windows.h>
#include <cstdlib>

//Funkcja której używałem zanim poznałem magię escape codes, działa na zasadzie ustawienia odpowiednich koordynatów w terminalu windowsowym
void gotoXY(HANDLE console,COORD CursorPosition, int x, int y) 
{ 
	CursorPosition.X = x; 
	CursorPosition.Y = y; 
	SetConsoleCursorPosition(console,CursorPosition); 
}