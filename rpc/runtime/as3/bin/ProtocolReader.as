package bin
{
	import flash.utils.ByteArray;

	public class ProtocolReader
	{
		public static function readBytes(r:IReader, b:ByteArray, length:uint):void
		{
			r.readBytes(b, 0, length);
		}
		public static function readInt64(r:IReader):Int64
		{
			var i:Int64 = new Int64();
			i.h_ = r.readUnsignedInt();
			i.l_ = r.readUnsignedInt();
			return i;
		}
		public static function readDouble(r:IReader):Number
		{
			return r.readDouble();
		}
		public static function readFloat(r:IReader):Number
		{
			return r.readFloat();
		}
		public static function readInt32(r:IReader):int
		{
			return r.readInt();
		}
		public static function readUInt32(r:IReader):uint
		{
			return r.readUnsignedInt();
		}
		public static function readInt16(r:IReader):int
		{
			return r.readShort();
		}
		public static function readUInt16(r:IReader):uint
		{
			return r.readUnsignedShort();
		}
		public static function readInt8(r:IReader):int
		{
			return r.readByte();
		}
		public static function readUInt8(r:IReader):uint
		{
			return r.readUnsignedByte();
		}
		public static function readBool(r:IReader):Boolean
		{
			return r.readBoolean()
		}
		public static function readEnum(r:IReader, max:uint):uint
		{
			var e:uint = r.readUnsignedByte();
			if(e >= max)
				throw "Invalid enum";
			return e;
		}
		public static function readString(r:IReader, maxLen:uint):String
		{
			var len:uint = readDynSize(r, maxLen);
			return r.readUTFBytes(len);		}
		
		public static function readDynSize(r:IReader, max:uint):uint
		{
			var b:uint = readUInt8(r);
			var n:uint = (b & 0XC0)>>6;
			var s:uint = (b & 0X3F);
			for(var i:uint = 0; i < n; i++)
			{
				b = readUInt8(r);
				s = (s<<8)|b;
			}
			if(s > max) throw "invalid size";
			return s;
		}
	}
}