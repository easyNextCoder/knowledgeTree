	string str;
    while(getline(cin,str))
    {
        stringstream sstr(str);
        string  token[2], tmp;
        int count = 0;
		//4 4 4 4-joker JOKER下面一行代码将这个例子分为两个部分
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
	...