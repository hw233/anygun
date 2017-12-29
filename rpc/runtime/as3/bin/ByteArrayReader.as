package bin
{
	import flash.utils.ByteArray;
	import flash.utils.Endian;
	
	/** Read bytes from a ByteArray. */
	public class ByteArrayReader implements IReader
	{
		private var ba_:ByteArray;
		public function ByteArrayReader(ba:ByteArray)
		{
			ba_ = ba;
			ba_.endian = Endian.LITTLE_ENDIAN;
			ba_.position = 0;
		}
		public function readBoolean():Boolean
		{
			return ba_.readBoolean();
		}
		public function readByte():int
		{
			return ba_.readByte();
		}
		public function readBytes(bytes:ByteArray, offset:uint, length:uint):void
		{
			ba_.readBytes(bytes, offset, length);
		}
		public function readDouble():Number
		{
			return ba_.readDouble();
		}
		public function readFloat():Number
		{
			return ba_.readFloat();
		}
		public function readInt():int
		{
			return ba_.readInt();
		}
		public function readShort():int
		{
			return ba_.readShort();
		}
		public function readUnsignedByte():uint
		{
			return ba_.readUnsignedByte();
		}
		public function readUnsignedInt():uint
		{
			return ba_.readUnsignedInt();
		}
		public function readUnsignedShort():uint
		{
			return ba_.readUnsignedShort();
		}
		public function readUTFBytes(length:uint):String
		{
			return ba_.readUTFBytes(length);
		}
	}
}