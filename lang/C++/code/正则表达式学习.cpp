#include <iostream>
#include <regex>
#include <string>

using namespace std;

int main(){
	string item("8a98080798989988aaajfkdjfkjasdljf");
	regex r("(^((\\s)*)?([+-]?))[0-9]*");
	smatch results;
	
	if(regex_search(item, results, r))
		cout<<results.str()<<endl;
		
		 
	
	
	return 0;
}
