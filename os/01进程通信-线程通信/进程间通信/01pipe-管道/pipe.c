#include<stdio.h>
#include<unistd.h>

int main()
{
    int fd[2] = {0};  // 两个文件描述符
    pid_t pid;
    char buff[20];

    printf("original fd0 is:%d, fd1 is:%d\n", fd[0], fd[1]);

    if(pipe(fd) < 0)  // 创建管道
        printf("Create Pipe Error!\n");

    
    printf("after pipe fd0 is:%d, fd1 is:%d\n", fd[0], fd[1]);

    if((pid = fork()) < 0)  // 创建子进程
        printf("Fork Error!\n");
    else if(pid > 0)  // 父进程
    {
        close(fd[0]); // 关闭读端
        write(fd[1], "hello world\0\n", 13);
        printf("father finished.\n");
    }
    else
    {
        close(fd[1]); // 关闭写端
        read(fd[0], buff, 20);
        printf("%s\n", buff);
        printf("son finished.\n");
    }

    return 0;
}