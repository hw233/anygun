#include "Struct.h"

bool Struct::findField( const std::string& name )
{	
	if(FieldContainer::findField(name))
		return true;
	return super_?super_->findField(name):false;
}

size_t Struct::getFieldNum()
{
	return super_?super_->getFieldNum() + fields_.size():fields_.size();
}