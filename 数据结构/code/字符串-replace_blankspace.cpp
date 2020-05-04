#include <iostream>

using namespace std;


class Solution {
public:
	void replaceSpace(char *str,int length) {
        int blank = 0;
        for(int i = 0; i<length; i++){
            if(str[i] == ' '){
                blank++;
            }
        }
        
        char * str_c = new char[length+blank*2+1];
        /*
        swap<char*>(str_c, str);
        copy_if(str, str_c, [&](char item){
            if()
        })
        */
        int count = 0;
        int ori_count = 0;
        while(count < length+blank*2){
            if(str[ori_count] == ' '){
                str_c[count++] = '%';
                str_c[count++] = '2';
                str_c[count++] = '0';
                ori_count++;
            }else{
                str_c[count++] = str[ori_count++];
            }
        }
        str_c[count]='\0';
        str = str_c;
        std::cout<<str_c<<std::endl;
        return;
	}
};

int main(){
	Solution solution;
	char*s = "hello world";
	solution.replaceSpace(s, 11);
	cout<<s<<endl;
}
