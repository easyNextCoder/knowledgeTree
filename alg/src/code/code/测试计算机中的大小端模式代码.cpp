#include <iostream>

using namespace std;
union Student{
	int i;
	unsigned char ch[2];
};

#define BIG_ENDIAN 0
#define LITTLE_ENDIAN 1
int TestByteOrder()
{
	short int word = 0x0001;
	char *byte = (char*)&word;
	return (byte[0]?LITTLE_ENDIAN:BIG_ENDIAN);
}
  
int main()
{
	/*
	Student student;
	student.i = 0x1420;
	printf("%d %d", student.ch[0], student.ch[1]);
	*/
	
	cout<<TestByteOrder()<<endl;
	long long a = 1, b = 2, c = 3;
	//入栈顺序：先c 后b 再a
	//入栈后的分布如：0x01000000  0x00000000 0x02000000 0x00000000 0x03000000
	printf("%d %d %d\n", a, b, c);
	
	printf("%d, %d, %d", 'a', 'b','c');
	
	return 0;
	
}
