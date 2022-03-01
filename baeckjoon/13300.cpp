#include <iostream>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int n, k;
	cin >> n >> k;
	int f, g, arr[2][6] = {};
	for (int i=0; i<n; i++)
	{
		cin >> f >> g;
		arr[f][g-1]++;
	}
	int ret = 0;
	for (int i=0; i<2; i++)
		for (int j=0; j<6; j++)
			ret += !!(arr[i][j] % k) + arr[i][j] / k;
	cout << ret;
	return (0);
}
