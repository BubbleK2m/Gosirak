package main

import (
	"flag"
	"strings"
	"time"

	"github.com/DSMdongly/gogyo"
	"github.com/DSMdongly/gosirak/gosirak"
)

func main() {
	now := time.Now()

	year := flag.Int("year", now.Year(), "year")
	month := flag.Int("month", int(now.Month()), "month")
	day := flag.Int("day", now.Day(), "day")

	schoolKind := flag.String("school-kind", "high-school", "kind of school")
	schoolRegion := flag.String("school-region", "daejeon", "region of school")
	schoolCode := flag.String("school-code", "G100000170", "code of school")

	flag.Parse()

	kind := gosirak.GetSchoolKind(*schoolKind)
	region := gosirak.GetSchoolRegion(*schoolRegion)
	code := gogyo.SchoolCode(strings.ToUpper(*schoolCode))

	school := gogyo.NewSchool(kind, region, code)
	gosirak.ShowDailyMenu(*school, *year, *month, *day)
}
