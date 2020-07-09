package CricleQueue

const QueueSize = 100

type CricleQueue struct {
	data [QueueSize]interface{}		//存储数据的结构
	front int	//头部
	rear int	// 尾部
}

func initQueue(q *CricleQueue)  {	//初始化，头部尾部重合，为空
	q.front = 0
	q.rear = 0
}

func Queuelenght(q *CricleQueue) int  {	//队列长度
	return (q.rear - q.front + QueueSize) % QueueSize
}

func EnQueue(q *CricleQueue, data interface{})  (err error){

}

func DeQueue(q *CricleQueue)  (data interface{},err error){

}