#include <iostream>
#include <conio.h>
#include <windows.h>
#include <cstdlib>


void gotoXY(HANDLE console,COORD CursorPosition, int x, int y) 
{ 
	CursorPosition.X = x; 
	CursorPosition.Y = y; 
	SetConsoleCursorPosition(console,CursorPosition); 
}