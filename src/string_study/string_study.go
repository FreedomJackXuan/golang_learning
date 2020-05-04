package string_study

/*
runtime/string.go
type stringStruct struct {
	str unsafe.Pointer
	len int
}
string 的pointer指向字节数组,默认为UTF-8编码，存储unicode字符， len返回字节数组的长度
*/



func String_Study() {
	var s string
	println(s == "")
	//println(s == nil) 字符串默认为""，而不是nil



}
