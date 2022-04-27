#include <iostream>
#include <windows.h>
void intro()
{
    for(int i=0;i<15;i++)
    {
        cout << endl;
    }
    SetConsoleTextAttribute( //funckja ustawiajaca kolor tekstu
            GetStdHandle(STD_OUTPUT_HANDLE), // wyjście konsoli
                 4 // kod koloru
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
    SetConsoleTextAttribute( 
            GetStdHandle(STD_OUTPUT_HANDLE),
                 11 
                 );
    cout << endl;
    cout << endl;
    cout << "                                                                            @Mateusz_Redzimski" << endl;
    SetConsoleTextAttribute( 
            GetStdHandle(STD_OUTPUT_HANDLE), 
                 15 
                 );
    
}
