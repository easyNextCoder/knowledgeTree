#pragma once
#include <string>
#include <stack>
#include <deque>
#include <string.h>
#include <iostream>



#define MESSAGE_LENGTH 140

using namespace std;

class MessageBody
{
public:
	MessageBody() = default;
	MessageBody(int src_id, int des_id, string isms)
	{
		this->src_id = src_id;
		this->des_id = des_id;
		const char* src = isms.c_str();		
			
		strcpy(_ms, src);
	}
	
	MessageBody(MessageBody& msb)
	{
		this->src_id = msb.src_id;
		this->des_id = msb.des_id;
		string tmp = msb.get_message();
		const char* src = tmp.c_str();
		strcpy(_ms, src);
		
	}
	void set_src_id(int value)
	{
		src_id = value;
	}
	void set_des_id(int value)
	{
		des_id = value;
	}
	int get_src_id()
	{
		return src_id;
	}
	int get_des_id()
	{
		return des_id;
	}
	string get_message()
	{
		string rval;
		rval  = _ms;
		return rval;	
	}
	
	
private:

	int src_id = -1;
	int des_id = -1;
	char _ms[MESSAGE_LENGTH];
	int messageLen = 0;
};

/*

class MessageBase
{
public:
	MessageBase(MessageBody& ims):ms(ims)
	{
		int i = 0;
		for(i = 0; i<MESSAGE_LENGTH; i++)
		{
			if(ms._ms[i] == '\0')
				break;
		}
		
		this->ms_string.copy(ms._ms, i, 0);
	}

	MessageBase(MessageBase& mbs):ms(mbs.ms),ms_string(mbs.ms_string){}

	int get_src_id()
	{
		return ms.src_id;
	}

	int get_des_id()
	{
		return ms.des_id;
	}

	string get_message()
	{
		return ms_string;
	}


	MessageBody ms;
	string ms_string;
};

*/
