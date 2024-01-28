#include <stdio.h>

void collatz(int num) {
  while (num != 1) {
    int eval = num;
    if (eval % 2 == 0) {
      num = eval / 2;
    } else {
      num = 3 * eval + 1;
    }
    printf("%d\n", num);
  }
}

int main() {
  int num;

  printf("Enter number: ");
  scanf_s("%d", &num);

  collatz(num);

  return 0;
}
