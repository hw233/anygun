// import other file.
#import Import.arpc

/* Structure inheritance. */
struct DerivedStruct : StructBase
{
	int32 aaa_;
};

/* Service inheritance. */
service DerivedService : ServiceBase
{
	method3(DerivedStruct d);
	method4();
};
