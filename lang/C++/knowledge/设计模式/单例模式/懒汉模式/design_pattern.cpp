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
    static Singleton* m_instance;  
    Singleton(){}  
public:  
    static Singleton* getInstance();  
};  
  
Singleton* Singleton::getInstance()  
{  
    if(NULL == m_instance)  
    {  
        Lock();//借用其它类来实现，如boost  
        if(NULL == m_instance)  
        {  
            m_instance = new Singleton;  
        }  
        UnLock();  
    }  
    return m_instance;  
} 

int main()
{
    
    return 0;
}