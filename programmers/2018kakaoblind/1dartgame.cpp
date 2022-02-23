#include <string>

using namespace std;

enum	State
{
	STATE_NUMBER,
	STATE_SQUARE,
	STATE_BONUS
};

State	stateNumber(char str, int *i, int *preScore, int *score)
{
	static int	num = 0;

	if ('0' <= str && str <= '9')
	{
		num = num * 10 + str - '0';
		return (STATE_NUMBER);
	}
	else
	{
		*preScore = *score;
		*score = num;
		num = 0;
		--(*i);
	}
	return (STATE_SQUARE);
}

State	stateSquare(char str, int *score)
{
	if ('D' == str)
		*score = *score * *score;
	else if ('T' == str)
		*score = *score * *score * *score;
	return (STATE_BONUS);
}

State	stateBonus(char str, int *i, int *preScore, int *score, int *result)
{
	if ('*' == str)
	{
		*preScore *= 2;
		*score *= 2;
	}
	else if ('#' == str)
		*score *= -1;
	else
		--(*i);
	*result += *preScore;
	return (STATE_NUMBER);
}

int	solution(string dartResult)
{
	int		result = 0;
	int		score = 0;
	int		preScore = 0;
	State	state = STATE_NUMBER;
	int		i = 0;
	while (i < dartResult.size())
	{
		if (STATE_NUMBER == state)
			state = stateNumber(dartResult[i], &i, &preScore, &score);
		else if (STATE_SQUARE == state)
			state = stateSquare(dartResult[i], &score);
		else if (STATE_BONUS == state)
			state = stateBonus(dartResult[i], &i, &preScore, &score, &result);
		++i;
	}
	if (STATE_BONUS == state)
		result += preScore + score;
	else if(STATE_NUMBER == state)
		result += score;
	return result;
}
