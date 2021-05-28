
#include<windows.h>


using namespace std;

void okno()
{
    HANDLE handle = GetStdHandle(STD_OUTPUT_HANDLE); // uchwyt standardowego wyj�cia
    COORD c2; // struktura potrzebna do ustawienia rozmiar�w bufora pami�ci
    c2.X = 120; // szeroko�� na 120 szeroko�ci znak�w
    c2.Y = 50; // wysoko�� na 50 wysoko�ci znak�w
    SetConsoleScreenBufferSize(handle, c2); // ustawia rozmiar bufora (wy�wietlanego tekstu)

    SMALL_RECT sr; // struktura wykorzystywana do ustawienia rozmiaru okna
    sr.Left = 0; // na zero
    sr.Top = 0; // na zero
    sr.Right = c2.X-1; // szeroko�� o 1 mniejsza od bufora
    sr.Bottom = c2.Y-1; // wysoko�� o 1 mniejsza od bufora
    SetConsoleWindowInfo(handle,true,&sr); // ustawia rozmiar okna (jednostka to szeroko�� i wysoko�� pojedynczego znaku)
    ShowScrollBar(GetConsoleWindow(), SB_BOTH, FALSE);
    SetWindowLong(GetConsoleWindow(), GWL_STYLE, GetWindowLong(GetConsoleWindow(), GWL_STYLE)&~WS_SIZEBOX);

}

