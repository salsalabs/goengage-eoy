package eoy

import (
	"time"

	goengage "github.com/salsalabs/goengage/pkg"
)

//ActivityForm contains a basic set of values for an activity page.
type ActivityForm struct {
	Key         string
	Name        string
	CreatedDate *time.Time
}

//Form reads a channel of activities to retrieve ActivityIDs.  Those
//are used to populate the Activity table in the database.
func Form(rt *Runtime, c chan goengage.Fundraise) (err error) {
	rt.Log.Println("Form: start")
	for true {
		r, ok := <-c
		if !ok {
			break
		}
		s := ActivityForm{}
		rt.Log.Printf("Form: %v\n", r.ActivityID)
		rt.DB.Where("key = ?", r.ActivityFormID).First(&s)
		if s.CreatedDate == nil {
			s.Key = r.ActivityFormID
			s.Name = r.ActivityFormName
			t := time.Now()
			s.CreatedDate = &t
			rt.DB.Create(&s)
		}
	}
	rt.Log.Println("Form: end")
	return nil
}