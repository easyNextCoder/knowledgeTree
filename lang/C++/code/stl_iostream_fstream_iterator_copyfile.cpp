#include <iostream>
#include <fstream>
#include <iterator>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;
class si{
	public:
		bool i;
};
char buf[100];
int main(){
	vector<int>vec(10,2);
	ifstream File("test.rar",ios::binary);
	istream_iterator<int>eos;
	istream_iterator<int>read_iter(File);
	
	
	ofstream outFile("testx.rar", ios::binary|ios::out);
	
	
	/*//������ʹ�� 
	 while(!File.eof()){ //ʵ���ļ��ĸ���
	 cout<<"hello."<<endl;
	 cout<<File.gcount()<<endl;
	 cout<<sizeof(buf)<<endl;
                            File.read(buf, sizeof(buf));        
                            outFile.write(buf, File.gcount());
                        }
	File.close();
	outFile.close();
	*/
	
	
	/*
	bool c;
	while (File.get(c))  //ÿ�ζ�ȡһ���ַ�
        outFile.put(c);  //ÿ��д��һ���ַ�
    File.close();
    outFile.close();
	*/
	
	
	ostream_iterator<int>write_iter(outFile);
	
	while(read_iter != eos)
	{
		*write_iter = *read_iter;
		++read_iter;
		++write_iter;
	}
	
	return 0;
}



