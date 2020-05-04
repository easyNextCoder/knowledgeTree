#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    bool judege(vector<int> sequence, int start, int end) {
        if (end - start <= 1) {
            return true;
        }
        int root = sequence[start];
        int index = 0;
        for (int i = start + 1; i < end; i++) {
            if (root < sequence[i])
                continue;
            else {
                index = i;
                for (int j = i; j < end; j++) {
                    if (root > sequence[j]) {
                        if (j == end - 1) {
                            return judege(sequence, start + 1, index) && judege(sequence, index, end);
                        }
                        continue;
                    }
                    else {
                        return false;
                    }
                    
                }
            }
        }
    }
    bool VerifySquenceOfBST(vector<int> sequence) {

        reverse(sequence.begin(), sequence.end());
        if (sequence.size() == 0)
            return false;
        else {
            //
            return judege(sequence, 0, sequence.size());

        }
    }
};


int main() {

    Solution solution;
    vector<int> input = { 4,8,6,12,16,14,10 };
    cout << solution.VerifySquenceOfBST(input) << endl;;

    vector<int> input1 = { 4,8,6,12,16,14,10 };
    vector<int> input2 = { 46,86,6,124,164,144,104 };
    auto iter = find_first_of(input1.begin(), input1.end(), input2.begin(), input2.end());
    cout << *iter << endl;
}