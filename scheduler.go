// 予約在庫情報をリセットするタスクスケジューラ
package main

import (
	"Go_TaskSheduler/db"
	"context"
	"fmt"
	"time"
)

// デフォルトの遅延時間(s)
const delay int = 60

func scheduler() error {
	var cnt int
	var flag bool
	var d time.Duration = delay * time.Second

	fmt.Println("Scheduler started")

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

		// スケジューラでのメイン処理
		if flag {
			// Delete
			_, err := client.Stock.FindUnique(
				db.Stock.ID.Equals(*s.ID),
			).Delete().Exec(ctx)
			if err != nil {
				return fmt.Errorf("在庫テーブル削除エラー : %w", err)
			}
		}

		// Stockテーブルの内容を一括取得
		stock, err := client.Stock.FindMany().Exec(ctx)
		if err != nil {
			return fmt.Errorf("在庫テーブル取得エラー : %w", err)
		}

		// 比較用に時刻を設定
		etime := stock[0].End

		for _, value := range stock {
			// fmt.Printf("index : %d, value : %+v\n", index, value)

			// 終了時刻がより早い場合はその時間をセット
			if value.End.Before(etime) {
				etime = value.End
			}
		}

		fmt.Println("etime :", etime)

		// 現在時刻との間隔を求める
		dtime := etime.Sub(GetTime())
		fmt.Println(dtime)

		// dtimeがdelayよりも早い場合
		// その間隔分遅延し、遅延後にタスク処理を実行
		if dtime < d {

		}

		time.Sleep((d * 1) * time.Second)
	}
}
