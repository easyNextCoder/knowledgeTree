#include <iostream>
#include <set>

using namespace std;

int main()
{
	set<int> test;
	test.insert(1);
	cout<<test.erase(1)<<endl;
	cout<<test.erase(1)<<endl;
	return 0;
}

/*
int main()
{
    std::set<int> c = {1, 2, 3, 4, 5, 6, 7, 8, 9};
 
    // erase all odd numbers from c
    for(auto it = c.begin(); it != c.end(); ) {
        if(*it % 2 == 1)
            it = c.erase(it);
        else
            ++it;
    }
 
    for(int n : c) {
        std::cout << n << ' ';
    }
}
*/
