#include <stdio.h>
//https://blog.csdn.net/dengtong7258/article/details/101888527
//��������Ա���Ա���95ҳ 
//extern const buffer_size;
int main()
{
	int buffer_size = 10;
	int arr[buffer_size] = {0};
	int i = 0;
	for(i = 0; i<buffer_size; i++)
		printf("%d\n", arr[i]);
	
	return 0;
}
