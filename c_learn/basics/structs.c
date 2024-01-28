#include <stdio.h>

void main() {
  typedef struct {
    int age;
    char *name;
  } PERSON;

  PERSON flavio = {37, "Flavio"};
  printf("%d,%s", flavio.age, flavio.name);
}
