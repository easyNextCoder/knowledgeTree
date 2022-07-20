#include "stdio.h"
#include "unistd.h"
int main(){
	
	char buf[] = "hello linux\n";
	printf("%s",buf);
	while(1);
	return 0;
}
