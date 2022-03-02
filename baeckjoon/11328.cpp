#include <iostream>

using namespace std;

bool check(string a, string b)
{
	int arr1[26]={}, arr2[26]={};

	for (size_t j=0; j<a.size(); j++)
		arr1[a[j]-'a']++;
	for (size_t j=0; j<b.size(); j++)
		arr2[b[j]-'a']++;
	for (size_t j=0; j<26; j++)
		if (arr1[j] != arr2[j])
			return (false);
	return (true);
}

int main(void)
{
	int n;
	string a, b;

	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	cin >> n;
	for (int i=0; i<n; i++)
	{
		cin >> a >> b;
		if (check(a, b))
			cout << "Possible\n";
		else
			cout << "Impossible\n";
	}
	return (0);
}
