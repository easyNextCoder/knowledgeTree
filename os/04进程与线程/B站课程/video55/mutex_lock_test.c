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
pthread_mutex_t mutex;
void* thread_fun1(void* arg){
	
	while(1){
		//加锁，对整个机构体访问进行加锁，防止产生错乱
		pthread_mutex_lock(&mutex);
		stu.id = i;
		stu.age = i;
		stu.name = i;
		i++;
		if(stu.id != stu.age || stu.id != stu.name || stu.age != stu.name){
			printf("thread1 %d, %d, %d\n", stu.id, stu.age, stu.name);
			break;
		}
		//访问变量完成，需要进行解锁，只有这样其他线程才能访问
		pthread_mutex_unlock(&mutex);
	}
	return (void*)0;
}
void* thread_fun2(void* arg){
	
	while(1){
		//加锁，对整个结构体访问进行加锁，防止产生错乱
		pthread_mutex_lock(&mutex);
		stu.id = i;
		stu.age = i;
		stu.name = i;
		i++;
		if(stu.id != stu.age || stu.id != stu.name || stu.age != stu.name){
			printf("thread2 %d, %d, %d\n", stu.id, stu.age, stu.name);
			break;
		}
		pthread_mutex_unlock(&mutex);
	}
	return (void*)0;
}


int main(){
	
	pthread_t tid1, tid2;
	int err1, err2;

	//对互斥量进行初始化，只有初始化过的互斥量才能使用
	err1 = pthread_mutex_init(&mutex, NULL);

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
