#include <iostream>
#include <functional>

using namespace std;

class GameCharacter;

int defaultHealthCalc(const GameCharacter& gc)
{
	cout << "defaultHealthCalc func." << endl;
	return 0;
}

class GameCharacter {
public:
	typedef function<int(const GameCharacter&)> HealthCalFunc;
	explicit GameCharacter(HealthCalFunc hcf = defaultHealthCalc)
		: healthFunc(hcf)
	{

	}

	int healthValue() const
	{
		return healthFunc(*this);
	}

private:
	HealthCalFunc healthFunc;
};

class EvilBadGuy :public GameCharacter {
public:
	explicit EvilBadGuy(HealthCalFunc hcf = defaultHealthCalc)
		:GameCharacter(hcf)
	{

	}

};

class EyeCandyCharacter :public GameCharacter {
 public:
	 
	explicit EyeCandyCharacter(HealthCalFunc hcf = defaultHealthCalc)
	: GameCharacter(hcf)
	{
		//...
	}
};

int loseHealthQuickly(const GameCharacter& gcq)
{
	int i = 20;
	while (i > 0)
	{
		cout << "losing health " << i << "." << endl;
		i -= 5;
	}
	return 0;
}
int loseHealthSlowly(const GameCharacter& gcs)
{
	int i = 20;
	while (i > 0)
	{
		cout << "losing health " << i << "." << endl;
		i -= 2;
	}
	return 0;
}

short calcHealth(const GameCharacter&)
{
	cout << "calcHealth()" << endl;
	return 0;
}

struct HealthCalculator {
	int operator()(const GameCharacter&)  
	{
		cout << "HealthCalculator: operator." << endl;
		return 0;
	}
};

class GameLevel {
public:
	int health(const GameCharacter&) const
	{
		cout << "GameLevel: float health(const GameCharacter&) const" << endl;
		return 0;
	}
};

int main()
{
	EvilBadGuy moreEbg(loseHealthQuickly);
	moreEbg.healthValue();

	EvilBadGuy normalEbg(loseHealthSlowly);
	normalEbg.healthValue();

	GameLevel gl;
	EvilBadGuy ebg1(std::bind(&GameLevel::health, gl, placeholders::_1));
	ebg1.healthValue();

	EyeCandyCharacter ecc(HealthCalculator);//?不能使用
	


	/*//下面这段代码测试：如何使用指向类成员函数的指针
	GameCharacter normal;
	typedef int (HealthCalculator::* p)(const GameCharacter&);
	p s = &HealthCalculator::operator();
	HealthCalculator hCalculator;
	(hCalculator.*s)(normal);//前面加个括号很重要
	*/

	return 0;
}



