#include <iostream>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int n;
	cin >> n;
	int arr[202] = {}, tmp;
	for (int i=0; i<n; i++)
	{
		cin >> tmp;
		arr[tmp+100]++;
	}
	cin >> tmp;
	cout << arr[tmp+100];
	return (0);
}
