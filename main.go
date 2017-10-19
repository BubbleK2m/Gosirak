package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/DSMdongly/gogyo"
)

func ShowDailyMenu(school gogyo.School, year int, month int, day int) {
	fmt.Printf("\n[ %2d/%2d/%2d ]\n", year, month, day)

	daily := school.GetDailyMenu(year, month, day)
	times := [3]string{"조식", "중식", "석식"}

	for index := 0; index < 3; index++ {
		meals := daily[index]

		if meals == nil {
			continue
		}

		time := times[index]
		fmt.Printf("\n[ %s ]\n\n", time)

		for _, meal := range meals {
			fmt.Printf("- %s\n", meal.Name)
		}
	}
}

func main() {
	now := time.Now()

	year := flag.Int("year", now.Year(), "year")
	month := flag.Int("month", int(now.Month()), "month")
	day := flag.Int("day", now.Day(), "day")

	schoolKind := flag.String("school-kind", "high-school", "kind of school")
	schoolRegion := flag.String("school-region", "daejeon", "region of school")
	schoolCode := flag.String("school-code", "G100000170", "code of school")

	flag.Parse()

	kind := gogyo.GetSchoolKind(*schoolKind)
	region := gogyo.GetSchoolRegion(*schoolRegion)
	code := gogyo.SchoolCode(strings.ToUpper(*schoolCode))

	school := gogyo.NewSchool(kind, region, code)
	ShowDailyMenu(*school, *year, *month, *day)
}
