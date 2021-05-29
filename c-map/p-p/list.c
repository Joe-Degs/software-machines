#include <stdio.h>
#include <stdlib.h>

typedef struct linked_list list;
struct linked_list {
    int data;
    list *next;
};

list *push(list *l, int data) {
    list *nl = malloc(sizeof(list));
    nl->data = data;
    nl->next = l;
    return nl;
}

list *pop(list *l, int *r) {
    if (l == NULL)
        return NULL;
    else
        *r = l->data;
        return l->next;
}

// god! why is understanding this piece of shit so fucken hard!
// i've been trying to understand this pointer to a pointer bullshit but fuck it
// i'm fucken done with c. Oh jesus! i thought c was cute, no more!
void reMove(list **l, int data) {
    list **t;
    for (t = l; *t != NULL; t = &(*t)->next) {
        // advance through the list
        // make the old list keep track of the new one's position
       //*old = t;
       if ((*t)->data == data) {
           // we have to remove t
           (*t) = (*t)->next;
           //free(*t);
           break;
       }
       //*old = t;
    }
}

void printData(list *t) {
    for (; t != NULL; t = t->next) {
        printf("%d -> ", t->data);
    }
    printf("NULL\n");
}

int main() {
    list *t = push(NULL, 1);
    t = push(t, 2);
    t = push(t, 3);
    t = push(t, 4);
    t = push(t, 5);
    t = push(t, 6);

    // print the list
    printData(t);

    // remove one item from list
    reMove(&t, 3);
    printData(t);
}
