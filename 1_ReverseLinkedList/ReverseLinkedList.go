package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := CreateTestData("[1,2,3,4]")
	x := reverseList(list)
	Print(x)
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func CreateTestData(data string) *ListNode {
	if data == "[]" {
		return nil
	}
	data = string([]rune(data)[1 : len(data)-1])
	res := strings.Split(data, ",")
	length := len(res)
	listNode := make([]ListNode, length)
	headVal, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}
	listNode[0] = ListNode{headVal, nil}
	for i := 1; i < length; i++ {
		headVal, _ = strconv.Atoi(res[i])
		listNode[i] = ListNode{headVal, nil}
		listNode[i-1].Next = &listNode[i]
	}
	return &listNode[0]
}

func Print(listNode *ListNode) {
	if listNode == nil {
		fmt.Println(nil)
	}
	var buffer strings.Builder
	buffer.WriteString("[")
	value := strconv.Itoa(listNode.Val)
	buffer.WriteString(value)
	temp := listNode.Next
	for temp != nil {
		buffer.WriteString(",")
		value = strconv.Itoa(temp.Val)
		buffer.WriteString(value)
		temp = temp.Next
	}
	buffer.WriteString("]")
	fmt.Println(buffer.String())
}
