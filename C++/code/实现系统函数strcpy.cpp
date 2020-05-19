#include <stdio.h>
#include <assert.h>

char* strcpy(char* strDest, const char* strSrc)
{
	assert(strDest != NULL || strSrc != NULL);
	char* strDestc = strDest;
	while((*strDest++ = *strSrc++) != '\0');
	return strDestc;
}

int main()
{
	char d[] = "hel";
	char s[] = "qqqqq";
	strcpy(0,s);
	printf("%s\n", d);
	return 0;
}
/* 
1100
2082995725

1100
2082885133

1
0
0
1

0
0
1
1

0
0
1
0

0
0
0
1

1
1
1
1
0



1
0
0
1

0
0
1
1

0
0
1
0

0
0
0
1

1
1
1
1

1
1143207437
0
1143209485
0
1143209485
1
1143209485
0
1143209485
0
1143209485
1
1143209485
1
1143209485
0
1143471629
0
1143471629
1
1143471629
0
1143471629
0
1143471629
0
1143471629
0
1143471629
1
1143471629
1
1143471629
1
1277689357
1
1546124813
1
2082995725

*/



