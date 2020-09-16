//返回值是tok最大为k的获利最大值
int f(int tok, int index, int have){

    if(tok == k || index == prices.size())return 0;//持有到最后还没有卖，就没有意义了
    int maxOut = 0;
    //决定是买还是卖的问题
    if(have)
    {
        int ravl1 = 0;
        int rval2 = 0;
        //继续持有
        rval1 = f(tok, index+1, have);
        //卖出
        rval2 = f(tok+1, index+1, !have)+prices[index];
        maxOut = max(maxOut, max(rval1, rval2));
    }else{
        //现在决定是买还是不买的问题
        int rval1 = 0;
        //不买
        rval1 = f(tok, index+1, have);
        //买了
        int rval2 = f(tok, index+1, !have)-prices[i];
        maxOut = max(maxOut, max(rval1, rval2));
    }
    return maxOut;

}

void maxProfit(int k, vector<int>& prices)
{


}