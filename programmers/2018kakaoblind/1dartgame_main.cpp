#include <iostream>
#include <string>

using namespace std;

int	solution(string dartResult);

void	check(string input, int respect)
{
	if (solution(input) != respect)
	{
		cout << "input\t: " << input << endl;
		cout << "result\t: " << solution(input) << endl;
		cout << "respect\t: " << respect << endl << endl;
	}
}

int	main(void)
{
	check("1S2D*3T", 37);
	check("1D2S#10S", 9);
	check("1D2S0T", 3);
	check("1S*2T*3S", 23);
	check("1D#2S*3S", 5);
	check("1T2D3D#", -4);
	check("1D2S3T*", 59);
	return (0);
}
