/*
#include <iostream>
#include <string>
#include <vector>
#include <sstream>

using namespace std;

int main()
{
    string sl;
    
    while(getline(cin, sl))
    {
        vector<bool> con;
        vector<string> scon, scon1;
        stringstream ss(sl);
        string tmp;
        bool output = false;
        while(ss>>tmp)
        {
            
            if(tmp == "and"){
                if(ss>>tmp)
                {
                    if(scon.size()>0)
                    {
                        string last = scon.back();
                        scon.pop_back();
                        string llast = tmp;
                        
                        if((last == "true" || last == "false") && (llast == "true" || llast == "false"))
                        {
                            if(last == "true" && llast == "true")
                                scon.push_back("true");
                            else
                                scon.push_back("false");
                        }else{
                            if(!output)
                            {
                                cout<<"error"<<endl;
                                output = true;
                            }
                            break;
                        }
                    }else{
                        if(!output)
                        {
                            cout<<"error"<<endl;
                            output = true;
                        }
                        break;
                    }
                    
                }else{
                    if(!output)
                    {
                        cout<<"error"<<endl;
                        output = true;
                    }
                    break;
                }
            }else{
                scon.push_back(tmp);
            }
        }
       
        
        for(int i = 0; i<scon.size(); ++i)
        {
            if(scon[i] == "true")
            {
                con.push_back(1);
            }else if(scon[i] == "false")
            {
                con.push_back(0);
            }else if(scon[i] == "or")
            {
                if(i == scon.size()-1 || con.empty()){
                    if(!output)
                    {
                        cout<<"error"<<endl;
                        output = true;
                    }
                    break;
                }else{
                    bool last = con.back();
                    con.pop_back();
                    string tmp = scon[i+1];
                    if(!(tmp == "true" || tmp == "false"))
                    {
                        if(!output)
                        {
                            cout<<"error"<<endl;
                            output = true;
                        }
                        break;
                    }
                    bool llast = false;
                    if(tmp == "true")llast = true;
                    con.push_back((last|llast));
                    i++;
                }
            }
        }
        
       bool res = false;
        if(con.size() == 1 && !output)
        {
            res = con[0];
            if(res)
                cout<<"true"<<endl;
            else
                cout<<"false"<<endl;
        }else{
            if(!output)
            {
                cout<<"error"<<endl;
                output = true;
            }
        }
        
    }
    return 0;
}
*/

#include <iostream>
#include <string>

using namespace std;

bool isMatch(string p, string t)
{
    if(p.empty())
    {
        return t.empty();
    }
    if(t.empty())
    {
        int i = 0; 
        while(i<p.size())
        {
            if(p[i] != '*')
            {
                return false;
            }
                
            i++;
        }
        return true;
    }
    if(p[0] == t[0] || p[0] == '?')
        return isMatch(p.substr(1), t.substr(1));
    else if(p[0] == '*')
    {
        bool rval = false;
        rval |= isMatch(p.substr(1), t);
        rval |= isMatch(p, t.substr(1));
        return rval;
    }else{
        return false;
    }
    return true;
}

int main()
{
    string p, t;
    cin>>p>>t;
    cout<<isMatch(p, t)<<endl;
    return 0;
}
