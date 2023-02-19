package ydaos

import "time"

type UserMessage struct {
	Id             int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Uid            int64     `xorm:"not null INT(11) index"`
	FromUid        int64     `xorm:"not null BIGINT(20)"`    // 发送者uid，系统邮件为0
	SysKey         string    `xorm:"not null VARCHAR(1024)"` // 系统邮件的唯一标识
	TplId          int64     `xorm:"not null BIGINT(20)"`
	Type           int       `xorm:"not null INT(11) DEFAULT 0"`        // 消息类型
	Title          string    `xorm:"not null text"`                     // 消息标题
	ContentData    string    `xorm:"not null text"`                     // 当TplId非0时，为模版参数json，TplId为0时，为消息体本身
	AttachmentData string    `xorm:"not null VARCHAR(1024) DEFAULT ''"` // 附件内容描述
	Args           string    `xorm:"not null VARCHAR(1024) DEFAULT ''"` // 用户客户端的一些附加的参数，比如跳转
	LogTag         int       `xorm:"not null INT(11) DEFAULT 0"`
	Status         int       `xorm:"not null INT(11) DEFAULT 1"` // 消息状态，值参考user_svr/app/common/consts的MessageStatusEnum
	DayReset       time.Time `xorm:"not null DATETIME"`
	Test3          int       `xorm:"not null BIGINT(20) DEFAULT 0"`
	/*
		sync
		1.可以识别到新增字段但是无法识别到字段内容的改动
		2.当改动字段名(同时可以更改字段属性)的时候旧的字段不会删除，然后会另外新建一个字段（新的属性也会生效）
	*/

	Created time.Time `xorm:"not null DATETIME created index"`
	Updated time.Time `xorm:"not null DATETIME updated"`
	Deleted time.Time `xorm:"deleted index"`
}

type User struct {
	Id   int64  `xorm:"not null pk autoincr BIGINT(20)"`
	Uid  int64  `xorm:"not null INT(11) index"`
	Name string `xorm:"not null text"`
}
