#include <iostream>
#include "Stack.h"

using namespace std;

void createStack_1303220145(Stack &S){
    S.Top = 0;
}
bool isEmpty_1303220145(Stack S){
    if (S.Top == 0){
        return true;
    }else{
        return false;
    }
}
bool isFull_1303220145(Stack S){
    if (S.Top == 15){
        return true;
    }else{
        return false;
    }
}
void push_1303220145(Stack &S, infotype x){
    if (!isFull_1303220145(S)){
        S.Top += 1;
        S.info[S.Top] = x;
    }else{
        cout << "Maaf Stack Penuh"<< endl;
    }
}
infotype pop_1303220145(Stack &S){
    infotype x;
    if (isEmpty_1303220145(S)){
        cout << "Maaf stack Kosong" << endl;
    }else{
        x = S.info[S.Top];
        S.Top = S.Top - 1;
    }
    return x;
}
void printInfo_1303220145(Stack S){
    int i;
    for (i = S.Top; i >= 1 ;i--){
        cout << S.info[i];
    }
}
