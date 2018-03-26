package main 

import (
    "bytes"
    "encoding/binary"
    "log"
)

func IntToHex(num int64) []byte{
    
    //IntToHex converts int to a byte array
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, num)
    if err != nil {
        log.Panic(err)
    }
    return buff.Bytes()
}
