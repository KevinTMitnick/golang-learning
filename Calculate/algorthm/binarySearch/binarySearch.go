package binarySearch

type Item struct {
	left int
	right int
	value int
}

var config = []*Item{
	{0,10,1},
	{10, 20, 2},
	{20, 100, 8},
	{100, 1<<31, 10},
}

func GetPrice(price int) int  {
	start, end := 0, len(config)
	mid := (start + end) / 2

	for mid >= 0 && mid < len(config){
		if price < config[mid].left {
			end = mid -1
		} else if price >= config[mid].right{
			start = mid + 1
		} else {
			return config[mid].value
		}
		mid = (start + end) / 2
	}
	return -1
}
