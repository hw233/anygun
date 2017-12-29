#ifndef __ChannelBroadcaster_h__
#define __ChannelBroadcaster_h__

#include "ChannelConnection.h"
#include <set>
#include <map>

/** 针对一组 CHANNEL 发送广播消息的管理器.
	
	ChannelBroadcaster 管理一个 CHANNEL 的集合，如果通过 ChannelBroadcaster 
	进行消息的发送，这个消息会广播给这个集合中所有的 CHANNEL 对象。

	ChannelBroadcaster 建立在 ChannelConnection 的数据广播接口上。
*/
template<class CHANNEL>
class ChannelBroadcaster 
{
public:
	ChannelBroadcaster()
	{
	}

	virtual ~ChannelBroadcaster()
	{
	}

	/** 添加一个channel对象.
		@warning 这个channel对象必须有效，并且在被ChannelBroadcaster管理过程中
				 不能被关闭.
	*/
	void addChannel(CHANNEL* ch)
	{
		ACE_ASSERT(ch->isValid());
		ChannelConnection* conn = ch->getConn();
		ACE_ASSERT(conn);
		// 查找并添加这个channel对应的connection.
		std::pair<typename ConnToChanMap::iterator, bool> r = 
			connToChans_.insert(std::pair<ChannelConnection*, ChannelSet>(conn, ChannelSet()));
		r.first->second.insert(ch);
	}

	/** 移除一个channel对象. */
	void removeChannel(CHANNEL* ch)
	{
		ACE_ASSERT(ch->isValid());
		ChannelConnection* conn = ch->getConn();
		ACE_ASSERT(conn);
		typename ConnToChanMap::iterator iter = connToChans_.find(conn);
		ACE_ASSERT(iter != connToChans_.end());
		int r = iter->second.erase(ch);
		ACE_ASSERT(r == 1);
		if(iter->second.size() == 0)
			connToChans_.erase(iter);
	}

	/** 清除所有的channel对象. */
	void clearChannels()
	{
		connToChans_.clear();
	}

	/** 准备广播数据. 
		此操作为集合中的每个connection准备好channel集合，等待接下来调用 fillSendingData 
		填充广播数据.
	*/
	void initSendingData()
	{
		for(typename ConnToChanMap::iterator iter = connToChans_.begin(); iter != connToChans_.end(); ++iter)
		{
			ChannelConnection* conn = iter->first;
			conn->initBCSendingData((std::set<Channel*>&)iter->second);
		}
	}

	/** 填充广播数据. 
		此操作为集合中的每个connection填充好待广播的数据，等待接下来的 flushSendingData
		发送广播数据.
	*/
	void fillSendingData(void* data, size_t size)
	{
		for(typename ConnToChanMap::iterator iter = connToChans_.begin(); iter != connToChans_.end(); ++iter)
		{
			ChannelConnection* conn = iter->first;
			conn->fillSendingData(data, size);
		}
	}

	/** 发送广播数据. 
		此操作为集合中的每个connection做最终的数据发送，广播结束.
	*/
	void flushSendingData()
	{
		for(typename ConnToChanMap::iterator iter = connToChans_.begin(); iter != connToChans_.end(); ++iter)
		{
			ChannelConnection* conn = iter->first;
			conn->flushSendingData();
		}
	}

	typedef std::set<CHANNEL*> ChannelSet;
	typedef std::map<ChannelConnection*, ChannelSet>	ConnToChanMap;
	ConnToChanMap	connToChans_;
};

#endif//__ChannelBroadcaster_h__
