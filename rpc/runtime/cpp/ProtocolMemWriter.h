#ifndef __ProtocolMemWriter_h__
#define __ProtocolMemWriter_h__
#include "Common.h"
#include "ProtocolWriter.h"
/** 从一个指定内存区读取数据的读取器实现.
*/
class ProtocolMemWriter : public ProtocolWriter
{
public:
	ProtocolMemWriter(void* b, size_t l):
	buffer_((char*)b),
	length_(l),
	wtptr_(0)
	{}

	virtual void write(const void* data, size_t len)
	{
		if(length_ < wtptr_ + len)
			return;
		::memcpy(buffer_ + wtptr_, data, len);
		wtptr_ += len;
	}

private:
	char*			buffer_;
	size_t			length_;
	size_t			wtptr_;
};


#endif//__ProtocolMemWriter_h__