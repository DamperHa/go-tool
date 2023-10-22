package basic

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

// 尽量定义零值可用的类型
// C语言竟然没有对变量自动赋0值，而是任意值；

// C语言中，我们首先需要对变量进行初始化操作。

type FinishDecorationMessage struct {
	StuId          int `json:"stuId"`
	BizId          int `json:"bizId"`
	CourseId       int `json:"courseId"`
	BatchId        int `json:"batchId"`        // 批次id
	UnitId         int `json:"unitId"`         // 单元id
	PlanId         int `json:"planId"`         // 场次id
	DramaId        int `json:"dramaId"`        // 剧本id
	StuCouId       int `json:"stuCouId"`       // 购课id
	ChapterId      int `json:"chapterId"`      // 章id
	ChapterLogicId int `json:"chapterLogicId"` // 章逻辑id

	SectionId      string `json:"sectionId"`
	SectionLogicId string `json:"sectionLogicId"`

	UnlockType int `json:"unlockType"` // 解锁类型，1：解锁章，2：解锁场次

	MetaData map[string]string `json:"metadata"`
}

func TestDemo1(t *testing.T) {
	arg := FinishDecorationMessage{}
	r, _ := jsoniter.MarshalToString(arg)
	fmt.Println(r)
}

func TestSlice(t *testing.T) {
	s := make([]int, 2)
	s = append(s, 1)
	fmt.Println(s)
}
