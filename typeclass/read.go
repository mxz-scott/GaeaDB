package typeclass

// read失败时返回错误信息，出错时字节数组返回[]byte{}
type Read interface {
	Read([]byte) ([]byte, error) // 返回尚未解析的字节数组
}
