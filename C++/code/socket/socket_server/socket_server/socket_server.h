#pragma once
#include "../../head/message_format.h"
#include "../../head/output_format.h"
#include "../../head/constants.h"

#define MAX_THREAD_NUM 100
class server_message:public message_base {
	
public:
	server_message(message_body& msb):message_base(msb){}
	server_message(message_base& ms) :message_base(ms) {}

};