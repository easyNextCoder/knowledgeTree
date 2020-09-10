#include <iostream>

using namespace std;

class CostEstimate{
    private:
    static const double FudgeFactor;
};

const double CostEstimate::FudgeFactor = 1.35;


int main()
{
    CostEstimate cet;
    return 0;
}