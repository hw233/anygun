package bin
{
	import flash.utils.ByteArray;
	import flash.utils.Endian;
	
	public class ProtocolWriter
	{
		public static function writeBytes(w:IWriter, b:ByteArray):void
		{
			w.writeBytes(b, 0, b.length);
		}
		public static function writeInt64(w:IWriter, v:Int64):void
		{
			w.writeUnsignedInt(v.h_);
			w.writeUnsignedInt(v.l_);
		}
		public static function writeDouble(w:IWriter, v:Number):void
		{
			w.writeDouble(v);
		}
		public static function writeFloat(w:IWriter, v:Number):void
		{
			w.writeFloat(v);
		}
		public static function writeInt32(w:IWriter, v:int):void
		{
			w.writeInt(v);
		}
		public static function writeUInt32(w:IWriter, v:uint):void
		{
			w.writeUnsignedInt(v);
		}
		public static function writeInt16(w:IWriter, v:int):void
		{
			w.writeShort(v);
		}
		public static function writeUInt16(w:IWriter, v:uint):void
		{
			w.writeShort(v);
		}
		public static function writeInt8(w:IWriter, v:int):void
		{
			w.writeByte(v);
		}
		public static function writeUInt8(w:IWriter, v:uint):void
		{
			w.writeByte(v);
		}
		public static function writeBool(w:IWriter, v:Boolean):void
		{
			w.writeBoolean(v);
		}
		public static function writeEnum(w:IWriter, v:uint):void
		{
			w.writeByte(v);
		}
		public static function writeString(w:IWriter, v:String):void
		{
			var temp:ByteArray = new ByteArray();
			temp.position = 0;
			temp.writeUTFBytes(v);
			writeDynSize(w, temp.length);
			w.writeBytes(temp, 0, temp.length);
		}
		public static function writeDynSize(w:IWriter, s:uint):void
		{
			var temp:ByteArray=new ByteArray;
			temp.endian = Endian.LITTLE_ENDIAN;
			temp.writeUnsignedInt(s);
			var n:uint=0;
			if(s <= 0X3F)
				n = 0;
			else if(s <= 0X3FFF)
				n = 1;
			else if(s <= 0X3FFFFF)
				n = 2;
			else if(s <= 0X3FFFFFFF)
				n = 3;
			temp[n] |= (n<<6);
			for(var i:int = n; i >= 0; i--)
				w.writeByte(temp[i]);
		}
	}
}