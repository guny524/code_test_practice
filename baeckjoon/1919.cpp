#include <iostream>

using namespace std;

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	string a, b;
	int arr[26] = {};

	cin >> a >> b;
	for (int i=0; i<a.size(); i++)
		arr[a[i]-'a']++;
	for (int i=0; i<b.size(); i++)
		arr[b[i]-'a']--;

	int ret = 0;
	for (int i=0; i<26; i++)
		ret += abs(arr[i]);
	cout << ret;
	return (0);
}
