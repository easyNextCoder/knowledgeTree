#include <stdio.h> 
struct lseg_st{
	unsigned int a;
	unsigned int b;
};
int main()
{
	register  unsigned int regs;
    struct lseg_st les;
    
    int addr;
    les.a = 0xffffffff;
    les.b = 0x08;
	
     printf("addr 0x%X\n", regs);
    asm
    (
             "mov %1, %0\n"
            "int $80\n"
            :"=r"(regs):"r"(&les):"memory"
    );
	
    printf("addr 0x%X\n", les);
    printf("addr 0x%X\n", regs);
     
	return 0;
}
