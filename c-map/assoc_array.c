#include<stdlib.h>
#include<string.h>

// assoc_array is a singly linked list with two items
// (key, value) in each of the sets.
struct assoc_array {
    char *key;
    void *value;
    struct assoc_array *next;
};
typedef struct assoc_array assoc_array_t

// hash map are just tables that have keys mapping to some value.
// how to maintain this table is the big problem
// how do you create a table of keys that map to some value than you
// can get anytime you want.

// returns a pointer to an assoc_array which is a singly linked list
// with two values.
assoc_array_t *assoc_array_new(char *k, char *v) {
    assoc_array_t *a = malloc(sizeof(assoc_array_t));
    a->key = strdup(k);
    a->value = strdup(v);
    a->next = NULL;
    return a;
}

// returns the value if the current entry in the list has the key.
char *assoc_array_get_if(assoc_array_t *a, char *k) {
    char *r = NULL;
    if (a != NULL && strcmp(a->key , k) == 0) {
        r = strdup(a->value);
    }
    return r;
}
