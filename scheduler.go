// 在庫情報を監視するタスクスケジューラ
// 現在はStockテーブルのEnd時刻と同期している
package main

import (
	"Go_TaskSheduler/db"
	"context"
	"fmt"
	"time"
)

// デフォルトの遅延時間(s)
const d int = 20

func scheduler() error {
	var cnt int
	delay := time.Duration(d) * time.Second

	fmt.Println("Scheduler started. delay time :", delay)

	// データベース接続用クライアントの作成
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return fmt.Errorf("クライアント接続エラー : %w", err)
	}
	defer func() {
		// クライアントの切断
		if err := client.Prisma.Disconnect(); err != nil {
			panic(fmt.Errorf("クライアント切断エラー : %w", err))
		}
	}()

	ctx := context.Background()

	for {
		cnt++

		// Stockテーブルで、現在時刻よりも後の情報を取得
		stock, err := client.Stock.FindMany(
			db.Stock.End.After(GetTime()),
		).Exec(ctx)
		if err != nil {
			return fmt.Errorf("在庫テーブル取得エラー : %w", err)
		}

		// 現在より後の情報がない場合
		if len(stock) == 0 {
			fmt.Printf("[cnt.%d] 更新予定なし\n", cnt)
			time.Sleep(delay)
			continue
		}

		// 比較用にStockテーブルの1行をセット
		stock_one := stock[0]

		for _, value := range stock {
			// fmt.Printf("index : %d, value : %+v\n", index, value)

			// 終了時刻がより早い場合はその行を新たにセット
			if value.End.Before(stock_one.End) {
				stock_one = value
			}
		}

		// 現在時刻との間隔を求める
		duration := stock_one.End.Sub(GetTime())

		// durationがdelayよりも短い場合
		// その間隔分遅延し、遅延後にタスク処理を実行
		if duration < delay {
			fmt.Printf("[cnt.%d] 更新実行: %v後...", cnt, duration)
			time.Sleep(duration)

			if err := task(client, stock_one); err != nil {
				fmt.Println("エラー")
				return fmt.Errorf("タスク実行エラー : %w", err)
			} else {
				fmt.Println("完了")
			}

		} else {
			fmt.Printf("[cnt.%d] 次回更新: %v (%v後)\n", cnt, stock_one.End, duration)
			time.Sleep(delay)
		}
	}
}
