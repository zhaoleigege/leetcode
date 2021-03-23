#include <stdio.h>
void converse(int *array, int start, int end)
{
    while (start < end)
    {
        int temp = array[start];
        array[start] = array[end];
        array[end] = temp;

        start++;
        end--;
    }
}
/* 数组的转置ba=(a'b')' */
void rotate(int *nums, int numsSize, int k)
{
    if (k >= numsSize)
        k = k - numsSize;

    /*
    * 先旋转前numsSize-k位，再旋转后k位，最后整个数组再旋转一次
    */
    converse(nums, 0, numsSize - k - 1);
    converse(nums, numsSize - k, numsSize - 1);
    converse(nums, 0, numsSize - 1);
}

int main(void)
{
    int len = 1;
    int array[] = {-1};

    rotate(array, len, 2);

    for (int i = 0; i < len; i++)
        printf("%d ", array[i]);

    return 0;
}