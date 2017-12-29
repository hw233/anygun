package bin
{
	import flash.utils.ByteArray;

	public interface IWriter
	{
		function writeBoolean(value:Boolean):void;
		function writeByte(value:int):void;
		function writeBytes(byte:ByteArray, offset:uint, length:uint):void;
		function writeDouble(value:Number):void;
		function writeFloat(value:Number):void;
		function writeInt(value:int):void;
		function writeMultiByte(value:String, charSet:String):void;
		function writeShort(value:int):void;
		function writeUnsignedInt(value:uint):void;
	}
}