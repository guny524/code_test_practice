#include <iostream>
#include <list>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	string str;
	list<char> l;

	cin >> str;
	for (auto c : str)
		l.push_back(c);

	int	cnt;
	auto iter = l.end();

	cin >> cnt;
	while(cnt--)
	{
		cin >> str;
		switch (str[0])
		{
			case 'L':
				if (iter != l.begin())
					--iter;
				break;
			case 'D':
				if (iter != l.end())
					++iter;
				break;
			case 'B':
				if (iter != l.begin())
					iter = l.erase(--iter);
				break;
			case 'P':
				cin >> str;
				l.insert(iter, str[0]);
				break;
		}
	}
	for (auto c : l)
		cout << c;
	return (0);
}
