#include <iostream>

using namespace std;

int	recursive(int s, const int& n, const int& k, int r, int c)
{
	int y = r % n;
	int x = c % n;
	if ((n-k)/2 <= x && x < (n-k)/2 + k
		&& (n-k)/2 <= y && y < (n-k)/2 + k)
		return (1);
	else if (!s)
		return (0);
	else
		return (recursive(s-1, n, k, r/n, c/n));
}

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int s, n, k, r1, r2, c1, c2;
	cin >> s >> n >> k >> r1 >> r2 >> c1 >> c2;
	for (int r=r1; r<=r2; r++)
	{
		for (int c=c1; c<=c2; c++)
			cout << recursive(s, n, k, r, c);
		cout << "\n";
	}
	return (0);
}
