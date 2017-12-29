package bin
{
	import flash.utils.ByteArray;
	public class Int64
	{
		public var h_:uint = 0;
		public var l_:uint = 0;
		public function isZero():Boolean
		{
			return h_ == 0 && l_ == 0;
		}
		public static function equal(srcObj:Int64,destObj:Int64):Boolean
		{
			return (srcObj.h_ == destObj.h_) && (srcObj.l_ == destObj.l_);
		}
	}
}