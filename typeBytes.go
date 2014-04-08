package typeBytes

import (
	"github.com/virtao/GoEndian"
	"unsafe"
)

const (
	INT_SIZE     int = int(unsafe.Sizeof(0))
	UINT_SIZE    int = int(unsafe.Sizeof(uint(0)))
	FLOAT32_SIZE int = 4
	FLOAT64_SIZE int = 8
)

func IntToBytes(n int) (ret []byte) {
	return xToBytes(unsafe.Pointer(&n), INT_SIZE)
}

func BytesToInt(n []byte) (ret int) {
	p := bytesToX(n, INT_SIZE)
	ret = *((*int)(p))
	return ret
}

func UintToBytes(n uint) (ret []byte) {
	return xToBytes(unsafe.Pointer(&n), UINT_SIZE)
}

func BytesToUint(n []byte) (ret uint) {
	p := bytesToX(n, UINT_SIZE)
	ret = *((*uint)(p))
	return ret
}

func Float32ToBytes(n float32) (ret []byte) {
	return xToBytes(unsafe.Pointer(&n), FLOAT32_SIZE)
}

func BytesToFloat32(n []byte) (ret float32) {
	p := bytesToX(n, INT_SIZE)
	ret = *((*float32)(p))
	return ret
}

func Float64ToBytes(n float64) (ret []byte) {
	return xToBytes(unsafe.Pointer(&n), FLOAT64_SIZE)
}

func BytesToFloat64(n []byte) (ret float64) {
	p := bytesToX(n, FLOAT64_SIZE)
	ret = *((*float64)(p))
	return ret
}

//np = 原始数据地址，size = 原始数据长度
func xToBytes(np unsafe.Pointer, size int) (ret []byte) {
	ret = make([]byte, size)
	for i := 0; i < size; i++ {
		ret[i] = *((*byte)(unsafe.Pointer((uintptr(np) + uintptr(i)))))
	}

	return ret
}

//n = 原始数据，size = 目标数据长度
func bytesToX(n []byte, size int) (ret unsafe.Pointer) {
	ns := len(n)
	bytes := make([]byte, size)
	if ns > size { //数组长度太大，截取前size个字节
		copy(bytes, n[:size])
	} else { //数组长度太小，根据大小端字节序保存，对于小端存储，原始字节顺序必须确定与小端存储顺序一致（高地址存低位，或者数组的高元素编号存低位数）
		if endian.IsBigEndian() {
			copy(bytes[size-ns:], n)
		} else {
			copy(bytes[:ns], n)
		}
	}

	ret = getSliceArrayPointer(unsafe.Pointer(&bytes))
	return ret
}

func getSliceArrayPointer(p unsafe.Pointer) (ret unsafe.Pointer) {
	return unsafe.Pointer(*(*uintptr)(p))
}

func Int16ToBytes(n int16) (ret []byte) {
	return Uint16ToBytes(uint16(n))
}
func BytesToInt16(buf []byte) (ret int16) {
	return int16(BytesToUint16(buf))
}
func Uint16ToBytes(n uint16) (ret []byte) {
	ret = make([]byte, 2)
	endian.Endian.PutUint16(ret, n)
	return ret
}
func BytesToUint16(buf []byte) (ret uint16) {
	return endian.Endian.Uint16(buf)
}

func Int32ToBytes(n int32) (ret []byte) {
	return Uint32ToBytes(uint32(n))
}
func BytesToInt32(buf []byte) (ret int32) {
	return int32(BytesToUint32(buf))
}
func Uint32ToBytes(n uint32) (ret []byte) {
	ret = make([]byte, 4)
	endian.Endian.PutUint32(ret, n)
	return ret
}
func BytesToUint32(buf []byte) (ret uint32) {
	return endian.Endian.Uint32(buf)
}

func Int64ToBytes(n int64) (ret []byte) {
	return Uint64ToBytes(uint64(n))
}
func BytesToInt64(buf []byte) (ret int64) {
	return int64(BytesToUint64(buf))
}
func Uint64ToBytes(n uint64) (ret []byte) {
	ret = make([]byte, 8)
	endian.Endian.PutUint64(ret, n)
	return ret
}
func BytesToUint64(buf []byte) (ret uint64) {
	return endian.Endian.Uint64(buf)
}
