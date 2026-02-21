#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* hello(char* s) {
    char* result = malloc(strlen(s) + 12);
    sprintf(result, "Hello %s\n", s);
    return result;
}