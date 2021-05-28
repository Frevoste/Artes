#include <iostream>
#include <windows.h>
void intro()
{
    for(int i=0;i<15;i++)
    {
        cout << endl;
    }
    SetConsoleTextAttribute( // ta funkcja ustawia kolor tekstu i t³a
            GetStdHandle(STD_OUTPUT_HANDLE), // pobieranie standardowego uchwytu wyjœcia konsoli
                 4 // kod koloru (dotyczy koloru t³a jak i tekstu
                 );
    cout << "                      " << "      ___           ___                         ___           ___     " << endl;
    cout << "                      " << "     /\\  \\         /\\  \\                       /\\__\\         /\\__\\    " << endl;
    cout << "                      " << "    /::\\  \\       /::\\  \\         ___         /:/ _/_       /:/ _/_   " << endl;
    cout << "                      " << "   /:/\\:\\  \\     /:/\\:\\__\\       /\\__\\       /:/ /\\__\\     /:/ /\\  \\  " << endl;
    cout << "                      " << "  /:/ /::\\  \\   /:/ /:/  /      /:/  /      /:/ /:/ _/_   /:/ /::\\  \\ " << endl;
    cout << "                      " << " /:/_/:/\\:\\__\\ /:/_/:/__/___   /:/__/      /:/_/:/ /\\__\\ /:/_/:/\\:\\__\\ " << endl;
    cout << "                      " << " \\:\\/:/  \\/__/ \\:\\/:::::/  /  /::\\  \\      \\:\\/:/ /:/  / \\:\\/:/ /:/  /" << endl;
    cout << "                      " << "  \\::/__/       \\::/~~/~~~~  /:/\\:\\  \\      \\::/_/:/  /   \\::/ /:/  / " << endl;
    cout << "                      " << "   \\:\\  \\        \\:\\~~\\      \\/__\\:\\  \\      \\:\\/:/  /     \\/_/:/  /  " << endl;
    cout << "                      " << "    \\:\\__\\        \\:\\__\\          \\:\\__\\      \\::/  /        /:/  /   " << endl;
    cout << "                      " << "     \\/__/         \\/__/           \\/__/       \\/__/         \\/__/    " << endl;
    SetConsoleTextAttribute( // ta funkcja ustawia kolor tekstu i t³a
            GetStdHandle(STD_OUTPUT_HANDLE), // pobieranie standardowego uchwytu wyjœcia konsoli
                 11 // kod koloru (dotyczy koloru t³a jak i tekstu
                 );
    cout << endl;
    cout << endl;
    cout << "                                                                            @Mateusz_Redzimski" << endl;
    SetConsoleTextAttribute( // ta funkcja ustawia kolor tekstu i t³a
            GetStdHandle(STD_OUTPUT_HANDLE), // pobieranie standardowego uchwytu wyjœcia konsoli
                 15 // kod koloru (dotyczy koloru t³a jak i tekstu
                 );
    for(int i=0;i<16;i++)
    {
        cout << endl;
    }
}
