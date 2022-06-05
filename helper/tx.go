package helper

import (
	"database/sql"
	"iteung-api/exception"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		exception.PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		exception.PanicIfError(errorCommit)
	}
}
