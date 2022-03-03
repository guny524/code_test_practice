#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void	indexing(vector<int> &arr, vector<int> &arrh)
{
	for (int h=0; h<arrh.size(); h++)
		arrh[h] = upper_bound(arr.begin(), arr.end(), h+1) - lower_bound(arr.begin(), arr.end(), h+1);
	for (int h=arrh.size()-2; h>=0; h--)
		arrh[h] += arrh[h+1];
}

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	int n, h;
	cin >> n >> h;
	vector<int> even((n+1)/2), odd(n/2);
	for (int i=0; i<n; i++)
	{
		if (i%2)
			cin >> odd[i/2];
		else
			cin >> even[i/2];
	}
	sort(even.begin(), even.end());
	sort(odd.begin(), odd.end());

	vector<int> evenh(h), oddh(h);
	indexing(even, evenh);
	indexing(odd, oddh);

	vector<int> arr(h);
	for (int i=0; i<h; i++)
		arr[i] = evenh[i] + oddh[h-i-1];

	sort(arr.begin(), arr.end());

	cout << arr[0] << " " << upper_bound(arr.begin(), arr.end(), arr[0]) - arr.begin();
}
