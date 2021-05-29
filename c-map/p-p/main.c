#include<stdlib.h>
#include<stdio.h>
#include<string.h>

/* *
 * Having a very hard time understanding why anybody will need a pointer to a 
 * pointer in this life. Its a fucken daunting exercise writing c i'll tell you
 * that.
 */
 
int allocstr(int size, char **retptr) {
   char *p = malloc(size + 1);
  if (p == NULL)
     return 0;
  else
     *retptr = p;
     return 1; 
}

int main() {
    char *str = "hello every fucken body?";
    char *copyStr;
    if (allocstr(strlen(str), &copyStr))
        strcpy(copyStr, str);
    else
        fprintf(stderr, "out of fucken memory\n");
}

