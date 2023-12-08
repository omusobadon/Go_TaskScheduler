// スケジューラで実行される処理
package main

import (
	"Go_TaskSheduler/db"
	"context"
	"fmt"
)

func task(c *db.PrismaClient, s db.StockModel) error {
	ctx := context.Background()

	// Delete
	_, err := c.Stock.FindUnique(
		db.Stock.ID.Equals(s.ID),
	).Delete().Exec(ctx)
	if err != nil {
		return fmt.Errorf("在庫テーブル削除エラー : %w", err)
	}

	return nil
}
