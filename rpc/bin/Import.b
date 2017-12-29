// Single line comment.
/* Mutiline comment.
*/

/* Enumeration Definition. */
enum EnumName
{
	EN1,
	EN2,
	EN3,
};

struct StructType
(skipcomp)	// Skip data compression.
{
	string aaa_;
	int32 bbb_;
};

/* Structure Definition.*/
struct StructBase
{
	/* Basic types */
	double 	double_;
	float	float_;
	int64	int64_;
	uint64	uint64_;
	int32	int32_;
	uint32	uint32_;
	int16	int16_;
	uint16	uint16_;
	int8	int8_;
	uint8	uint8_;
	bool	bool_;
	/* Enum type. */
	EnumName enum_;
	/* struct type */
	StructType struct_;
	/* String type.
		length is 65535.
	*/
	string	string_;
	// String with explicit length.
	string[32]	string1_;
	/* Array type. 
		Array type can contain any other types(including user defined structure or enumeration types) other than array.
		length is 65535.
	*/
	array<double> 	doubleArray_;
	array<float>	floatArray_;
	array<int64>	int64Array_;
	array<uint64>	uint64Array_;
	array<int32>	int32Array_;
	array<uint32>	uint32Array_;
	array<int16>	int16Array_;
	array<uint16>	uint16Array_;
	array<int8>		int8Array_;
	array<uint8>	uint8Array_;
	array<bool>		boolArray_;
	array<string> 	strArray_;
	array<EnumName> enumArray_;
	array<StructType> structArray_;
	// Array with explicit length.
	array[8]<string[16]> strarray1_;
	/* Byte array type.
		This type equals array<uint8>.
	*/
	bytes	bytes_;
	// Byte array with explicit length.
	bytes[64] bytes11_;
};

/* Service type. */
service ServiceBase
{
	method1(string[32] username, string[32] password);
	method2(StructBase s);
};

