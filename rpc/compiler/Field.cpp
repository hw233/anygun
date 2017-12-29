#include "Field.h"

bool FieldContainer::findField( const std::string& name )
{
	for( size_t i = 0; i < fields_.size(); i++ )
		if( fields_[i].getName() == name )
			return true;
	return false;
}