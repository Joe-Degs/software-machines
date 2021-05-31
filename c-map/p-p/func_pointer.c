#include <stdio.h>

int Add(int a, int b) {
    return a + b;
}
// how to use this for callbacks.
int sqr(int a, int (*f)(int)) {
    return (*f)(a);
}

int callback(int a) {
    return a*a;
}

void main() {
    int c;
    int (*p)(int, int); // declaring a pointer that points to a func that takes ints.
    // this could be useful for callbacks or  something.
    p = &Add;
    c = (*p)(2, 3);
    printf("%d\n", c);

    // the callback thing works just fine.
    int sq = sqr(5, &callback); // passing the callback function without the taking
    // its reference is the same thing. 
    printf("%d\n", sq);
}
