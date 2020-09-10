#include <iostream>
#include <vector>
#include <mutex>


using namespace std;

mutex m;

void Lock()
{
    m.lock();
}

void UnLock()
{
    m.unlock();
}

class Singleton  
{  
private:  
    static Singleton m_instance;  
    Singleton(){
        cout<<"construct Singleton"<<endl;
    }  
    ~Singleton(){
        cout<<"destroy Singleton"<<endl;
    }

public:  
    static Singleton* getInstance();  
};  

Singleton Singleton::m_instance; 

Singleton* Singleton::getInstance()  
{  
    return &m_instance;  
} 

int main()
{
    cout<<stod("10")<<endl;
    //cout<<(Singleton::getInstance())<<endl;
    //cout<<(Singleton::getInstance())<<endl;
    
    return 0;
}