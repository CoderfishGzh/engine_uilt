package DB

import (
	_ "database/sql"
	"engine_utils/model"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// 初始化
func MysqlInit() {
	dsn := "root:122513gzhGZH!!@tcp(decs.pcl.ac.cn:1762)/search_engine?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	var err error
	DB, err = gorm.Open(
		mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		panic(err)
	}
}

// 通过docid到mysql中查找对应的关系model
func Get_docid_value(docids []string) (docs []model.Docs) {
	DB.Where("id IN (?)", docids).Find(&docs)
	return docs
}
