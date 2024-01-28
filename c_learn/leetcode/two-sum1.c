#include <stdio.h>

void main() {
  int nums[4] = {2, 7, 11, 15};
  int numSize;
  int target = 9;
  int returnSize;
}

int *twoSum(int *nums, int numsSize, int target, int *returnSize) {
  for (int i = 0; i < sizeof(nums); i++) {
    for (int j = 0; j < sizeof(nums); j++) {
      int found = nums[j] + nums[i];
      if (target == found) {
        int scrt[2] = {i, j};
        printf("%f", scrt);
        return scrt;
      } else {
        return 0;
      }
    }
  }
}
