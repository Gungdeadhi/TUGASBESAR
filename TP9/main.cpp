#include <iostream>
#include "Stack.h"

using namespace std;

int main()
{
    int i = 1;
    infotype a;
    Stack S;
    createStack_1303220145(S);
    while (i <= 11){
        cout << "Masukkan Character" << endl;
        cin >> a;
        push_1303220145(S,a);
        i++;
    }
    printInfo_1303220145(S);
    cout << endl;
    for (i = 1; i <= 8;i++){
         pop_1303220145(S);
    }
    printInfo_1303220145(S);

}
