package drivers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ranomics/ranomics-web/conf"
	"github.com/revel/revel"
)

var DB *sql.DB

func InitDB() {
	connstring := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DbUser(), conf.DbPass(), conf.DbName())

	var err error
	DB, err = sql.Open("postgres", connstring)
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")
	revel.INFO.Println("Loading genes from DB...")

}
