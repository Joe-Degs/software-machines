#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* 
 * The goal is to find a substr in a string variable. in higher level languages,
 * you just call one function to get this done. But c is no other languages so
 * lets goooo.... */

typedef enum {
    ok = 0,
    not_ok
} status;

status findSubstr(char* str, char* substr, int* start) {
    int len = strlen(str);
    int subLen = strlen(substr);
    //printf("%d %d\n",len, subLen);

    for (int i = 0; i < len; i++) {
        if (subLen > (len - i)) {
            return not_ok;
        }
        if (str[i] == substr[0]) {
            char A[subLen+1];
            for (int j = 0; j < subLen+1; j++) {
                //printf("%d\n", j);
                if (j == subLen) {
                    A[j] = '\0';
                    break;
                }
                A[j] = str[i+j];
            }
            if (strcmp(A, substr) == 0) {
                *start = i;
                return ok;
            }
        }
    }

    return not_ok;
}

void main() {
    char* str = "joe is bae";
    char *substr = "hello";
    int start;
    status r = findSubstr(str, substr, &start);
    if (r == ok) {
        printf("starting index is %d\n", start);
    } else {
        printf("substr not in str\n");
    }
}
