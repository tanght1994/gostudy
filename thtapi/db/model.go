package db

// BaseTime 创建时间, 修改时间
type BaseTime struct {
	CreatedAt int64 `gorm:"column:created_at;type:bigint;index" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;type:bigint;index" json:"updated_at"`
}

// UserPassword 用户密码
type UserPassword struct {
	BaseTime
	Username string `gorm:"column:username;primary_key;type:char(100)" json:"username"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
}

// TableName 表名
func (UserPassword) TableName() string {
	return "user_password"
}

// UserInfo 用户信息
type UserInfo struct {
	BaseTime
	Username string `gorm:"column:username;primary_key;type:char(255)" json:"username"`
	Nickname string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
}

// TableName 表名
func (UserInfo) TableName() string {
	return "user_info"
}

// GroupURL 组可访问的URL
type GroupURL struct {
	BaseTime
	GroupName string `gorm:"column:group_name;primary_key;type:char(100)" json:"group_name"`
	URL       string `gorm:"column:url;primary_key;type:char(500)" json:"url"`
}

// TableName 表名
func (GroupURL) TableName() string {
	return "group_url"
}

// UserGroup 用户拥有的组
type UserGroup struct {
	BaseTime
	Username  string `gorm:"column:username;primary_key;type:char(100)" json:"username"`
	GroupName string `gorm:"column:group_name;primary_key;type:char(100)" json:"group_name"`
}

// TableName 表名
func (UserGroup) TableName() string {
	return "user_group"
}

// URLWhiteList 免权限URL
type URLWhiteList struct {
	BaseTime
	URL string `gorm:"column:url;primary_key;type:char(500)" json:"url"`
}

// TableName 表名
func (URLWhiteList) TableName() string {
	return "url_whitelist"
}

// ServerAddr 免权限URL
type ServerAddr struct {
	BaseTime
	ServerName string `gorm:"column:server_name;primary_key;type:char(100)" json:"server_name"`
	ServerAddr string `gorm:"column:server_addr;primary_key;type:char(100)" json:"server_addr"`
	Status     string `gorm:"column:status;type:char(100)" json:"status"`
}

// TableName 表名
func (ServerAddr) TableName() string {
	return "server_addr"
}

// URL2URL URL对应的URL
type URL2URL struct {
	BaseTime
	OriginURL string `gorm:"column:origin_url;primary_key;type:char(500)" json:"origin_url"`
	TargetURL string `gorm:"column:target_url;primary_key;type:char(500)" json:"target_url"`
}

// TableName 表名
func (URL2URL) TableName() string {
	return "url2url"
}

// TableModified 用于记录表的修改
// 每次表中有数据被修改, 就将Tag+1
type TableModified struct {
	Name string `gorm:"column:name;primary_key;type:char(100)" json:"name"`
	Tag  int64  `gorm:"column:tag;type:bigint" json:"tag"`
}

// TableName 表名
func (TableModified) TableName() string {
	return "table_modified"
}
