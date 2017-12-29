package bin
{
	import flash.utils.ByteArray;
	/**
	 * Field Mask.
	 */
	public class FieldMask
	{
		public function readBit():Boolean
		{
			var bytePos:uint = bitPos_ >> 3;
			var  r:Boolean = ((masks_[bytePos] & (128 >> (int)(bitPos_ & 0X00000007))) != 0)? true : false;
			bitPos_++;
			return r;			
		}
		public function writeBit(b:Boolean):void
		{
			var bytePos:uint = bitPos_ >> 3;
			if(bytePos >= masks_.length)
				masks_.writeByte(0);
			if(b)
				masks_[bytePos] |= (128>>(bitPos_&0X00000007));
			bitPos_++;
		}
		public function getMasks():ByteArray
		{
			return masks_;
		}
		private var masks_:ByteArray = new ByteArray();
		private var bitPos_:uint =0;
	}
}
