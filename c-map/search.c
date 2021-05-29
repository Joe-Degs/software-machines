#include<stdlib.c>
#include "assoc_array.c"

/*
 * This hash map implementation was broken down into 3 parts for easy
 * understanding.
 * we have an associative array, a search type and the map itself.
 *  */


struct search {
    char *term, *value; /* term is the key we want the associated value for */
    assoc_array_t *cursor, *memo; /* cursor is where we are in the list */
    // memo is the previous cursor
};
typedef struct search search_t;

/*
 * on a new search, we have the key/term we are looking for
 * the current item in the assoc_array as the cursor
 * no memo because its a new search
 * no value yet because we havent started the search
 * */
search_t *search_new(assoc_array_t *a, char *k) {
    search_t *s = malloc(sizeof(search_t));
    s->term = k;
    s->value = NULL;
    s->cursor = a;
    s->memo = NULL;
    return s;
}

// check if the current cursor is the has the value we have the key for.
void search_step(search_t *s) {
    s->value = assoc_array_get_if(s->cursor, s->term);
}

// if we do have the value and cursor is not null
// we continue the search for value.
int searching(search_t *s) {
    return s->value == NULL && s->cursor != NULL;
}

// find loops through the items in the assoc_array, looking for the
// potential value. We stop if we get the value.
// on every run, if we have the value stop the loop
// if we don't have the value but the cursor is NULL stop the loop
// if we don't have the value and the cursor is not NULL move on to
// the next item in the list and keep the previous item as a memo.
search_t *search_find(assoc_array_t *a, char *k) {
    search_t *s = search_new(a, k);
    for (search_step(s); searching(s); search_step(s)) {
        s->memo = s->cursor;
        s->cursor = s->cursor->next;
    }
    return s;
}
