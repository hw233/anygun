#if defined (_MSC_VER)
	// Disable vc warnings.
	#define _CRT_SECURE_NO_WARNINGS
	#pragma warning(disable:4996)
	#pragma warning(disable:4267)
	#pragma warning(disable:4129)//unrecognized character escape sequence
	#pragma warning(disable:4819)//The file contains a character that cannot be represented in the current code page (936). Save the file in Unicode format to prevent data loss
#endif

#include "Channel.h"
#include "ChannelConnection.h"

Channel::Channel():
guid_(0),
conn_(NULL)
{
}

Channel::~Channel()
{
}

bool Channel::isValid()
{
	return (conn_ == NULL)?false:true;
}

void Channel::fillBegin()
{
	if(conn_)
		conn_->initChannelSendingData(this);
}

void Channel::fill(void* data, size_t size)
{
	if(conn_)
		conn_->fillSendingData(data, size);
}

void Channel::fillEnd()
{
	if(conn_)
		conn_->flushSendingData();
}

bool Channel::handleClose()
{
	conn_ = NULL;
	return true;
}
