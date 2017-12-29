
import struct

class FieldMaskReader:
    def __init__(self, b, p, n):
        self.pos = 0
        self.mask = []
        while n > 0:
            self.mask.append(struct.unpack('B', b[p:p+1])[0])
            p += 1
            n -= 1
    def get(self):
        b = (self.mask[self.pos>>3] & (128 >> (self.pos & 0X00000007)) and True) or False
        self.pos += 1
        return b
    
def dynSizeReader(b, p):
    bb = struct.unpack('B', b[p:p+1])[0]
    p += 1
    n = (bb & 0XC0)>>6
    s = (bb & 0X3F)
    while n > 0:
        bb = struct.unpack('B', b[p:p+1])[0]
        p += 1
        s = (s<<8)|bb
        n -= 1
    return s, p

def numberReader(f, n, b, p, fm):
    if fm == None or fm.get():
        return struct.unpack(f, b[p:p+n])[0], p+n
    else:
        return 0, p
    
def int64Reader(b, p, valMax, fm):
    return numberReader('<q', 8, b, p, fm)
def uint64Reader(b, p, valMax, fm):
    return numberReader('<Q', 8, b, p, fm)
def doubleReader(b, p, valMax, fm):
    return numberReader('<d', 8, b, p, fm)
def floatReader(b, p, valMax, fm):
    return numberReader('<f', 4, b, p, fm)
def int32Reader(b, p, valMax, fm):
    return numberReader('<i', 4, b, p, fm)
def uint32Reader(b, p, valMax, fm):
    return numberReader('<I', 4, b, p, fm)
def int16Reader(b, p, valMax, fm):
    return numberReader('<h', 2, b, p, fm)
def uint16Reader(b, p, valMax, fm):
    return numberReader('<H', 2, b, p, fm)
def int8Reader(b, p, valMax, fm):
    return numberReader('<b', 1, b, p, fm)
def uint8Reader(b, p, valMax, fm):
    return numberReader('<B', 1, b, p, fm)
def boolReader(b, p, valMax, fm):
    if fm:
        return fm.get(), p
    else:
        if b[p] == '\000':
            return False, p+1
        else:
            return True, p+1
def stringReader(b, p, valMax, fm):
    if fm == None or fm.get():
        l, p = dynSizeReader(b, p)
        if l > valMax:
            raise
        return b[p:p+l], p+l
    else:
        return "", p
def enumReader(b, p, valMax, fm):
    e, p = uint8Reader(b, p, 0, fm)
    if e > valMax:
        raise
    return e, p

def read(rdr, b, p, arrMax, valMax, fm):
    if arrMax > 0:
        if fm == None or fm.get():
            s, p = dynSizeReader(b, p)
            if s > arrMax:
                raise
            else:
                arr = []
                while s > 0:
                    i, p = rdr(b, p, valMax, None)
                    arr.append(i)
                    s -= 1
                return arr, p
        else:
            return [], p
    else:
        return rdr(b, p, valMax, fm)