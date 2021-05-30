#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void printCharsInString() {
    char* hello = "hello the motherfucken world!";
    for (int i = 0; i < strlen(hello); i++) {
        printf("%d: '%c'\n", i, hello[i]);
    }
}

void main() {
    printCharsInString();
}
