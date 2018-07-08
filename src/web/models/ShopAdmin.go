package models

type ShopAdmin struct {
	AdminId    uint   `xorm:"pk" xorm:"autoincr"`
	AdminUser  string `xorm:"notnull" xorm:"unique(shop_admin_adminuser_adminpass)" xorm:"unique(shop_admin_adminuser_adminemail)"`
	AdminPass  string `xorm:"notnull" xorm:"unique(shop_admin_adminuser_adminpass)"`
	AdminEmail string `xorm:"notnull" xorm:"unique(shop_admin_adminuser_adminemail)"`
	LoginTime  uint   `xorm:"notnull"`
	LoginIp    uint64 `xorm:"notnull"`
	CreateTime uint   `xorm:"notnull"`
}
