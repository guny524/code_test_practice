#include <iostream>

using namespace std;

void blank(const int n, const int line)
{
	for (int blank=0; blank<n-line-1; blank++)
		cout << " ";
}

void draw(const int n, const int line)
{
	if (n == 3)
	{
		switch (line)
		{
			case 2:
				cout << "*****";
				break;
			case 1:
				cout << "* ";
			case 0:
				cout << "*";
		}
		return ;
	}
	draw(n/2, line%(n/2));
	if (line < n/2)
		return ;
	blank(n/2, line%(n/2));
	cout << " ";
	blank(n/2, line%(n/2));
	draw(n/2, line%(n/2));
}

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int n;
	cin >> n;
	for (int line=0; line<n; line++)
	{
		blank(n, line);
		draw(n, line);
		blank(n, line);
		cout << " \n";
	}
	return (0);
}
