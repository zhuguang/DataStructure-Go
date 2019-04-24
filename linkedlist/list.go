package linkedlist


type dictEntry struct {
	key  *interface{}
	prev *dictEntry
	next *dictEntry
}

func dictEntryCreate() *dictEntry {
	dict := &dictEntry{key: nil, next: nil}
	return dict
}

func (d *dictEntry) dictEntryPush(key interface{}) {
	newDict := dictEntryCreate()
	newDict.key = &key
	newDict.next = d.next
	d.next = newDict
}
