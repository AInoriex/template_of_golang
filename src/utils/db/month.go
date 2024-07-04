package db

import (
	"fmt"
	"time"
)

func GetTableNameAddMonth(pre string) string {
	y := time.Now().Year()  //年
	m := time.Now().Month() //月
	return fmt.Sprintf(pre+"%d%02d", y, m)
}
