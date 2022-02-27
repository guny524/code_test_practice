#include <iostream>
#include <vector>

using namespace std;

int	main(void)
{
	vector<int> arr(26, 0);
	string str;
	ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cin >> str;
	for (int i=0; i<str.size(); i++)
		arr[str[i]-'a'] += 1;
	for (int i=0; i<26; i++)
		cout << arr[i] << " ";
	return (0);
}
