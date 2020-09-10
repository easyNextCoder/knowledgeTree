#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
#include <set>
#include <map>
#include <unordered_map>
#include <functional>

using namespace std;


template<class T>
class test_template {
public:

    constexpr bool operator()(const T a, const T b) const
    {
        return a > b;
    }
};
vector<string> trulyMostPopular(vector<string>& names, vector<string>& synonyms) {
    map < string, int, greater<string> > mnames;
    unordered_map<string, string>syn;
    for (auto item : names)
    {
        int index = item.find_first_of('(');
        mnames.insert({ item.substr(0, index),
                         atoi(item.substr(index + 1, item.size() - 1 - (index + 1)).c_str()) });
    }

    for (auto item : synonyms)
    {
        int index = item.find_first_of(',');

        string name1 = item.substr(1, index - 1);
        string name2 = item.substr(index + 1, item.size() - 1 - (index + 1));

        //找到自己的最终祖先
        while (syn.find(name1) != syn.end())
            name1 = syn[name1];

        while (syn.find(name2) != syn.end())
            name2 = syn[name2];

        if (name1 != name2)
        {
            if (name1 > name2)swap(name1, name2);
            syn.insert({ name2, name1 });
            int times = mnames[name2] + mnames[name1];
            mnames.erase(name2);
            mnames[name1] = times;
        }
    }

    vector<string>out;
    auto iter = mnames.begin();

    while (iter != mnames.end())
    {
        string s = iter->first + "(" + to_string(iter->second) + ")";
        out.push_back(s);
        iter++;
    }

    return out;
}




int main()
{


    Solution solution;
    vector<string>names = { "John(15)","Jon(12)","Chris(13)","Kris(4)","Jo(10)","Johnny(4)","Christopher(19)" };
    vector<string>synonyms = { "(Jon,John)","(Johnny,Jo)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)" };
    vector<string> rval = solution.trulyMostPopular(names, synonyms);
    cout << "here" << endl;
    for (auto item : rval)
    {
        cout << "for" << endl;
        cout << item<< endl;
    }
    return 0;
}