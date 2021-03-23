#include <stdio.h>

int removeDuplicates(int *nums, int numsSize)
{
    if (numsSize < 1)
        return 0;

    int size = 1;
    for (int i = 1; i < numsSize; i++)
    {
        if (nums[i] != nums[size - 1])
        {
            nums[size] = nums[i];
            size = size + 1;
        }
    }

    return size;
}

int main(void)
{
    int array[] = {1, 1, 2};

    printf("%d", removeDuplicates(array, 3));

    return 0;
}