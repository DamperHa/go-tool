package basic

import (
	"fmt"
	"testing"
)

// 领用golang写一个堆的代码
type IntSlice []int

func (r *IntSlice) Len() int {
	return len(*r)
}

func (r *IntSlice) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func (r *IntSlice) Less(i, j int) bool {
	return (*r)[i] < (*r)[j]
}

func (r *IntSlice) Push(x interface{}) {
	*r = append(*r, x.(int))
}

// TODO 直接对原数组操作有什么问题呢？
func (r *IntSlice) Pop() interface{} {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	// heap.Init()

	str := "CREATE TABLE `xes_video_note_%d` (\n  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `plan_id` int(11) NOT NULL DEFAULT '0' COMMENT '场次id',\n  `stu_id` int(11) NOT NULL DEFAULT '0' COMMENT '学生id',\n  `video_timestamp` int(11) NOT NULL DEFAULT '0' COMMENT '标记打点时间戳',\n  `img_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片url',\n  `offset` int(11) NOT NULL DEFAULT '0' COMMENT '视频偏移量',\n  `role` int(11) NOT NULL DEFAULT '0' COMMENT '角色，1：主讲，2：辅导，3：学生',\n  `is_use_offset` int(11) NOT NULL DEFAULT '0' COMMENT '1：录直播，其他：0',\n  `content` varchar(500) NOT NULL DEFAULT '' COMMENT '自定义文本',\n  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',\n  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除标记，0：未删除，1：删除',\n  `biz_id` int(11) NOT NULL DEFAULT '0' COMMENT '场次类型',\n  `business_line_id` tinyint(4) NOT NULL DEFAULT '10' COMMENT '业务线(10.学而思网校,素质app 20.学而思网课,高中app 30.学小思，学科app)',\n  PRIMARY KEY (`id`),\n  KEY `plan_stu` (`stu_id`,`plan_id`) USING BTREE COMMENT '场次与学生id建索引'\n) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='视频打点数据——范志豪';"
	for i := 0; i < 16; i++ {
		res := fmt.Sprintf(str, i)
		fmt.Println(res)
	}

}

func TestSliceCopy(t *testing.T) {
	originSlice := []int{1, 2, 3, 4}

	newSlice := originSlice[1:3]

	newSlice[0] = 100

	fmt.Println(originSlice, newSlice)
}

func TestSliceCopyPointer(t *testing.T) {
	originSlice := []int{1, 2, 3, 4}

	newSlice := originSlice

	newSlice[0] = 100

	fmt.Println(originSlice, newSlice)
}

func TestDelSlice(t *testing.T) {
	originSlice := make([]int, 100, 100)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				old := originSlice
				n := len(old)
				if n > 0 {
					old = old[:n-1]
					originSlice = old
				}

				fmt.Println(originSlice)
			}

		}()
	}

	select {}
}

func TestDelSliceV2(t *testing.T) {
	originSlice := make([]int, 100, 100)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				n := len(originSlice)
				if n > 0 {
					originSlice = originSlice[:n-1]
				}

				fmt.Println(originSlice)
			}

		}()
	}

	select {}
}
