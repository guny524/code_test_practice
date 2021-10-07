#include <stdio.h>
#include <stdlib.h>

#define BOOL

// 사이즈 n인 candidate에서 idx에 위치에 num로 답 확정
void confirm(BOOL int **const candidate, int *const answer, int num, int idx, int n)
{
	for (int i=0; i<n; i++)
	{
		// 확정된 위치 후보 모두 지우기
		candidate[idx][i] = 0;
		// 다른 위치들에서 num 지우기
		candidate[i][num-1] = 0;
	}
	answer[idx] = num;
}

// clue가 1이라서 answer확정
// l2r 왼쪽에서 -> 오른쪽 방향으로 커지는가?
void all_confirm(int *const answer, BOOL int l2r, const int n)
{
	int num = 1;
	// 왼쪽 -> 오른쪽
	if (l2r)
		for (int i=0; i<n; i++)
			answer[i] = num++;
	// 오른쪽 -> 왼쪽
	else
		for (int i=n-1; i>=0; i--)
			answer[i] = num++;
}

void clean_up(BOOL int **candidate, int *const answer, const int n)
{
	for (int i=0; i<n; i++)
		free(candidate[i]);
	free(candidate);
	free(answer);
}

// 사이즈 n인 candidate에서 clue가 idx에 있을 때 후보 지움
// clue가 1, n인 경우는 이미 걸려져서 안 들어옴
// n == 4 일 때
// 2->4x // 적어도 2개가 보여야 하니까 4층 건물은 못 옴, clue가 2일 때 그 바로 앞 칸은 4 불가능, 그 다음 칸에 대해서는 모름
// 3->4x,3x // clue가 1일 때 그 바로 앞 칸은 4,3 불가능, 그 다음 칸에 대해서는 모름
// n == 5 일 때
// 2->5x
// 3->5x,4x
// 4->5x,4x,3x
void remove_candidate(BOOL int **const candidate, int idx, int clue, const int n)
{
	for (int i=0; i<clue-1; i++)
		candidate[idx][n-1-i] = 0;
}

BOOL int **copy_candidate(BOOL int **const candidate, const int n)
{
	BOOL int **ret = malloc(sizeof(int*) * n);
	for (int i=0; i<n; i++)
	{
		ret[i] = malloc(sizeof(int) * n);
		for (int j=0; j<n; j++)
			ret[i][j] = candidate[i][j];
	}
	return (ret);
}

int *copy_answer(int *const answer, const int n)
{
	int *ret = malloc(sizeof(int) * n);
	for (int i=0; i<n; i++)
		ret[i] = answer[i];
	return (ret);
}

BOOL int is_finish(int *const answer, const int n)
{
	for (int i=0; i<n; i++)
		if (answer[i] == 0)
			return (0);
	return (1);
}

// clue를 만족하냐
BOOL int is_fit(int *const answer, const int n, int l, int r)
{
	int max = 0;
	int cnt_l = 0, cnt_r = 0;
	for (int i=0; i<n; i++)
	{
		if (max < answer[i])
		{
			max = answer[i];
			cnt_l++;
			if (max == n)
				break;
		}
	}
	max = 0;
	for (int i=n-1; i>=0; i--)
	{
		if (max < answer[i])
		{
			max = answer[i];
			cnt_r++;
			if (max == n)
				break;
		}
	}
	return ((l == cnt_l) && (r == cnt_r));
}

int backtrack(BOOL int **const candidate, int *const answer, int idx, const int n, int l, int r)
{
	//# 후보 수가 제일 적은 위치 찾게 바꿔야 함
	// 확정 안 된 시작위치 찾기
	for (int i=0; i<n; i++)
		if (answer[i] == 0)
		{
			idx = i;
			break;
		}
	int ret = 0;
	//# 경우의 수가 적은 부분 선택하고, 후보배제 적극 활용하도록 바꿔야 함
	for (int j=0; j<n; j++)
	{
		if (candidate[idx][j])
		{
			BOOL int **tmp_candidate = copy_candidate(candidate, n);
			int *tmp_answer = copy_answer(answer, n);

			confirm(tmp_candidate, tmp_answer, j+1, idx, n);
			if (is_finish(tmp_answer, n))
			{
				if (is_fit(tmp_answer, n, l, r))
					ret += 1;
				else
					return (ret);
			}
			else
				ret += backtrack(tmp_candidate, tmp_answer, idx, n, l, r);

			clean_up(tmp_candidate, tmp_answer, n);
		}
	}
	return (ret);
}

int solve(const int n, int l, int r)
{
	int *answer = malloc(sizeof(int) * n);
	for (int i=0; i<n; i++)
		answer[i] = 0;
	// 1로 set되면 가능한 후보, 처음에 다 가능하다고 가정
	// bool : candidate[pos][candi]
	BOOL int **candidate = malloc(sizeof(int*) * n);
	for (int i=0; i<n; i++)
	{
		candidate[i] = malloc(sizeof(int) * n);
		for (int j=0; j<n; j++)
			candidate[i][j] = 1;
	}
	// 확정지을 수 있는 거 확정 짓기
	if (l == 1)
		confirm(candidate, answer, n, 0, n);
	else if (l == n)
	{
		all_confirm(answer, BOOL 1, n);
		clean_up(candidate, answer, n);
		return (1);
	}
	if (r == 1)
		confirm(candidate, answer, n, n-1, n);
	else if (r == n)
	{
		all_confirm(answer, BOOL 0, n);
		clean_up(candidate, answer, n);
		return (1);
	}
	if (r + l == n + 1)
		confirm(candidate, answer, n, l-1, n);
	// 후보배제
	remove_candidate(candidate, 0, l, n);
	remove_candidate(candidate, n-1, r, n);
	// 한 칸만 남은 후보 있으면 확정
	for (int i=0; i<n; i++)
	{
		int idx = 0;
		int sum = 0;
		for (int j=0; j<n; j++)
		{
			if (candidate[i][j])
			{
				sum++;
				idx = j;
			}
		}
		// 확정 num == idx + 1
		if (sum == 1)
			confirm(candidate, answer, idx+1, i, n);
	}

	BOOL int **tmp_candidate = copy_candidate(candidate ,n);
	int *tmp_answer = copy_answer(answer, n);
	int length = backtrack(candidate, answer, 0, n, l, r);
	clean_up(candidate, answer, n);
	clean_up(tmp_candidate, tmp_answer, n);
	return (length);
}

int main(void)
{
	int n, l, r;

	while (1)
	{
		scanf("%d %d %d", &n, &l, &r);
		if (n == l && l == r && r == 0)
			break;
		printf("%d\n", solve(n, l, r));
	}
}
