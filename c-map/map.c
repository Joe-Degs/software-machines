#include<stdlib.h>
#include<stdio.h>
#include<limits.h>
#include<string.h>
#include "search.c"

struct map {
    int size;
    assoc_array_t **chains; /* this, i do not understand */
};
typedef struct map map_t;

map_t *map_new(int size) <{
    map_t *m = malloc(sizeof(map_t));
    m->chains = malloc(sizeof(assoc_array_t) * size);
    for (int i = 0; i < size; i++) {
        m->chains[i] = NULL
    }
    m->size = size;
    return m;
}

int map_chain(map_t *m, char *k) {
    unsigned long int b;
    for (int i = strlen(k) - 1; b < ULONG_MAX && i > 0; i--) {
        b = b << 8;
        b += k[i];
    }
    return b % m->size
}

char *map_get(map_t *m, char *k) {
    search_t *s = search_find(m->chains[map_chain(m, k)], k);
    if (s != NULL) {
        return s->value;
    }
    return NULL;
}

void map_set(map_t *m, char *k, char *v) {
    int b = map_chain(m, k);
    assoc_array_t *a = m->chains[b];
    search_t *s = search_find(a, k);
    if (s->value != NULL) {
        s->cursor->value = strdup(v);
    }
}
