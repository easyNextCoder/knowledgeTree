#pragma once
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;

void output(vector<pair<string,string>>inform_pair) {
	for_each(inform_pair.begin(), inform_pair.end(), [](pair<string, string>item) {
		cout << item.first << ": " << item.second << endl;
		});
}