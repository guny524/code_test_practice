#include<stdio.h>

void draw(int size, int col, int row)
{
	int div = size/3;
	if ( (div <= col && col < div*2) && (div <= row && row < div*2) )
	{
		printf(" ");
		return;
	}
	else if (size == 1)
	{
		printf("*");
		return;
	}
	draw(div, col % div, row % div);
}

int main()
{
	int n;
	scanf("%d", &n);
	for (int i = 0; i < n; i++)
	{
		for (int j = 0; j < n; j++)
		{
			draw(n, i, j);
		}
		printf("\n");
	}
	return 0;
}