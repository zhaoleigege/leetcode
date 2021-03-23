#include <stdio.h>
#include <stdlib.h>

int max(int num1, int num2, int value)
{
    int gap = num1 - num2;

    return gap > value ? gap : value;
}

int main(void)
{
    int n;
    scanf("%d", &n);

    int *array = malloc(n * sizeof(int));
    for (int i = 0; i < n; i++)
        scanf("%d", &array[i]);

    int gap = array[1] - array[0];
    int min = array[1] > array[0] ? array[0] : array[1];

    for (int i = 2; i < n; i++)
    {
        gap = max(array[i], min, gap);
        min = min > array[i] ? array[i] : min;
    }

    printf("最大差值%d\n", gap);

    return 0;
}