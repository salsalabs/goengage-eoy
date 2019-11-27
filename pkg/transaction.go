package eoy

import (
	goengage "github.com/salsalabs/goengage/pkg"
)

//Transaction reads a channel of activities to retrieve TransactionIDs.  Those
//are used to populate the Transaction table in the database.
func Transaction(rt *Runtime, c chan goengage.Fundraise) (err error) {
	rt.Log.Println("Transaction: start")
	for true {
		r, ok := <-c
		if !ok {
			break
		}
		rt.Log.Printf("Transaction: %v\n", r.ActivityID)

		if len(r.Transactions) != 0 {
			for _, c := range r.Transactions {
				rt.DB.Create(&c)
			}
		}
		return nil
	}
	rt.Log.Println("Transaction: start")
	return nil
}