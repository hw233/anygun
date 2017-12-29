
import struct

class FieldMaskWriter:
    def __init__(self, n):
        self.mask = []
        while n > 0:
            self.mask.append(0)
            n -= 1
        self.pos = 0
    def set(self, b):
        if b:
            self.mask[self.pos>>3] |= (128 >> (self.pos & 0X00000007))
        self.pos += 1
    def write(self):
        return struct.pack('B'*len(self.mask), *self.mask)
        
def dynSizeWriter(b, s):
    if s <= 0X3F:
        b.append(struct.pack('>B', s))
    elif s <= 0X3FFF:
        s |= (1<<14)
        b.append(struct.pack('>H', s))
    elif s <= 0X3FFFFF:
        s |= (2<<22)
        b.append(struct.pack('>I', s)[1:])
    elif s <= 0X3FFFFFFF:
        s |= (3<<30)
        b.append(struct.pack('>I', s))

def numberWriter(f, b, v, fm):
    m = (v != 0)
    if fm:fm.set(m)
    if fm == None or m:
        b.append(struct.pack(f, v))
    
def int64Writer(b, v, fm):
    numberWriter('<q', b, v, fm)
def uint64Writer(b, v, fm):
    numberWriter('<Q', b, v, fm)
def doubleWriter(b, v, fm):
    numberWriter('<d', b, v, fm)
def floatWriter(b, v, fm):
    numberWriter('<f', b, v, fm)
def int32Writer(b, v, fm):
    numberWriter('<i', b, v, fm)
def uint32Writer(b, v, fm):
    numberWriter('<I', b, v, fm)
def int16Writer(b, v, fm):
    numberWriter('<h', b, v, fm)
def uint16Writer(b, v, fm):
    numberWriter('<H', b, v, fm)
def int8Writer(b, v, fm):
    numberWriter('<b', b, v, fm)
def uint8Writer(b, v, fm):
    numberWriter('<B', b, v, fm)
def boolWriter(b, v, fm):
    if fm:
        fm.set(v)
    else:
        if v:
            b.append('\001')
        else:
            b.append('\000')
def stringWriter(b, v, fm):
    l = len(v)
    if fm:
        if l:
            fm.set(True)
            dynSizeWriter(b, l)
            b.append(v)
        else:
            fm.set(False)
    else: 
        dynSizeWriter(b, l)
        b.append(v)
def enumWriter(b, v, fm):
    return uint8Writer(b, v, fm)

def write(wtr, a, b, v, fm):
    if a:
        l = len(v)
        m = l > 0;
        if fm:fm.set(m)
        if fm == None or m:
            dynSizeWriter(b, l)
            i = 0
            while i < l:
                wtr(b, v[i], None)
                i += 1
    else:
        return wtr(b, v, fm)