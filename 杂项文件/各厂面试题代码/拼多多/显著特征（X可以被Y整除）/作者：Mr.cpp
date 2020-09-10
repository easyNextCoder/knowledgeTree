/*
    题目：给定两个数N,M；接下来是M行，求1-N之间的个数可以被集合中的某些（从1-m个）元素整除的个数
    输入数据：10 2
              2
              3
    输出数据：7
    解题思路：这道题要用到容斥原理https://blog.csdn.net/sodacoco/article/details/81635420
             基本思想就是，先全部计算出来，然后将重复的部分剪掉
*/

#include <iostream>
#include <sstream>
#include <string>
#include <algorithm>
#include <queue>
#include <stdio.h>
#include <limits.h>
#include <map>
#include <string.h>
#include <stack>
#include <cmath>
#include <iomanip>
#include <assert.h>
#include <limits.h>
#include <set>
#include <unordered_map>

#define debug

using namespace std;

typedef long long LL ;

LL arr[20];
LL n, m;

 LL gcd(LL a,LL b) //求两个数的最大公约数
 {
     return b ? gcd(b,a%b):a;
}
  LL lcm(LL m,LL g)  //求两个数的最小公倍数
  {
     return m/gcd(m,g)*g;
 }

LL lcmmulp(vector<LL>& tmp) {
    if (tmp.size() == 1) return tmp[0];
    else if (tmp.size() == 2) return lcm(tmp[0], tmp[1]);
    else {
        LL ret = lcm(tmp[0], tmp[1]);
        for (int i = 2; i < tmp.size(); i++) {
            ret = lcm(ret, tmp[i]);
        }
        return ret;
    }
}

LL solve(LL val) {
    vector<LL> tmp;
    int idx = 0;
    while (val > 0) {
        if ((val & 1) == 1) {
            tmp.push_back(arr[idx]);
        }
        idx++;
        val = val >> 1;
    }
    LL flag = tmp.size() % 2 ? 1 : -1;
    int ret = lcmmulp(tmp);
    return (n / ret) * flag;
}


int main() {


    cin >> n >> m;

    for (int i = 0; i < m; i++) cin >> arr[i];
    
    LL ans = 0;
    
    int eval = (1 << m) - 1;
    
    for (int val = 1; val <= eval; val++) {
        ans += solve(val);
    }
    
    printf("%d", ans);
    
    return 0;

}
