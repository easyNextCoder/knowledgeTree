//heap sort
#include <iostream>
#include <random>
#include <algorithm>
using namespace std;

#define LEN 50

void swap(int* a, int* b) {
	int temp = *a;
	*a = *b;
	*b = temp;
}

void adjust_heap(int* heap, int start, int end) {

	int son = start * 2 + 1;
	while (son < end) {

	}
	if (son2 < end) {

	}
	else {

	}

	return;
}

int main() {
	int* heap = new int[LEN];
	default_random_engine e;
	heap[0] = INT_MAX;
	for (int i = 1; i < LEN; i++) {
		heap[i] = e() % 300;
	}

	for (int i = LEN/2; i>=1; i--)
		adjust_heap(heap, i, LEN);

	for (int i = LEN - 1; i >= 1; i--) {
		swap(heap+i, heap+1);
		adjust_heap(heap,1, LEN);
	}

	for (int i = 0; i < 50; i++) {
		cout << heap[i] << endl;
	}



	return 0;
}