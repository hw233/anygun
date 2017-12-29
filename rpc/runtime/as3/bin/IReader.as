package bin
{
	import flash.utils.ByteArray;
	
	public interface IReader
	{
		function readBoolean():Boolean;
		function readByte():int;
		function readBytes(bytes:ByteArray, offset:uint, length:uint):void;
		function readDouble():Number;
		function readFloat():Number;
		function readInt():int;
		function readShort():int;
		function readUnsignedByte():uint;
		function readUnsignedInt():uint;
		function readUnsignedShort():uint;
		function readUTFBytes(length:uint):String;
	}
}