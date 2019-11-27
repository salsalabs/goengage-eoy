package eoy

//harvester declares functions that process data.
type harvester func(rt *Runtime) (err error)

//Harvest retrieves data from the database in various permutations of slicing
//and dicing, then stores them into a spreadsheet.  The spreadsheet is written
//to disk when done.
func (rt *Runtime) Harvest() (err error) {
	functions := []harvester{
		ThisYear,
		Months,
		YearOverYear,
		MonthOverMonth,
		AllDonors,
		TopDonors,
		ActivityPages,
		ProjectedRevenue,
	}

	for _, r := range functions {
		err := r(rt)
		if err != nil {
			return err
		}
	}
	return err
}

// ThisYear selects data for ThisYear, sorts it, tweaks it, then stores it into
//the spreadsheet.
func ThisYear(rt *Runtime) (err error) {

	return err
}

// Months selects data for Months, sorts it, tweaks it, then stores it into
//the spreadsheet.
func Months(rt *Runtime) (err error) {

	return err
}

// YearOverYear selects data for YearOverYear, sorts it, tweaks it, then stores it into
//the spreadsheet.
func YearOverYear(rt *Runtime) (err error) {

	return err
}

// MonthOverMonth selects data for MonthOverMonth, sorts it, tweaks it, then stores it into
//the spreadsheet.
func MonthOverMonth(rt *Runtime) (err error) {

	return err
}

// AllDonors selects data for AllDonors, sorts it, tweaks it, then stores it into
//the spreadsheet.
func AllDonors(rt *Runtime) (err error) {

	return err
}

// TopDonors selects data for TopDonors, sorts it, tweaks it, then stores it into
//the spreadsheet.
func TopDonors(rt *Runtime) (err error) {

	return err
}

// ActivityPages selects data for ActivityPages, sorts it, tweaks it, then stores it into
//the spreadsheet.
func ActivityPages(rt *Runtime) (err error) {

	return err
}

// ProjectedRevenue selects data for ProjectedRevenue, sorts it, tweaks it, then stores it into
//the spreadsheet.
func ProjectedRevenue(rt *Runtime) (err error) {

	return err
}
