GoTypeBytes
===========

A basic type of conversion tool for Go. It can achieve mutual conversion between primitive types and arrays of bytes. 

Types that can be converted:

int/int16/int32/int64/uint/uint16/uint32/uint64/float32/float64 <=> []byte

Sample code:



```
package main

import (
    "encoding/hex"
    "fmt"
    "github.com/virtao/GoTypeBytes"
)

func main() {
    fmt.Println("typeBytes库测试：")
    typeBytesTest()
}

func typeBytesTest() {
    fmt.Println("float32 size : ", typeBytes.FLOAT32_SIZE)
    fmt.Println("float64 size : ", typeBytes.FLOAT64_SIZE)
    fmt.Println("int size : ", typeBytes.INT_SIZE)
    fmt.Println("uint size : ", typeBytes.UINT_SIZE)

    var testInt int = 0x1f123400ff123400
    var testInt16 int16 = 0x1f12
    var testInt32 int32 = 0x1f123400
    var testInt64 int64 = 0x1f123400ff123400
    var testUint uint = 0xff123400ff123400
    var testUint16 uint16 = 0xff12
    var testUint32 uint32 = 0xff123400
    var testUint64 uint64 = 0xff123400ff123400
    var testFloat32 float32 = -1234.5678 //x86_64 PC上存储形式应该为：2b 52 9a c4
    var testFloat64 float64 = -1234.5678 //x86_64 PC上存储形式应该为：ad fa 5c 6d 45 4a 93 c0

    buf := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    fmt.Printf("IntToBytes : 0x%0x to %v\n", testInt, hex.EncodeToString(typeBytes.IntToBytes(testInt)))
    fmt.Printf("Int16ToBytes : 0x%0x to %v\n", testInt16, hex.EncodeToString(typeBytes.Int16ToBytes(testInt16)))
    fmt.Printf("Int32ToBytes : 0x%0x to %v\n", testInt32, hex.EncodeToString(typeBytes.Int32ToBytes(testInt32)))
    fmt.Printf("Int64ToBytes : 0x%0x to %v\n", testInt64, hex.EncodeToString(typeBytes.Int64ToBytes(testInt64)))
    fmt.Printf("UintToBytes : 0x%0x to %v\n", testUint, hex.EncodeToString(typeBytes.UintToBytes(testUint)))
    fmt.Printf("Uint16ToBytes : 0x%0x to %v\n", testUint16, hex.EncodeToString(typeBytes.Uint16ToBytes(testUint16)))
    fmt.Printf("Uint32ToBytes : 0x%0x to %v\n", testUint32, hex.EncodeToString(typeBytes.Uint32ToBytes(testUint32)))
    fmt.Printf("Uint64ToBytes : 0x%0x to %v\n", testUint64, hex.EncodeToString(typeBytes.Uint64ToBytes(testUint64)))
    fmt.Printf("Float32ToBytes : %f to %v\n", testFloat32, hex.EncodeToString(typeBytes.Float32ToBytes(testFloat32)))
    fmt.Printf("Float64ToBytes : %f to %v\n", testFloat64, hex.EncodeToString(typeBytes.Float64ToBytes(testFloat64)))

    //数组长度刚好
    fmt.Printf("BytesToInt(8Bytes) : %s to 0x%0x\n", hex.EncodeToString(buf[2:10]), typeBytes.BytesToInt(buf[2:10]))
    //数组长度太大，截取前INT_SIZE个字节
    fmt.Printf("BytesToInt(>8Bytes) : %s to 0x%0x\n", hex.EncodeToString(buf), typeBytes.BytesToInt(buf))
    //数组长度太小，根据大小端字节序保存
    fmt.Printf("BytesToInt(2Bytes) : %s to 0x%0x\n", hex.EncodeToString(buf[:2]), typeBytes.BytesToInt(buf[:2]))
}
```

Output:

```
    typeBytes库测试：
    float32 size :  4
    float64 size :  8
    int size :  8
    uint size :  8
    IntToBytes : 0x1f123400ff123400 to 003412ff0034121f
    Int16ToBytes : 0x1f12 to 121f
    Int32ToBytes : 0x1f123400 to 0034121f
    Int64ToBytes : 0x1f123400ff123400 to 003412ff0034121f
    UintToBytes : 0xff123400ff123400 to 003412ff003412ff
    Uint16ToBytes : 0xff12 to 12ff
    Uint32ToBytes : 0xff123400 to 003412ff
    Uint64ToBytes : 0xff123400ff123400 to 003412ff003412ff
    Float32ToBytes : -1234.567749 to 2b529ac4
    Float64ToBytes : -1234.567800 to adfa5c6d454a93c0
    BytesToInt(8Bytes) : 030405060708090a to 0xa09080706050403
    BytesToInt(>8Bytes) : 0102030405060708090a to 0x807060504030201
    BytesToInt(2Bytes) : 0102 to 0x201
```