/*
 *作用：测试线程同步的重要性
 *作者: 徐永康
 *日期：2019.11.23
 */

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

struct Student{
	int id;
	int age;
	int name;
};

struct Student stu;
int i;
void* thread_fun1(void* arg){
	
	while(1){
		stu.id = i;
		stu.age = i;
		stu.name = i;
		i++;
		if(stu.id != stu.age || stu.id != stu.name || stu.age != stu.name){
			printf("thread1 %d, %d, %d\n", stu.id, stu.age, stu.name);
			break;
}
	}
	return (void*)0;
}
void* thread_fun2(void* arg){
	
	while(1){
		stu.id = i;
		stu.age = i;
		stu.name = i;
		i++;
		if(stu.id != stu.age || stu.id != stu.name || stu.age != stu.name){
			printf("thread2 %d, %d, %d\n", stu.id, stu.age, stu.name);
			break;
}
	}
	return (void*)0;
}


int main(){
	
	pthread_t tid1, tid2;
	int err1, err2;

	err1 = pthread_create(&tid1, NULL, thread_fun1, NULL);
	if(err1 != 0){
		printf("create thread1 failed.\n");
		return 0;
	}
	err1 = pthread_create(&tid2, NULL, thread_fun2, NULL);
	if(err1 != 0){
		printf("create thread2 failed.\n");
		return 0;
	}
	pthread_join(tid1, NULL);
	pthread_join(tid2, NULL);

	return 0;
}
