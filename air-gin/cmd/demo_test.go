package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	str := "CREATE TABLE `xes_video_note_%d` ( `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id', `plan_id` int(11) NOT NULL DEFAULT '0' COMMENT '场次id', `stu_id` int(11) NOT NULL DEFAULT '0' COMMENT '学生id', `video_timestamp` int(11) NOT NULL COMMENT '标记打点时间戳', `img_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片url', `offset` int(11) NOT NULL DEFAULT '0' COMMENT '视频偏移量', `role` int(11) NOT NULL DEFAULT '0' COMMENT '角色，1：主讲，2：辅导，3：学生', `is_use_offset` int(11) NOT NULL DEFAULT '0' COMMENT '1：录直播，其他：0', `content` varchar(500) NOT NULL COMMENT '自定义文本', `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间', `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除标记，0：未删除，1：删除', `biz_id` int(11) NOT NULL DEFAULT '0' COMMENT '场次类型', `business_line_id` tinyint(4) NOT NULL COMMENT '业务线(10.学而思网校,素质app 20.学而思网课,高中app 30.学小思，学科app)', PRIMARY KEY (`id`), KEY `plan_stu` USING BTREE (`stu_id`, `plan_id`) COMMENT '场次与学生id建索引' ) ENGINE = InnoDB CHARSET = utf8mb4 COMMENT '视频打点数据——范志豪';"
	for i := 0; i <= 15; i++ {
		fmt.Printf(str, i)
		fmt.Println()
	}
}
