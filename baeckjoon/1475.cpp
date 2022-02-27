#include <iostream>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int arr[10] = {};
	string str;
	cin >> str;
	for (int i=0; i<str.size(); i++)
		arr[str[i]-'0']++;
	int tmp = arr[6] + arr[9];
	if (tmp % 2)
		arr[6] = tmp/2 + 1;
	else
		arr[6] = tmp/2;
	int max = 0;
	for (int i=0; i<9; i++)
		if (max < arr[i])
			max = arr[i];
	cout << max;
	return (0);
}
