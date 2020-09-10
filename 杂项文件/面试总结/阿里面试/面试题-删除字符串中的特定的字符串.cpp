#include <iostream>
#include <vector>
#include <map>
#include <algorithm>
#include <ctime>

using namespace std;

struct node {
	bool isWord;
	map<char, node*> nexts;
	node() :isWord(false) {};
};
node* root = new node();

void insert(node* root, string s)
{
	for (int i = 0; i < s.size(); ++i)
	{
		if (root->nexts.count(s[i]) == 0)
		{
			root->nexts[s[i]] = new node();
		}
		root = root->nexts[s[i]];
		if (i == s.size() - 1)root->isWord = true;
	}
}
int maxLen = 0;
int check(node* root, string& s, int left)
{
	int right = min(left + maxLen+1, (int)s.size());
	for (int i = left; i < right; ++i)
	{
		if (root->nexts.count(s[i]) == 0)
		{
			return -1;
		}
		else {
			root = root->nexts[s[i]];
		}
		if (root->isWord)return i;

	}
	return -1;
}
string news;
bool toDele(string s, node* root)
{
	int i = 0;
	for ( i = 0; i < s.size(); ++i)
	{
		int right = check(root, s, i);
		if (right == -1)continue;
		news.clear();
		for (int j = 0; j < i; ++j)
			news.push_back(s[j]);
		for (int j = right + 1; j < s.size(); ++j)
			news.push_back(s[j]);
		return true;
	}
	if (i == s.size())return false;
}

int main()
{
	/*
	string s, tmp;
	int n;
	cin >> s;
	cin >> n;
	for (int i = 0; i < n; ++i)
	{
		cin >> tmp;
		maxLen = max(maxLen, (int)tmp.size());
		insert(root, tmp);
	}
	news = s;
	while (toDele(news, root))
	{
		;
	}
	cout << news << endl;
	*/

	string old(100001, 's');
	string news = old;
	int start1 = clock();
	bool flag = true;;
	for (int i = 0; i < 10000; i++)
	{
		news.erase(10+flag, 1);
		if (flag)flag = !flag;
	}
	int end1 = clock();
	news = old;
	string news1;
	int start2 = clock();
	
	for (int i = 0; i < 10000; i++)
	{
		news1.clear();
		for(int j = i; j<10000; ++j)
			news1.push_back(news[j]);
	}
	int end2 = clock();
	cout << (end1 - start1) << ":" << (end2 - start2) << endl;
	return 0;
}