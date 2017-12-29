#ifndef __ProtocolMemReader_h__
#define __ProtocolMemReader_h__

#include "Common.h"
#include "ProtocolReader.h"

/** 从一个指定内存区读取数据的读取器实现.
*/
class ProtocolMemReader : public ProtocolReader
{
public:
	ProtocolMemReader(void* b, size_t l):
	buffer_((char*)b),
	length_(l),
	rdptr_(0)
	{}

	virtual bool read(void* data, size_t len)
	{
		if(length_ < rdptr_ + len)
			return false;
		::memcpy(data, buffer_ + rdptr_, len);
		rdptr_ += len;
		return true;
	}

private:
	char*			buffer_;
	size_t			length_;
	size_t			rdptr_;
};


#endif//__ProtocolMemReader_h__