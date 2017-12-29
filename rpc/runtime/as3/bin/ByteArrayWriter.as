package bin
{
	import flash.utils.ByteArray;
	import flash.utils.Endian;

	/** Write bytes to a ByteArray. */
	public class ByteArrayWriter implements IWriter
	{
		private var ba_:ByteArray;
		public function ByteArrayWriter()
		{
			ba_ = new ByteArray();
			ba_.endian = Endian.LITTLE_ENDIAN;
			ba_.position = 0;
		}
		public function getByteArray():ByteArray
		{
			return ba_;
		}
		public function writeBoolean(value:Boolean):void
		{
			ba_.writeBoolean(value);
		}
		public function writeByte(value:int):void
		{
			ba_.writeByte(value);
		}
		public function writeBytes(byte:ByteArray, offset:uint, length:uint):void
		{
			ba_.writeBytes(byte, offset, length);
		}
		public function writeDouble(value:Number):void
		{
			ba_.writeDouble(value);
		}
		public function writeFloat(value:Number):void
		{
			ba_.writeFloat(value);
		}
		public function writeInt(value:int):void
		{
			ba_.writeInt(value);
		}
		public function writeMultiByte(value:String, charSet:String):void
		{
			ba_.writeMultiByte(value, charSet);
		}
		public function writeShort(value:int):void
		{
			ba_.writeShort(value);
		}
		public function writeUnsignedInt(value:uint):void
		{
			ba_.writeUnsignedInt(value);
		}
	}
}