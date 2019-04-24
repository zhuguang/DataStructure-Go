package linkedlist

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	dictDemo := dictEntryCreate()
	dictDemo.dictEntryPush("one")
	dictDemo.dictEntryPush(2)
	dictDemo.dictEntryPush("three")
	for dictDemo.next != nil {
		if dictDemo.key != nil {
			fmt.Printf("dictDemo:%v,", *dictDemo.key)
		}
		dictDemo = dictDemo.next
	}
}
