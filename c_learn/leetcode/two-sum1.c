#include <stdlib.h>
int *twoSum(int *nums, int numsSize, int target, int *returnSize) {
  int sum;
  int *ans = malloc(2 * sizeof(int));
  for (int i = 0; i < numsSize; i++) {
    for (int j = 0; j < numsSize; j++) {
      sum = nums[i] + nums[j];
      if (sum == target && i != j) {
        ans[0] = i;
        ans[1] = j;
        *returnSize = 2;
        printf("%d", *ans);
        return ans;
      } else {
        sum = 0;
      }
    }
  }
  *returnSize = 0;
  return NULL;
}

void main() {
  int nums[4] = {2, 7, 11, 15};
  int numSize = 4;
  int target = 9;
  int *returnSize;
  twoSum(nums, numSize, target, returnSize);
}
