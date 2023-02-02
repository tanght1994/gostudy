package db

// BaseTime 创建时间, 修改时间
type BaseTime struct {
	CreatedAt int64 `gorm:"column:created_at;type:bigint;index" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;type:bigint;index" json:"updated_at"`
}

// modelUserPassword 用户密码
type modelUserPassword struct {
	Username string `gorm:"column:username;primary_key;type:char(100)" json:"username"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
	BaseTime
}

// TableName 表名
func (modelUserPassword) TableName() string {
	return "user_password"
}

// modelUserInfo 用户信息
type modelUserInfo struct {
	Username string `gorm:"column:username;primary_key;type:char(255)" json:"username"`
	Nickname string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
	BaseTime
}

// TableName 表名
func (modelUserInfo) TableName() string {
	return "user_info"
}

// modelUserGroup 用户拥有的组
type modelUserGroup struct {
	Username string `gorm:"column:username;primary_key;type:char(100)" json:"username"`
	GroupID  uint32 `gorm:"column:group_id;primary_key;type:int unsigned" json:"group_id"`
	BaseTime
}

// TableName 表名
func (modelUserGroup) TableName() string {
	return "user_group"
}

// modelGroupInfo 组ID与组名对应关系
type modelGroupInfo struct {
	GroupID   uint32 `gorm:"column:group_id;primary_key;type:int unsigned;auto_increment;comment:'组ID'" json:"group_id"`
	GroupName string `gorm:"column:group_name;type:char(100);comment:'组名'" json:"group_name"`
	BaseTime
}

// modelProxyConfig 代理规则配置
type modelProxyConfig struct {
	ID   uint64 `gorm:"column:id;primary_key;type:bigint unsigned;auto_increment;comment:'充当主键,没其它用处'" json:"id"`
	Text string `gorm:"column:text;type:MEDIUMTEXT;comment:'配置文件'" json:"text"`
	BaseTime
}

// TableName 表名
func (modelProxyConfig) TableName() string {
	return "proxy_config"
}
