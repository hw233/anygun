#include "Service.h"


void Service::getParents(std::vector<Service*>& parents)
{
	if(super_)
		super_->getParents(parents);
	parents.push_back(this);
}

bool Service::findMethod( const std::string& name )
{
	for( size_t i = 0; i < methods_.size(); i++ )
		if( methods_[i].getName() == name )
			return true;
	return super_?super_->findMethod(name):false;
}

size_t Service::getMethodNum()
{
	int superMN = super_?super_->getMethodNum():0;
	return superMN + methods_.size();
}

void Service::getAllMethods(std::vector<Method*>& methods)
{
	if(super_)
		super_->getAllMethods(methods);
	for(size_t i = 0; i < methods_.size(); i++)
		methods.push_back(&(methods_[i]));
}