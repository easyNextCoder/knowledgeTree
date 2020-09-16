//https://www.nowcoder.com/questionTerminal/0a92c75f5d6b4db28fcfa3e65e5c9b3f
#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include <map>
#include <algorithm>

using namespace std;

int main()
{
    string str;
    while(getline(cin,str))
    {
        stringstream sstr(str);
        string  token[2], tmp;
        int count = 0;
        while(getline(sstr, tmp, '-'))
        {
            token[count++] = tmp;
        }
        
        vector<string> hand[2];
        for(int i = 0; i<2; i++)
        {
            stringstream tmp(token[i]);
            string tstr;
            while(tmp>>tstr)
            {
                hand[i].push_back(tstr);
            }
        }
        map<string, int>order = {
            {"3",0},{"4",1},{"5",2},{"6",3},{"7",4},{"8",5},{"9",6},
            {"10",7},{"J",8},{"Q",9},{"K",10},{"A",11},{"2",12},
            {"joker",13},{"JOKER",14},
        };
        if(hand[0].size() == hand[1].size())
        {
            int hsize = hand[0].size();
 
            int outIndex = 0;
            if(order[hand[0][0]] < order[hand[1][0]])
            {
                outIndex = 1;
            }
            for(auto item:hand[outIndex])
            {
                cout<<item<<" ";
            }
            
        }else{
            vector<string> *con[2];
            int count = 0;
            con[0] = &hand[0];
            con[1] = &hand[1];
            for(int i = 0; i<2; i++)
            {
                if(hand[i].size() == 2 || hand[i].size() == 4)
                {
                    con[count++] = &hand[i];
                    con[count++] = &hand[1-i];
                    break;
                }
            }
            if(con[0]->size() == 2)
            {
                if((con[0]->at(0) == "joker" || con[0]->at(0) == "JOKER")&&
                    (con[0]->at(1) == "joker" || con[0]->at(1) == "JOKER"))
                {
                   cout<<"joker JOKER"<<endl; 
                }else{
                    if(con[1]->size() == 4)
                    {
                        for(auto item:*con[1])
                            cout<<item<<" ";
                    }else{
                        cout<<"ERROR"<<endl;
                    }
                }
            }else if(con[0]->size() == 4)
            {
                if(con[1]->size() == 4)
                {
                    int outIndex = 0;
                    if(con[0]->at(0)<con[1]->at(1))
                        outIndex = 1;
                    for(auto item:*con[outIndex])
                        cout<<item<<" ";
                }else{
                    for(auto item:*con[0])
                        cout<<item<<" ";
                }
            }else{
                cout<<"ERROR"<<endl;
            }
            
        }
        
        
    }
    
    return 0;
}