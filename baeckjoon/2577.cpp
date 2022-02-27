#include <iostream>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int a, b, c, num;
	int arr[10] = {};
	cin >> a >> b >> c;
	num = a * b * c;
	while (num)
	{
		arr[num%10] += 1;
		num /= 10;
	}
	for (int i=0; i<10; i++)
		cout << arr[i] << "\n";
	return (0);
}
