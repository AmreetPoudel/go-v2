package epoch

import (
	"fmt"
	"time"
)

func main() {
	// layout := "2006-01-02 15:04:05"
	ref_date := "2006-01-02"
	ref_date2 := "2006/02/01"

	// now := time.Now()
	// unixTime := now.Unix()

	// fmt.Println(now)
	// fmt.Println(unixTime)
	// unix := time.Unix(unixTime, 0)
	// fmt.Println(unix)

	d1, _ := time.Parse(ref_date, "2024-02-12")
	d2, _ := time.Parse(ref_date, "2024-02-11")

	fmt.Println(d1)
	fmt.Println(d2)
	//comparing two time
	fmt.Println(d2.Compare(d1))
	// add 7 days to parsed date
	fmt.Println(d1)
	fmt.Println(d1.AddDate(0, 0, 7))
	//check weekday
	fmt.Println(d1.Weekday())
	// formating date differently
	d3, _ := time.Parse(ref_date2, "2024/02/11")

	fmt.Println(d2)
	fmt.Println(d3)

}
