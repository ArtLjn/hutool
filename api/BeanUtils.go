package api

type BeanUtils interface {
	CopyProperties(itemOne interface{}, itemTwo interface{})
	IsEmpty(interface{}, interface{}) bool
	ParseList(strList string, newList interface{})
}
