#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <cstdio>

using namespace std;

int main(){
    int num = 0;
    vector<string>stringCon; 
    while(cin>>num){
    	cin.ignore();
        while(num--){
			string tmp;
            
            getline(cin, tmp);
            if(tmp == "stop"){
            	 
                sort(stringCon.begin(), stringCon.end(),[](string& a, string& b)->bool{return a.length()<b.length();});
                for(auto item:stringCon){
                    cout<<item<<endl;
                }
                break;
            }else{
                stringCon.push_back(tmp);
                if(num == 0){
                    sort(stringCon.begin(), stringCon.end(),[](string& a, string& b)->bool{return a.length()<b.length();});
                    for(auto item:stringCon){
                        cout<<item<<endl;
                    }
                    break;
                }
            }
        }
        stringCon.clear();
    }
    
    
    return 0;
}
