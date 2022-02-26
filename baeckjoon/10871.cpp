#include <iostream>

using namespace std;

int	main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int	n, x, num;
	cin >> n >> x;
	while (n--)
	{
		cin >> num;
		if (num < x)
			cout << num << " ";
	}
	return (0);
}
