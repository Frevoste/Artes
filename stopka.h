
#include <iostream>
using namespace std;

void stopka()
{

    SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),14);
    cout << "                          " << "  \\---------------------------------------------------------/" << endl;
    SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),4);
    gotoXY(console,CursorPosition,33,38); cout <<   "  _|_|                _|                         "; 
    gotoXY(console,CursorPosition,33,39); cout << " _|    _|  _|  _|_|  _|_|_|_|    _|_|      _|_|_|  "; 
    gotoXY(console,CursorPosition,33,40); cout << " _|_|_|_|  _|_|        _|      _|_|_|_|  _|_|      ";
    gotoXY(console,CursorPosition,33,41); cout << " _|    _|  _|          _|      _|            _|_|" ;
    gotoXY(console,CursorPosition,33,42); cout << " _|    _|  _|            _|_|    _|_|_|  _|_|_|    " ;
}
