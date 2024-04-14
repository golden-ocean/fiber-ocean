package database

import (
	"context"
	"fmt"

	"github.com/golden-ocean/fiber-ocean/ent"
)

// import (
// 	"github.com/jmoiron/sqlx"
// )

//	func WithTransaction(db *sqlx.DB, txFunc func(*sqlx.Tx) error) error {
//		tx, err := db.Beginx()
//		if err != nil {
//			return err
//		}
//		defer func() {
//			if p := recover(); p != nil {
//				tx.Rollback()
//				panic(p) // re-throw panic after Rollback
//			} else if err != nil {
//				tx.Rollback() // err is non-nil; don't change it
//			} else {
//				err = tx.Commit() // err is nil; if Commit returns error update err
//			}
//		}()
//		err = txFunc(tx)
//		return err
//	}
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
