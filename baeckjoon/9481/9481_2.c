#include <stdio.h>
#include <stdlib.h>
#define BOOL

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

void cover(Node *col)
{

}

void dlx_search(Node *head, int k)
{
	if (head->right == head)
	{
		cnt++;
		return ;
	}
	int min = ;
	for ()
	{
		choose column
		cover(column)
	}
}

typedef struct	_Node
{
	struct _Node	*up;
	struct _Node	*down;
	struct _Node	*left;
	struct _Node	*right;
	struct _Node	*col;
	int				size;

	int				pos;
	int				num;
}				Node;

Node *new_node(Node *col, int pos, int num)
{
	Node *ret = malloc(sizeof(Node));
	ret->up = ret->down = ret->left = ret->right = NULL;
	ret->col = col;
	ret->size = 0;
	ret->pos = pos;
	ret->num = num;
	return (ret);
}

Node *init_head()
{
	Node *head, *ret;
	head = ret = new_node(NULL, 0, 0);
	// tmp arr for fast init rows
	Node **pos_col = malloc(sizeof(Node*));
	Node **num_col = malloc(sizeof(Node*));
	// set position column
	for (int pos=0; pos<n; pos++)
	{
		head->right = new_node(NULL, pos, 0);
		head->right->left = head;
		head = head->right;
		pos_col[pos] = head;
	}
	// set num column
	for (int num=0; num<n; num++)
	{
		head->right = new_node(NULL, 0, num+1);
		head->right->left = head;
		head = head->right;
		num_col[num] = head;
	}
	// for circular loop
	head->right = ret;
	ret->left = head;

	// set rows
	for (int pos=0; pos<n; pos++)
	{
		for (int num=0; num<n; num++)
		{
			Node *pos_node = new_node(pos_col[pos], pos, num+1);
			Node *num_node = new_node(num_col[num], pos, num+1);
			//#
		}
	}

	free(pos_col);
	free(num_col);
	return (ret);
}

int n, l, r, cnt;

int main(void)
{
	while (1)
	{
		scanf("%d %d %d", &n, &l, &r);
		if (n == l && l == r && r == 0)
			break;
		cnt = 0;
		Node *head = init_head();
		dlx_search(head, 0);
	}
}
