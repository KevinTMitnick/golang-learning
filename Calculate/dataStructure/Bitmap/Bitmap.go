package Bitmap

import (
	"fmt"
	"strings"
)

const (
	bitSize = 8
)

var bitmask = []byte{1, 1 <<1, 1<< 2, 1<< 3, 1<< 4, 1<< 5, 1<<6, 1<<7}

// 首字母小写 只能调用 工厂函数 创建
type bitmap struct {
	bits []byte
	bitcount uint64	// 已填入的数量
	capsize uint64	 //容量
}

//
func NewBitmap(maxnum uint64) *bitmap {
	return &bitmap{make([]byte, (maxnum+7)/bitSize),0, maxnum}
}

// 索引
func (b *bitmap) offset(num uint64) (byteIndex uint64, bitPos byte)  {
	byteIndex = num / bitSize
	if byteIndex >= uint64(len(b.bits)){
		panic(fmt.Sprintf("Runtime error: Index value %d out of range", byteIndex))
		return
	}
	bitPos = byte(num % bitSize)
	return byteIndex, bitPos
}

// 填入数字
func (b *bitmap) Set(num uint64)  {
	byteIndex, bitPos := b.offset(num)
	b.bits[byteIndex] |= bitmask[bitPos]
}


// 清除填入数字
func (b *bitmap) Reset(num uint64)  {
	byteIndex, bitPos := b.offset(num)

	// 重置为空位
	b.bits[byteIndex] &= ^bitmask[bitPos]
	b.bitcount--
}

// 判断数字是否在 位图中
func (b *bitmap) Check(num uint64) bool {
	byteIndex := num / bitSize
	if byteIndex >= uint64(len(b.bits)) {
		return false
	}
	bitPos := num % bitSize
	// 右移 bitPos位 和 1进行按位与
	return !(b.bits[byteIndex]&bitmask[bitPos] == 0)
}


// bitmap 容量
func (b *bitmap) Size() uint64  {
	return uint64(len(b.bits) * bitSize)
}

// 位图是否为空
func (b *bitmap) IsEmpty() bool  {
	return b.bitcount == 0
}

// 是否填满
func (b *bitmap) IsFull() bool  {
	return b.bitcount == b.capsize
}

// 已填入的数字个数
func (b *bitmap) Count() uint64  {
	return b.bitcount
}


// 获取填入的数字切片
func (b *bitmap) GetData() []uint64  {
	var data []uint64
	count := b.Size()

	for index := uint64(0); index < count; index++{
		if b.Check(index){
			data = append(data, index)
		}
	}
	return data
}


func byteToBinaryString(data byte) string {
	var str strings.Builder
	for index := 0; index < bitSize; index++ {
		if (bitmask[7-index] & data) == 0 {
			str.WriteString("0")
		} else {
			str.WriteString("1")
		}
	}
	return str.String()
}

// 获取 位图中索引对应的字符串
func (b *bitmap) String() string  {
	var str strings.Builder

	for index := len(b.bits) -1 ; index >= 0 ;index--{
		str.WriteString(byteToBinaryString(b.bits[index]))
		str.WriteString(" ")
	}
	return str.String()
}



