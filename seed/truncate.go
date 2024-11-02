package seed

import (
	"fmt"

	"gorm.io/gorm"
)

type SeedHandler struct {
	Db *gorm.DB
}

func (sh *SeedHandler) TruncateTable(tables []string) {
	for _, v := range tables {
		_ = sh.Db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", v))
	}
}
