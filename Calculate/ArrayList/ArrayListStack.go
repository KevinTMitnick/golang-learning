package ArrayList

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}		//弹出
	Push(data interface{})	//压入
	IsFull() bool			//判断栈是否满了
	IsEmpty() bool			//判断栈是否为空
}


type Stack struct {
	myarray * ArrayList
	capsize int			// 最大范围
}

func NewArrayListStack() *Stack  {
	mystack := new(Stack)
	mystack.myarray = NewArryList()
	mystack.capsize = 10
	return mystack
}

func (mystack *Stack) Clear()  {
	mystack.myarray.Clear()
	mystack.capsize = 10
}

func (mystack *Stack) Size() int  {
	return  mystack.myarray.TheSize
}

func (mystack *Stack) Pop() interface{}  {
	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
		mystack.myarray.Delete(mystack.myarray.TheSize-1)
		return last
	}
	return nil
}

func (mystack *Stack) Push( data interface{})  {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}

func (mystack *Stack) IsFull() bool{
	if mystack.myarray.TheSize >= mystack.capsize {
		return true
	} else {
		return false
	}
}

func (mystack *Stack) IsEmpty() bool  {
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

