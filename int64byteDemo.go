package main
import (
	"fmt"
	"encoding/binary"
)

func main() {
	var i int64 = 2323
	buf := Int64ToBytes(i)
	fmt.Println(buf)
	fmt.Println(BytesToInt64(buf))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}