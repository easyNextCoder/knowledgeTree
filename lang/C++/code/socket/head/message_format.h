#pragma once
#include <string>
#include <stack>
#include <deque>
#define MESSAGE_LENGTH 140

using namespace std;

class message_body {
public:
	message_body() = default;
	message_body(int src_id, int des_id, string isms) {
		this->src_id = src_id;
		this->des_id = des_id;
		int i = 0;
		for (i = 0; i < isms.size(); i++) {
			_ms[i] = isms[i];
		}
		_ms[i] = '\0';

	}

	int src_id = -1;
	int des_id = -1;
	char _ms[MESSAGE_LENGTH];

};

class message_base {
public:

	message_base(message_body& ims) :ms(ims) {
		int i = 0; 
		for (i = 0; i < MESSAGE_LENGTH; i++) {
			if (ms._ms[i] == '\0')
				break;
		}
		ms_string = string(ms._ms, i);
	}

	message_base operator=(message_base& isms) {
		this->ms = isms.ms;
		this->ms_string = isms.ms_string;
		return *this;
	}

	int get_src_id() {
		return ms.src_id;
	}

	int get_des_id() {
		return ms.des_id;
	}

	string get_message() {
		return ms_string;
	}

private:
	message_body ms;
	string ms_string;

};