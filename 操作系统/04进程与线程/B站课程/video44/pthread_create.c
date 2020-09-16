/*
 *AUTHOR: XUYONGKANG
 *DATE:   2019-11-19
 *
 *
 * getpid() 
 * pthread_self()
 * int pthread_create(pthread_t * thread,
 *					  const pthread_attr_t *attr,
 *					  void *(* start_routine)(void*),
 *					  void *arg);
 * 第一个参数：新线程id,创建成功系统回填
 * 第二个参数：新线程到属性，NULL为默认属性
 * 第三个参数：新线程到启动函数
 * 第四个参数：传递给新线程
*/
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>
struct ThreadInfo{
	char * name;
	int length;
};
typedef struct  ThreadInfo thread_info;
void print_id(thread_info *input_thread){
	pid_t pid;
	pthread_t tid;

	pid = getpid();
	tid = pthread_self();

	printf("%s pid %u, tid %u\n", input_thread->name, pid, tid);

}
void * thread_fun(thread_info *arg){
	print_id(arg);
	return (void*)0;
}

int main(){
	pthread_t ntid;	
	int err;
	thread_info init_thread_info;
	init_thread_info.name = "new pthread";
	err = pthread_create(&ntid, NULL, thread_fun,&init_thread_info);
	if(err != 0){
		printf("err create thread fialed.\n");
		return 0;
	}
	thread_info source_thread_info;
	source_thread_info.name = "main thread";
	print_id(&source_thread_info);
	sleep(2);
}
