#include <stdio.h>
int maxProfit(int *prices, int pricesSize)
{
    if (pricesSize < 2)
        return 0;

    int sum = 0;

    int min = prices[0];
    int max = prices[1];
    if (max < min)
        min = max;

    for (int i = 2; i < pricesSize; i++)
    {
        if (prices[i] >= max)
        {
            max = prices[i];
        }
        else
        {
            sum = sum + (max - min);
            min = prices[i];
            max = min;
        }
    }

    if (max > min)
        sum = sum + (max - min);

    return sum;
}

int main(void)
{
    int array[] = {7, 6, 4, 3, 1};

    printf("%d", maxProfit(array, 5));

    return 0;
}