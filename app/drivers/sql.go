package drivers

import (
	"database/sql"
	"fmt"
	"github.com/ivanzhangio/ruth-fishtank-api/conf"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

var DB *sql.DB

func InitDB() {
	connstring := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DB_USER, conf.DB_PASS, conf.DB_NAME)

	var err error
	DB, err = sql.Open("postgres", connstring)
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")

}
