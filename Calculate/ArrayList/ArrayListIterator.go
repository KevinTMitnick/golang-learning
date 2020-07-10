package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool 	//是否有下一个
	Next() (interface{}, error)
	Remove()	//删除
	GetIndex() int 	//得到索引
}

type Iterable interface {
	Iterator() Iterator
}

// 构造指针访问数组
type ArrayListIterator struct {
	list * ArrayList
	currenindex int
}

func (list * ArrayList) Iterator() *ArrayListIterator {
	it := new(ArrayListIterator)
	it.currenindex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool  {
	return it.currenindex < it.list.TheSize //是否有下一个
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("Not Next")
	}
	value, err := it.list.Get(it.currenindex)	//抓取当前数据
	it.currenindex++
	return value,err
}

func (it *ArrayListIterator) Remove()  {
	it.currenindex--
	it.list.Delete(it.currenindex)
}

func (it *ArrayListIterator)	 GetIndex() int  {
	return it.currenindex
}