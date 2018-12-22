package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int    //主键
	UserName string //字段1
	PassWord string //字段2
}

type Article struct {
	Id       int       `orm:"pk;auto"`         //主键
	ArtiName string    `orm:"size(20)"`        //文章标题
	Atime    time.Time `orm:"auto_now"`        //文章创建时间
	Acount   int       `orm:"default(0);null"` //文章阅读量
	Acontent string    `orm:"size(500)"`       //文章内容
	Aimg     string    `orm:"size(100)"`       //文章存储路径
}

func init() {
	/*1.注册连接对象
	 * 第一个参数是数据库别名
	 * 第二个参数是数据库驱动
	 * 第三个参数是连接数据库字符串
	 */
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

	/*2.注册表
	 * 定义一个结构体 Person,注意字段首字母大写
	 * 方法中new结构体
	 */
	orm.RegisterModel(new(User), new(Article))

	/*3.建表
	 * 第一个参数对应注册的数据库别名
	 * 第二个参数是否强制更新，true会造成数据丢失
	 * 第三个参数是否可见sql语句
	 */
	orm.RunSyncdb("default", false, true)
}
