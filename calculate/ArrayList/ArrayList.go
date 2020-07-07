package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int 		//数组大小
	Get(index int) (interface{}, error)	// 抓取第几个元素
	Set(index int, newval interface{}) error	// 修改数据
	Insert(index int, newval interface{}) error	//插入
	Append(newval interface{}) 	// 追加
	Clear()
	Delete(index int) error
	String() string	//返回字符串
	Iterator() Iterator
}

// 数据结构: 字符串、整数、实数
type ArrayList struct {
	dataStore [] interface{}	// 数组存储
	TheSize int		// 数组大小
}

func NewArryList() *ArrayList  {
	list := new(ArrayList)
	list.dataStore = make([]interface{},0, 10)
	list.TheSize = 0
	return list
}

func (list *ArrayList) checkisFull() {
	if list.TheSize == cap(list.dataStore) { //判断内存使用

		// make  中间的参数， cap = 0， 没有开辟内存，
		newdataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize) //开辟双倍内存

		// copy(newdataStore, list.dataStore)	//内存越界的时候会有问题:

		//for i := 0; i < len(list.dataStore); i++ {
		//	newdataStore[i] = list.dataStore[i]
		//}

		list.dataStore = newdataStore		// 重新赋值
		//fmt.Println(list.dataStore)
		//fmt.Println(newdataStore)
	}
}

func (list *ArrayList) Size() int  {
	return list.TheSize
}

func (list *ArrayList)  Get(index int) (interface{}, error)	{
	if index  < 0  || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}


func (list *ArrayList)  Set(index int, newval interface{}) error {
	if index  < 0  || index >= list.TheSize {
		return errors.New("索引越界")
	}

	list.dataStore[index] = newval 	// 设置
	return nil
}

func (list *ArrayList)  Insert(index int, newval interface{}) error	{
	if index  < 0  || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkisFull()	// 检测内存，如果满了，自动追加

	list.dataStore = list.dataStore[:list.TheSize+1]	//插入数据的时候，需要把内存移动一位
	for i := list.TheSize; i>index;i--{
		list.dataStore[i] = list.dataStore[i - 1]
	}
	list.dataStore[index] = newval
	list.TheSize ++		// 索引追加
	return nil
}

func (list *ArrayList) Append(newval interface{}) 	{
	list.dataStore = append(list.dataStore, newval)
	list.TheSize ++
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)		//重新开辟内存
	list.TheSize = 0
}

func (list *ArrayList)  Delete(index int) error {		//删除
	list.dataStore = append(list.dataStore[:index],list.dataStore[index+1:]...)		//重新叠加
	list.TheSize--
	return nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}