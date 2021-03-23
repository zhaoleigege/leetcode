#include <stdio.h>

/**
 * 查找有序数组array前endIndex的值最后一个小于num的位置，并让数组整体往后移一位后插入num的值
 */
void move(int array[], int endIndex, int num)
{
    int index = endIndex - 1;
    while (index >= 0)
    {
        if (num > array[index])
            index--;
        else
            break;
    }

    for (int j = endIndex; j > index + 1; j--)
        array[j] = array[j - 1];

    array[index + 1] = num;
}

int main(void)
{
    int size = 10;
    int len = 3;

    int array[] = {90, 36, 4, 55, 71, 18, 0, 91, 89, 92};
    int result[len];

    int curIndex = 0;
    result[curIndex] = array[0];

    /* 先按顺序填充前len的数组 */
    for (int i = 1; i < len; i++)
    {
        if (array[i] > result[curIndex])
        {
            result[curIndex + 1] = result[curIndex];
            move(result, curIndex, array[i]);
        }
        else
        {
            result[curIndex + 1] = array[i];
        }

        curIndex = i;
    }

    /* 以后的数据和已填充的有序数组做对比 */
    for (int i = len; i < size; i++)
    {
        if (array[i] > result[len - 1])
        {
            move(result, len - 1, array[i]);
        }
    }

    for (int i = 0; i < len; i++)
        printf("%d ", result[i]);

    return 0;
}