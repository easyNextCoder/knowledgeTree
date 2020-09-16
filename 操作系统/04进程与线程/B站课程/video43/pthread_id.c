#include <pthread.h>
#include <stdio.h>
int main(){
	pid_t pid;
	pthread_t tid;
	
	pid = getpid();
	tid = pthread_self();

	printf("pid is:%u, tid is:%u\n", pid, tid);
	return 0;
}
