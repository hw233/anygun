#include "Enum.h"

bool Enum::findItem( const std::string& item )
{
	for( size_t i = 0; i < items_.size(); i++ )
		if( items_[i] == item )
			return true;
	return false;
}