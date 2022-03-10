#include <iostream>
#include <list>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	int tc;
	cin >> tc;

	for (int t=0; t<tc; t++)
	{
		string str;
		cin >> str;

		list<char> l;
		auto cursor = l.end();
		for (auto c : str)
		{
			switch (c)
			{
				case '<':
					if (l.begin() != cursor)
						--cursor;
					break;
				case '>':
					if (l.end() != cursor)
						++cursor;
					break;
				case '-':
					if (l.begin() != cursor)
						cursor = l.erase(--cursor);
					break;
				default:
					l.insert(cursor, c);
			}
		}

		for (auto c : l)
			cout << c;
		cout << "\n";
	}
	return (0);
}
