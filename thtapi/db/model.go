package db

// BaseTime 创建时间, 修改时间
type BaseTime struct {
	CreatedAt int64 `gorm:"column:created_at;type:bigint;index" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;type:bigint;index" json:"updated_at"`
}

// modelUserPassword 用户密码
type modelUserPassword struct {
	BaseTime
	Username string `gorm:"column:username;primary_key;type:char(100)" json:"username"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
}

// TableName 表名
func (modelUserPassword) TableName() string {
	return "user_password"
}

// modelUserInfo 用户信息
type modelUserInfo struct {
	BaseTime
	Username string `gorm:"column:username;primary_key;type:char(255)" json:"username"`
	Nickname string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
}

// TableName 表名
func (modelUserInfo) TableName() string {
	return "user_info"
}

// modelGroupURL 组可访问的URL
type modelGroupURL struct {
	BaseTime
	GroupName string `gorm:"column:group_name;primary_key;type:char(100)" json:"group_name"`
	URL       string `gorm:"column:url;primary_key;type:char(500)" json:"url"`
}

// TableName 表名
func (modelGroupURL) TableName() string {
	return "group_url"
}

// modelUserGroup 用户拥有的组
type modelUserGroup struct {
	BaseTime
	Username  string `gorm:"column:username;primary_key;type:char(100)" json:"username"`
	GroupName string `gorm:"column:group_name;primary_key;type:char(100)" json:"group_name"`
}

// TableName 表名
func (modelUserGroup) TableName() string {
	return "user_group"
}

// modelURLWhiteList 免权限URL
type modelURLWhiteList struct {
	BaseTime
	URL string `gorm:"column:url;primary_key;type:char(500)" json:"url"`
}

// TableName 表名
func (modelURLWhiteList) TableName() string {
	return "url_whitelist"
}

// modelSvcName2SvcAddr svc_name 与 svc_addr 对应关系
type modelSvcName2SvcAddr struct {
	BaseTime
	SvcName string `gorm:"column:svc_name;primary_key;type:char(100)" json:"svc_name"`
	SvcAddr string `gorm:"column:svc_addr;primary_key;type:char(100)" json:"svc_addr"`
	Status  string `gorm:"column:status;type:char(50)" json:"status"`
}

// TableName 表名
func (modelSvcName2SvcAddr) TableName() string {
	return "svc_name2svc_addr"
}

// modelURL2EndPoint URL对应的EndPoint, EndPoint就是SvcName+URL, 如server1/a/b/c
type modelURL2EndPoint struct {
	BaseTime
	URL      string `gorm:"column:url;primary_key;type:char(350)" json:"url"`
	EndPoint string `gorm:"column:end_point;primary_key;type:char(350)" json:"end_point"`
}

// TableName 表名
func (modelURL2EndPoint) TableName() string {
	return "url2end_point"
}

// modelTableModified 用于记录表的修改
// 每次表中有数据被修改, 就将Tag+1
type modelTableModified struct {
	BaseTime
	Name string `gorm:"column:name;primary_key;type:char(100)" json:"name"`
	Tag  int64  `gorm:"column:tag;type:bigint" json:"tag"`
}

// TableName 表名
func (modelTableModified) TableName() string {
	return "table_modified"
}
