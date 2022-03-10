#include <iostream>
#include <list>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	int n, k;
	cin >> n >> k;

	list<int> l;
	for (int i=1; i<=n; i++)
		l.push_back(i);

	cout << "<";
	auto iter = l.begin();
	while (!l.empty())
	{
		// for (auto c : l)
		// 	cout << c;
		// cout << "\n";
		for (int i=0; i<k-1; i++)
		{
			++iter;
			if (l.end() == iter)
				++iter;
		}
		cout << *iter;
		iter = l.erase(iter);
		if (l.end() == iter)
			++iter;
		if (!l.empty())
			cout << ", ";
	}
	cout << ">";
	return (0);
}
