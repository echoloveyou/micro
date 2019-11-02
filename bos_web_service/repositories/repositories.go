package repositories

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
)

var (
	dbs       = map[string]*Repository{}
	Confs     = map[string]MysqlConf{}
	OfficeBos = beego.AppConfig.String("db_database")
	messageFormate = "======启动======%s\n"
)

type Repository struct {
	*gorm.DB
}

type MysqlConf struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func NewRepository(conf MysqlConf) *Repository {
	var (
		err error
		repository Repository
	)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		conf.user, conf.password, conf.host, conf.port, conf.database, url.QueryEscape(beego.AppConfig.String("time_loc")))
	fmt.Println(dbUrl)
	fmt.Printf(messageFormate, "mysql 开始连接")
	repository.DB, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		fmt.Printf(messageFormate, "mysql 连接失败")
		panic(err)
	}
	fmt.Printf(messageFormate, "mysql 连接成功")
	return &repository
}

func NewMysqlConf(host, port, user, password, database string) MysqlConf {
	return MysqlConf{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
	}
}

func (repo *Repository) GDB(databases ...string) *Repository {
	var database string
	if len(databases) == 0 {
		database = OfficeBos
	} else {
		database = databases[0]
	}

	if repo.DB == nil {
		if dbs[database].DB != nil {
			repo := NewRepository(Confs[database])
			dbs[database] = repo
		} else {
			repo = dbs[database]
		}
	}
	return repo
}

func loadConf() map[string]MysqlConf {
	Confs[beego.AppConfig.String("bos_db_database")] = NewMysqlConf(
		beego.AppConfig.String("bos_db_host"),
		beego.AppConfig.String("bos_db_port"),
		beego.AppConfig.String("bos_db_user"),
		beego.AppConfig.String("bos_db_password"),
		beego.AppConfig.String("bos_db_database"),
	)

	//Confs[beego.AppConfig.String("a")] = NewMysqlConf(
	//	beego.AppConfig.String("a"),
	//	beego.AppConfig.String("a"),
	//	beego.AppConfig.String("a"),
	//	beego.AppConfig.String("a"),
	//	beego.AppConfig.String("a"),
	//)
	return Confs
}

//记得服务关闭时调用close方法关闭mysql连接
func Run()  {
	for _, v := range loadConf() {
		dbs[v.database] = NewRepository(v)
	}
}

func Close() {
	for _, v := range dbs {
		_ = v.Close()

	}
}