#include <iostream>
#include <vector>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	int n;
	cin >> n;

	int val[100000];
	for (int i=0; i<n; i++)
		cin >> val[i];

	int x;
	cin >> x;
	int ret = 0;

	bool arr[2000000];
	for (int i=0; i<n; i++)
	{
		if (x > val[i] && arr[x-val[i]])
			ret++;
		arr[val[i]] = true;
	}
	cout << ret;
	return (0);
}
