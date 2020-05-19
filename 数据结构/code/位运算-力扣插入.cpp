#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;


class Solution {
public:
	void toBinary(int N)
	{
		vector<char> con;
		for(int i = 0; i<32; i++)
		{
			if( ( ((unsigned int)(1)<<i) & N ) ==  ((unsigned int)(1)<<i) )
				con.push_back('1');
			else
				con.push_back('0');
		}
		reverse(con.begin(), con.end());
		for(auto item:con)
			cout<<item;
		cout<<endl;
	}
    int insertBits(int N, int M, int i, int j) {
        int count = 0;
        for(int k = i; k<j; k++)
        {
            long long  int bit = (((long long int)1<<count)&M)>>count;
            cout<<bit<<endl;
            if(bit == 1)bit=0;
            else bit = 1;
            cout<<"before change:"<<endl;
            //toBinary(N);
            N = ((int)N|((long long  int)1<<k)) & (~(bit<<k));
            /*
            cout<<"after change:"<<endl;
			toBinary((int)N|((long long  int)1<<k));
			toBinary(~(bit<<k));
			int tmp =  (int)N|((long long  int)1<<k); 
			int tmp2 = (~(bit<<k)) ; 
			toBinary((  tmp & tmp2));
			N = tmp&tmp2;
            */
			toBinary(M);
			count++;
        }
        toBinary(N);
        toBinary(2082885133);
        return N;
    }
};

int main()
{
	Solution solution;
	solution.insertBits(1143207437,
						1017033,
						11,
						31);
						
	cout<<"test 2&1"<<(2&1)<<endl;
	
	return 0;
}

/*
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;


class Solution {
public:
	void toBinary(int N)
	{
		vector<char> con;
		for(int i = 0; i<32; i++)
		{
			if( ( ((unsigned int)(1)<<i) & N ) ==  ((unsigned int)(1)<<i) )
				con.push_back('1');
			else
				con.push_back('0');
		}
		reverse(con.begin(), con.end());
		for(auto item:con)
			cout<<item;
		cout<<endl;
	}
    int insertBits(int N, int M, int i, int j) {
        int count = 0;
        for(int k = i; k<j; k++)
        {
            long long  int bit = (((long long int)1<<count)&M)>>count;
            cout<<bit<<endl;
            if(bit == 1)bit=0;
            else bit = 1;
            cout<<"before change:"<<endl;
            toBinary(N);
            ( (long long int)N|((long long  int)1<<k) & (~(bit<<k)) );
            cout<<"after change:"<<endl;
			toBinary((int)N|((long long  int)1<<k));
			toBinary(~(bit<<k));
			int tmp =  (int)N|((long long  int)1<<k); 
			int tmp2 = (~(bit<<k)) ; 
			toBinary((  tmp & tmp2));
			N = tmp&tmp2;
            toBinary(M);
			count++;
        }
        toBinary(N);
        toBinary(2082885133);
        return N;
    }
};

int main()
{
	Solution solution;
	solution.insertBits(1143207437,
						1017033,
						11,
						31);
						
	cout<<"test 2&1"<<(2&1)<<endl;
	
	return 0;
}
*/
