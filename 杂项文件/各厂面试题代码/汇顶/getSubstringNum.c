#include <stdio.h>

int del_sub(char* str, char* sub)
{
	if (str == 0 || sub == 0)
		return 0;
	int cnt = 0;
	for (int i = 0; str[i] != '\0'; ++i)
	{
		if (str[i] == sub[0])
		{
			int j = 0;
			for (j = 0; sub[j] != '\0'; ++j)
			{
				if (str[i + j] != sub[j])
					break;
			}
			if (sub[j] == '\0')cnt++;
		}
	}
	return cnt;
}

int main()
{
	char str[100];
	char sub[100];
	int num = 0;
	scanf("%s", str);
	scanf("%s", sub);

	num = del_sub(str, sub);
	printf("%d", num);

	return 0;
}