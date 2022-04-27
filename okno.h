
#include<windows.h>


using namespace std;

void okno()
{
    HANDLE handle = GetStdHandle(STD_OUTPUT_HANDLE); // uchwyt standardowego wyjścia
    COORD c2; // struktura potrzebna do ustawienia rozmiarow okna
    c2.X = 120; 
    c2.Y = 50; 
    SetConsoleScreenBufferSize(handle, c2); 

    SMALL_RECT sr; // struktura wykorzystywana do ustawienia rozmiaru okna
    sr.Left = 0; 
    sr.Top = 0; 
    sr.Right = c2.X-1; 
    sr.Bottom = c2.Y-1; 
    SetConsoleWindowInfo(handle,true,&sr); // ustawia rozmiar okna jednostką miary są znaki
    ShowScrollBar(GetConsoleWindow(), SB_BOTH, FALSE);
    SetWindowLong(GetConsoleWindow(), GWL_STYLE, GetWindowLong(GetConsoleWindow(), GWL_STYLE)&~WS_SIZEBOX);

}

