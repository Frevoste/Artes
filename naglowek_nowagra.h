
#include <iostream>
using namespace std;

void naglowek_nowagra()
{
    for(int i=0; i<4;i++)
        cout << endl;
     SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),4);
    cout << "                                 "  << "   _|_|                _|                         " << endl;
    cout << "                                 "  << " _|    _|  _|  _|_|  _|_|_|_|    _|_|      _|_|_|  " << endl;
    cout << "                                 "  << " _|_|_|_|  _|_|        _|      _|_|_|_|  _|_|      " << endl;
    cout << "                                 "  << " _|    _|  _|          _|      _|            _|_|" << endl;
    cout << "                                 "  << " _|    _|  _|            _|_|    _|_|_|  _|_|_|    " << endl;
    for(int i=0; i<2;i++)
        cout << endl;
    SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),14);
    cout << "                          " << "  /----------------------[  Nowa Gra  ]---------------------\\" << endl;
    SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE),15);

}
