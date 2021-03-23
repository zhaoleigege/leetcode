#include <stdio.h>
#include <stdlib.h>

int *twoSum(int *nums, int numsSize, int target);

int main(void)
{
    int size = 4;
    int numArray[] = {3,
                      5,
                      7,
                      2};

    int *indexs = twoSum(numArray, size, 6);

    if (indexs != NULL)
    {
        printf("[%d,%d]\n", indexs[0], indexs[1]);
        free(indexs);
    }

    return 0;
}

int *twoSum(int *nums, int numsSize, int target)
{
    int *result = malloc(2 * sizeof(int));

    for (int i = 0; i < numsSize; i++)
    {
        for (int j = i + 1; j < numsSize; j++)
        {
            if ((nums[i] + nums[j]) == target)
            {
                result[0] = i;
                result[1] = j;

                return result;
            }
        }
    }

    free(result);
    return NULL;
}