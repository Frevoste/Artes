#include<windows.h>
#include <iostream>
#include "okno.h"
#include "intro.h"
#include "menu.h"
using namespace std;

int main()
{
    system ("CLS");
    okno();
    intro();
    Sleep(3000);
    menu();
    return 0;
}
