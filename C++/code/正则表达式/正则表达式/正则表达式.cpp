#include <iostream>
#include <regex>
#include <string>
#include <sstream>

using namespace std;
class Solution {
public:
    int strToInt(string str) {
        stringstream sm(str);
        if (str.size() == 0) {
            return 0;
        }
        string tmp;
        sm >> tmp;
        string tod;


        regex r("(^((\\s)*)?([+-]?))[0-9]*");
        smatch pre_results;
        smatch results;
        regex ir("[1-9][0-9]*");

        if (regex_search(tmp, pre_results, r) ) {
            tmp = pre_results.str();

            if (regex_search(tmp, results, ir)) {
                cout << results.str() << endl;
                long long int mulF = 1;
                long long int final_int = 0;
                string results_string = results.str();
                long long int max_int = INT_MIN;
                max_int = -max_int;
                for (int i = results_string.size() - 1; i >= 0; i--) {

                    long long int increment = mulF * (results_string[i] - '0');
                    
                    if (increment >= max_int) {
                        final_int = max_int;
                        break;
                    }
                    /*
                    if(increment/mulF != (results_string[i] - '0')){
                        final_int = INT_MAX;//overflow
                        break;
                    }
                    */
                    final_int += increment;

                    long long int mulF_old = mulF;
                    mulF *= 10;
                    

                }
                find_if(tmp.begin(), tmp.end(), [&final_int, &max_int](char s)->bool {
                    if (s == '-') {
                        if (final_int >= max_int)
                            final_int = INT_MIN;
                        else
                            final_int = -final_int;
                        return true;
                    }
                    return false;
                });
                if (final_int <= INT_MAX)
                    return (int)final_int;
                else
                    return INT_MAX;
            }
        }
        else {
            return 0;
        }
        return 0;

    }
};

class Solution2 {
public:
    int isNumber(char s) {
        if (s >= '0' && s <= '9')
            return true;
        return false;
    }
    int strToInt(string str) {
        int c = 0;
        int neg = 0;
        unsigned int cutoff = INT_MAX / 10;
        unsigned int cutlim = 0;
        unsigned int overflow = 0;
        unsigned int acc = 0;

        while (str[c] == ' ') { c++; }

        if (str[c] == '-') { neg = 1; c++; }
        else if (str[c] == '+') { neg = 0; c++; }

        if (neg) cutlim = 8;
        else cutlim = 7;

        for (acc = 0; c < str.size(); c++)
        {
            if (!isNumber(str[c]))
                break;

            if (overflow || acc > cutoff || (acc == cutoff && str[c] - '0' > cutlim)) {
                overflow = 1;
                break;
            }
            else {
                acc *= 10;
                acc += str[c] - '0';
            }
        }

        if (overflow)
            acc = neg ? (INT_MIN) : INT_MAX;
        else
            if (neg)acc = (~acc) + 1;

        return acc;

    }
};

int main() {
	/*
	string item("8a98080798989988aaajfkdjfkjasdljf");
	regex r("(^((\\s)*)?([+-]?))[0-9]*");
	smatch results;

	if (regex_search(item, results, r))
		cout << results.str() << endl;
	

	*/
    cout << atoi("-9879099096543210") << endl;
    Solution2 solution;
    cout<<solution.strToInt(" -9128347a");

	return 0;
}
