#pragma once
#include "MessageFormat.h"

#define MAX_THREAD_NUM 100
class ServerMessage:public MessageBody
{
public:
	ServerMessage(MessageBody& msb):MessageBody(msb){}
	//ServerMessage(MessageBase& ms):MessageBase(ms){}
};
