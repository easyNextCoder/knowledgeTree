## 在for循环进行判断时容易产生的bug

for(int i = 0; i+patternRight.size()-1<str.size(); i++)
这里中间那个不要使用-1，而应该改成
for(int i = 0; i+patternRight.size()<str.size()-1; i++)
因为0+(unsigned int类型的)0-1会溢出成一个非常大的数字