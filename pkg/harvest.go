package eoy

//harvester declares functions that process data.
type harvester func(rt *Runtime) (err error)

//Harvest retrieves data from the database in various permutations of slicing
//and dicing, then stores them into a spreadsheet.  The spreadsheet is written
//to disk when done.
func (rt *Runtime) Harvest(fn string) (err error) {
	functions := []harvester{
		ThisYear,
		Months,
		YearOverYear,
		MonthOverMonth,
		AllDonors,
		TopDonors,
		ActivityForms,
		// ProjectedRevenue,
	}

	for _, r := range functions {
		err := r(rt)
		if err != nil {
			return err
		}
	}
	rt.Spreadsheet.SetActiveSheet(2)
	rt.Spreadsheet.DeleteSheet("Sheet1")
	rt.Spreadsheet.DeleteSheet("Sheet2")
	err = rt.StoreSpreadsheet(fn)
	return err
}

// ThisYear selects data for ThisYear, sorts it, tweaks it, then stores it into
//the spreadsheet.
func ThisYear(rt *Runtime) (err error) {
	sheet := rt.NewThisYearSheet()
	rt.Decorate(sheet)
	return err
}

//Months select monthds for the largest year in the months database.
func Months(rt *Runtime) (err error) {
	sheet := rt.NewMonthSheet()
	rt.Decorate(sheet)
	return err
}

// YearOverYear selects data for this , sorts it, tweaks it, then stores it into
//the spreadsheet.
func YearOverYear(rt *Runtime) (err error) {
	sheet := rt.NewYOYearSheet()
	rt.Decorate(sheet)
	return err
}

// MonthOverMonth selects data for MonthOverMonth, sorts it, tweaks it, then stores it into
//the spreadsheet.
func MonthOverMonth(rt *Runtime) (err error) {
	sheet := rt.NewMOMonthSheet()
	rt.Decorate(sheet)
	return err
}

// AllDonors shows details for all donors in the current year.
func AllDonors(rt *Runtime) (err error) {
	sheet := rt.NewAllDonorsSheet()
	rt.Decorate(sheet)
	return err
}

// TopDonors shows details for all donors in the current year.
func TopDonors(rt *Runtime) (err error) {
	sheet := rt.NewTopDonorsSheet()
	rt.Decorate(sheet)
	return err
}

// ActivityForms selects data for ActivityForms, sorts it, tweaks it, then stores it into
//the spreadsheet.
func ActivityForms(rt *Runtime) (err error) {
	sheet := rt.NewActivityFormSheet()
	rt.Decorate(sheet)
	return err
}

// // ProjectedRevenue selects data for ProjectedRevenue, sorts it, tweaks it, then stores it into
// //the spreadsheet.
// func ProjectedRevenue(rt *Runtime) (err error) {
// 	sheet := "Projected revenue"
// 	_ = rt.Spreadsheet.NewSheet(sheet)
//
// 	return err
// }

//StoreSpreadsheet saves the spreadsheet to disk.
func (rt *Runtime) StoreSpreadsheet(fn string) (err error) {
	err = rt.Spreadsheet.SaveAs(fn)
	return err
}
