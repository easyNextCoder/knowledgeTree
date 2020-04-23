#pragma once
#include "../../head/message_format.h"
#include "../../head/output_format.h"
#include "../../head/constants.h"

#define MAX_THREAD_NUM 100
class client_message:public message_base {
	
public:
	client_message(message_body& msb):message_base(msb){}
	client_message(message_base& ms) :message_base(ms) {}

};
