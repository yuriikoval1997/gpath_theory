package data_structures

type ArrayList struct {
	storage []string
	head int
}

func (arr *ArrayList) Add(str string) {
	if len(arr.storage) >= arr.head {
		var tmp = make([]string, int(float32(len(arr.storage)) * 1.61))
		copy(tmp, arr.storage)
	}
	arr.head++
	arr.storage[arr.head] = str
}

func (arr *ArrayList) Remove() (string, bool) {
	if arr.head == 0 {
		return "", false
	}
	if float64(len(arr.storage))/float64(arr.head) > 1.61 {
		var tmp = make([]string, int(float64(len(arr.storage))/1.61)+1)
		copy(tmp, arr.storage)
		arr.storage = tmp
	}
	var res = arr.storage[arr.head-1]
	arr.head--
	return res, true
}

func (arr *ArrayList) IsEmpty() bool {
	if arr.head == 0 {
		return true
	}
	return false
}
