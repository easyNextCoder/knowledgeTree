#include <iostream>
#include <string>

using namespace std;

class File {
	public:
		void setFileName(string fn)
		{
			fileName = fn;	
		} 
		string fileName;
};

class InputFile:virtual File {};

class OutputFile:virtual File {};

class IOFile:public InputFile, public OutputFile {

};

int main() {
	IOFile iof;
	//cout<<iof.InputFile::fileName<<endl;
	cout<<iof.setFileName()<<endl;
	return 0;
}
