#include<stdio.h>
#include<stdlib.h>

/* Cactus/spaghetti stacks with cells all over the place.
 * Instead of one contigous chunk ofmemory as stack, this is 
 * a list spread all over the place with the each cell pointing
 * to the next one after it.
 * This is a singly linked list of stack cells in the heap.
 * */

typedef struct stack STACK;
struct stack {
    int data;
    STACK *prev;
};

// push add a new cell to the top of the list, and returns it.
STACK *push(STACK *s, int data) {
    STACK *r = malloc(sizeof(STACK));
    r->data = data;
    r->prev = s;
    return r;
}

// return data at the top of the stack.
STACK *pop(STACK *s, int *r) {
    if (s == NULL)
        exit(1);
    *r = s->data;
    return s->prev;
}

// iterate through the list and return
// number of cells in the list.
int depth(STACK *s) {
    int r = 0;
    for (STACK *t = s; t != NULL; t = t->prev) {
        r++;
    }
    return r;
}

// free the number of items from the list
void gc(STACK **old, int items) {
    STACK *t;
    for (; items > 0 && *old != NULL; items--) {
        t = *old;
        *old = (*old)->prev;
        free(t);
    }
}

int main() {
    STACK *s = push(NULL, 1);
    s = push(s, 2);
    int l;
    printf("depth pre-gc = %d\n", depth(s));
    s = pop(s, &l);
    printf("pop pre-gc = %d\n", l);
    printf("depth after pop = %d\n", depth(s));
    
    gc(&s, 2);
    printf("depth post-gc = %d\n", depth(s));
}
