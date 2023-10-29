package entity

import "time"

type SysConfig struct {
	ID             string    `gorm:"column:config_id;primaryKey;not null;size:18"`
	ConfigName     string    `gorm:"column:config_name;not null;size:45"`
	ConfigGroup    string    `gorm:"column:config_group;not null;size:45"`
	ConfigValue    string    `gorm:"column:config_value;not null;type:text"`
	ConfigDesc     *string   `gorm:"column:config_desc;size:255"`
	ConfigPortalID *string   `gorm:"column:config_portal_id;default:null;size:45"`
	Restricted     int8      `gorm:"column:restricted;not null;size:1"`
	ModifiedBy     string    `gorm:"column:mdb;not null;size:18"`
	ModifiedDate   time.Time `gorm:"column:mdd;not null"`
}

type SysUser struct {
	UserID       string     `gorm:"column:user_id;primaryKey;size:18;not null"`
	Username     string     `gorm:"column:username;not null;size:45"`
	KataSandi    string     `gorm:"column:kata_sandi;not null;size:200"`
	NamaLengkap  string     `gorm:"column:nama_lengkap;not null;size:100"`
	Telepon      *string    `gorm:"column:telepon;size:15"`
	Email        *string    `gorm:"column:email;size:255"`
	JenisKelamin *string    `gorm:"column:jenis_kelamin;size:1"`
	Foto         *string    `gorm:"column:foto;size:100"`
	Status       *int       `gorm:"column:status"`
	LastLogin    *time.Time `gorm:"column:last_login;not null"`
	RegisteredBy *string    `gorm:"column:registered_by;size:45"`
	ModifiedDate *time.Time `gorm:"column:mdd;not null"`
}

type SysAuthLog struct {
	LogID      string    `gorm:"column:log_id;primaryKey;size:32;not null"`
	User       string    `gorm:"column:user;not null;size:100"`
	WaktuLogin time.Time `gorm:"column:waktu_login;not null"`
	IP         string    `gorm:"column:ip;not null;size:18"`
	UserAgent  string    `gorm:"column:user_agent;not null;size:150"`
	Keterangan *string   `gorm:"column:keterangan;size:255"`
	Status     string    `gorm:"column:status;not null;size:10"`
}

type SysGroup struct {
	GroupID         string    `gorm:"column:group_id;primaryKey;size:18;not null"`
	GroupName       string    `gorm:"column:group_name;not null;size:45"`
	GroupDesc       string    `gorm:"column:group_desc;not null;size:100"`
	GroupPortal     string    `gorm:"column:group_portal;not null;size:45"`
	GroupRestricted int       `gorm:"column:group_restricted;not null"`
	ModifiedBy      int       `gorm:"column:mdb;not null"`
	ModifiedDate    time.Time `gorm:"column:mdd;not null"`
}

type SysLoginAttempts struct {
	LoginID   string     `gorm:"column:login_id;primaryKey;size:18;not null"`
	IPAddress *string    `gorm:"column:ip_address;size:50"`
	Login     *string    `gorm:"column:login;size:50"`
	Time      *time.Time `gorm:"column:time"`
}

type SysMenu struct {
	MenuID       *uint     `gorm:"column:menu_id;primaryKey;autoIncrement"`
	PortalID     uint      `gorm:"column:portal_id;not null"`
	MenuName     string    `gorm:"column:menu_name;not null;size:45"`
	MenuDesc     string    `gorm:"column:menu_desc;not null;size:45"`
	MenuPosition string    `gorm:"column:menu_position;not null;size:20"`
	MenuOrder    int       `gorm:"column:menu_order;not null;type:int(3)"`
	MenuParent   int       `gorm:"column:menu_parent;not null"`
	MenuLink     string    `gorm:"column:menu_link;not null;size:45"`
	MenuShow     *int8     `gorm:"column:menu_show;type:int(1)"`
	MenuIcon     string    `gorm:"column:menu_icon;not null;size:45"`
	MenuFontIcon string    `gorm:"column:menu_fonticon;not null;size:45"`
	ModifiedBy   string    `gorm:"column:mdb;not null;size:18"`
	ModifiedDate time.Time `gorm:"column:mdd;not null"`

	// Define foreign key constraint in GORM model
	SysPortal SysPortal `gorm:"foreignKey:PortalID;references:PortalID;onDelete:CASCADE"`
}

type SysPermission struct {
	GroupID    string `gorm:"column:group_id;not null;size:18"`
	MenuID     uint   `gorm:"column:menu_id;not null"`
	Permission string `gorm:"column:permission;not null;size:4"`

	// Define primary key constraint in GORM model
	PrimaryKey struct{} `gorm:"primaryKey:group_id,menu_id"`
	SysGroup   SysGroup `gorm:"foreignKey:GroupID;references:GroupID;onDelete:CASCADE"`
	SysMenu    SysMenu  `gorm:"foreignKey:MenuID;references:MenuID;onUpdate:CASCADE;onDelete:CASCADE"`
}

type SysPortal struct {
	PortalID     uint      `gorm:"column:portal_id;primaryKey;autoIncrement"`
	PortalNumber string    `gorm:"column:portal_number;not null;size:2;unique"`
	PortalName   string    `gorm:"column:portal_name;not null;size:45"`
	PortalDesc   string    `gorm:"column:portal_desc;not null;size:255"`
	PortalLink   *string   `gorm:"column:portal_link;size:100"`
	MetaTitle    string    `gorm:"column:meta_title;not null;size:150"`
	MetaDesc     string    `gorm:"column:meta_desc;not null;size:150"`
	MetaTag      string    `gorm:"column:meta_tag;not null;size:150"`
	ModifiedBy   int       `gorm:"column:mdb;not null"`
	ModifiedDate time.Time `gorm:"column:mdd;not null"`
}

type SysUserGroup struct {
	UserID  string `gorm:"column:user_id;not null;size:18"`
	GroupID string `gorm:"column:group_id;not null;size:18"`
	Default int    `gorm:"column:default;not null;size:1"`

	// Define primary key constraint in GORM model
	PrimaryKey struct{} `gorm:"primaryKey:user_id,group_id"`

	// Define foreign key constraints in GORM model
	SysUser  SysUser  `gorm:"foreignKey:UserID;references:UserID;onUpdate:CASCADE;onDelete:CASCADE"`
	SysGroup SysGroup `gorm:"foreignKey:GroupID;references:GroupID"`
}
