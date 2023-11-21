#ifndef STACK_H_INCLUDED
#define STACK_H_INCLUDED

using namespace std;

typedef char infotype;

struct Stack{
    infotype info[15];
    int Top;
};

void createStack_1303220145(Stack &S);
bool isEmpty_1303220145(Stack S);
bool isFull_1303220145(Stack S);
void push_1303220145(Stack &S, infotype x);
infotype pop_1303220145(Stack &S);
void printInfo_1303220145(Stack S);




#endif // STACK_H_INCLUDED
