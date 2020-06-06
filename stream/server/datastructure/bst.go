package datastructure

type BST struct {
	Data  int32
	Left  *BST
	Right *BST
}

func Insert(bst *BST, data int32) *BST {
	if bst == nil {
		return &BST{Data: data}
	}
	if data < bst.Data {
		bst.Left = Insert(bst.Left, data)
		return bst
	}
	bst.Right = Insert(bst.Right, data)
	return bst
}

func ToArray(bst *BST) []int32 {
	if bst == nil {
		return nil
	}

	arr := ToArray(bst.Left)
	arr = append(arr, bst.Data)
	arr = append(arr, ToArray(bst.Right)...)
	return arr
}
