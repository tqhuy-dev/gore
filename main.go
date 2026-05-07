package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tqhuy-dev/gore/utilities"
)

type Orders struct {
	ID        int
	OrderCode string
	Amount    float64
	UserId    int
	Status    int8
	CreatedAt time.Time
}

func main() {
	// Thông tin kết nối
	dsn := "admin:admin@tcp(127.0.0.1:3306)/data"

	// Mở kết nối
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Lỗi khi mở kết nối: %v", err)
	}
	defer db.Close()

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		log.Fatalf("Không thể kết nối tới MySQL: %v", err)
	}
	arr := make([]Orders, 0, 10000)
	fmt.Println("✅ Kết nối MySQL thành công!")
	var placeholders []string
	var values []interface{}
	var counting int
	for i := 0; i < 30000000; i++ {
		arr = append(arr, Orders{
			ID:        i,
			OrderCode: utilities.RandomString(15, utilities.AlphanumericCharset),
			Amount:    10000,
			UserId:    1,
			Status:    1,
			CreatedAt: time.Now().AddDate(-1, 0, 0).Add(time.Duration(i) * time.Second),
		})
		counting += 1
		if counting == 10000 {
			for _, u := range arr {
				placeholders = append(placeholders, "(?, ?, ?, ?, ?)")
				values = append(values, u.CreatedAt, u.OrderCode, u.Amount, u.UserId, u.Status)
			}
			query := fmt.Sprintf("INSERT INTO orders (created_at, order_code, amount, user_id, status) VALUES %s",
				strings.Join(placeholders, ","))
			result, err := db.Exec(query, values...)
			if err != nil {
				panic(err)
			}
			rows, _ := result.RowsAffected()
			fmt.Printf("✅ Đã insert %d dòng\n", rows)

			placeholders = make([]string, 0)
			values = make([]interface{}, 0)
			arr = make([]Orders, 0, 10000)
			counting = 0
		}
	}

}
