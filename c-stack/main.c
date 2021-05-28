#include<stdio.h>
#include<stdlib.h>
#define STACK_MAX 100

// enum of error types that could occur
// given a stack
typedef enum {
    STACK_OK = 0,
    STACK_OVERFLOW,
    STACK_UNDERFLOW
} STACK_STATUS;

// typedef is the equivalent of type aliases in go.
// so basically STACK refers to a stack in memory.
typedef struct stack STACK;
struct stack {
    int data[STACK_MAX];
    int size;
};

STACK *NewStack() {
    STACK *s;
    s = malloc(sizeof(STACK)); /* malloc allocates memory on the heap to store var */
    s->size = 0;
    return s;
}

// pushes a new data item unto the stack
// if the push will result in the stack getting
// bigger than the stack max return an error
STACK_STATUS push(STACK *s, int data) {
    if (s->size < STACK_MAX) {
        // there is no len(array) function in c so
        // you have to keep len of array yourself
        // indexing is also not trivial in c
        s->data[s->size++] = data;
        return STACK_OK;
    }
    return STACK_OVERFLOW;
}

// return the last item that was pushed onto the stack.
// Its basically the item at the top of the stack
// it takes the stack and a pointer to an integer variable
// populates the int var with the poped item and reducess the size
// of the stack by 1
STACK_STATUS pop(STACK *s, int *d) {
    if (s->size > 0) {
        *d = s->data[s->size - 1]; /* put last item on stack in int var */
        s->size--; /* reduce s->size by one */
        return STACK_OK;
    }
    return STACK_UNDERFLOW;
}

int depth(STACK *s) {
    return s->size;
}

int main() {
    STACK *s = NewStack();
    int l, r;
    push(s, 3);
    push(s, 4);
    printf("depth = %d\n", depth(s));
    pop(s, &l);
    pop(s, &r);
    printf("l = %d, r = %d\n", l, r);
    printf("depth = %d\n", depth(s));   
}
