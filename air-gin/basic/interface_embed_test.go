package basic

import (
	"fmt"
	"testing"
)

type Writer interface {
	Write(str string)
}

type Reader interface {
	Read() (string, error)
}

type WriterImpl struct{}

func (r *WriterImpl) Write(str string) {
	fmt.Println(str)
}

type ReaderImpl struct{}

func (receiver *ReaderImpl) Read() (string, error) {
	fmt.Println("read")
	return "", nil
}

type ReaderWriter interface {
	Reader
	Writer
}

type ReaderWriterImpl struct {
	ReaderImpl
	WriterImpl
}

type ReaderWriterImplV2 struct {
	Reader
	Writer
}

func TestDemo(t *testing.T) {
	var rw ReaderWriter = &ReaderWriterImpl{}
	rw.Read()
	rw.Write("hello")
}

func TestDemoV2(t *testing.T) {
	var rw ReaderWriter = &ReaderWriterImplV2{
		Reader: &ReaderImpl{},
		Writer: &WriterImpl{},
	}

	rw.Read()
	rw.Write("hello")
}

// 考虑匿名嵌入
type SubV1 struct {
	Name int
}

func (receiver SubV1) Get() {
	fmt.Println("subv1 get")
}

type SubV2 struct {
	SubV1
	Name int
}

func (receiver SubV2) Get() {
	fmt.Println("subv2 get")
}

func TestEmbed(t *testing.T) {
	var sub SubV2
	sub.Get()
	sub.SubV1.Get()

	sub.Name = 1
	sub.SubV1.Name = 2
	fmt.Println(sub)
}

// 那么，我们在对象里面，嵌入一个接口呢
// 根据结果来，其实结构里面嵌入接口，和嵌入结构保持是一致的
type SubV3 struct {
	Writer
}

func (receiver SubV3) Write(str string) {
	fmt.Println("subV3", str)
}

func TestSubV3(t *testing.T) {
	sub := SubV3{}
	sub.Writer = &WriterImpl{}

	sub.Write("hello")
}
