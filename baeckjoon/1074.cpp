#include <iostream>

using namespace std;

int	recursion(int n, int r, int c, int div)
{
	if (n == 1)
		return (2*r + c);
	return (recursion(n-1, r%div, c%div, div/2) + div*div*(2*(r/div)+c/div));
}

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int n, r, c;
	cin >> n >> r >> c;
	cout << recursion(n, r, c, 1<<(n-1));
	return (0);
}
