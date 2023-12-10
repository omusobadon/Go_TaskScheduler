// スケジューラで実行される処理
package main

import (
	"Go_TaskSheduler/db"
	"context"
	"fmt"
)

func task(c *db.PrismaClient, s db.StockModel) error {
	ctx := context.Background()

	// 在庫数をデフォルトの個数にリセット
	// test, err := c.Stock.FindUnique(
	// 	db.Stock.ID.Equals(s.ID),
	// ).Exec(ctx)
	// if err != nil {
	// 	return fmt.Errorf("在庫テーブル削除エラー : %w", err)
	// }

	test, err := c.Product.FindMany(
		db.Product.ID.Equals(1),
	).Exec(ctx)
	if err != nil {
		return fmt.Errorf("テストエラー : %w", err)
	}

	fmt.Printf("test : %+v", test[0].Stock())

	return nil
}
