#include <stdio.h>
#include <string.h>

void main() {
  char spam[4][7];

  strcpy_s(spam[0], 7, "banana");
  strcpy_s(spam[1], 7, "apples");
  strcpy_s(spam[2], 5, "tofu");
  strcpy_s(spam[3], 5, "cats");

  // Array size/ string size = number of strings
  int string_num = sizeof(spam) / sizeof(spam[0]);

  for (int i = 0; i < string_num; i++) {
    printf("%s", spam[i]);
    if (i < string_num - 2) {
      printf(", ");
    } else if (i == string_num - 2) {
      if (string_num > 2) {
        printf(", ");
      }
      printf("and ");
    }
  }
}
