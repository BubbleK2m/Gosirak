package gosirak

import (
	"fmt"
	"strings"

	"github.com/DSMdongly/gogyo"
)

func GetSchoolKind(kind string) gogyo.SchoolKind {
	kinds := map[string]gogyo.SchoolKind{
		"kinder-garden":    gogyo.KinderGarden,
		"elementry-school": gogyo.ElementrySchool,
		"middle-school":    gogyo.MiddleSchool,
		"high-school":      gogyo.HighSchool,
	}

	return kinds[strings.ToLower(kind)]
}

func GetSchoolRegion(region string) gogyo.SchoolRegion {
	regions := map[string]gogyo.SchoolRegion{
		"seoul":     gogyo.Seoul,
		"incheon":   gogyo.Incheon,
		"busan":     gogyo.Busan,
		"gwangju":   gogyo.Gwangju,
		"daejeon":   gogyo.Daejeon,
		"daegu":     gogyo.Daegu,
		"sejong":    gogyo.Sejong,
		"ulsan":     gogyo.Ulsan,
		"gyeonggi":  gogyo.Gyeonggi,
		"kangwon":   gogyo.Kangwon,
		"chungbuk":  gogyo.Chungbuk,
		"chungnam":  gogyo.Chungnam,
		"gyeongbuk": gogyo.Gyeongbuk,
		"gyeongnam": gogyo.Gyeongnam,
		"jeonbuk":   gogyo.Jeonbuk,
		"jeonnam":   gogyo.Jeonnam,
		"jeju":      gogyo.Jeju,
	}

	return regions[strings.ToLower(region)]
}

func ShowDailyMenu(school gogyo.School, year int, month int, day int) {
	fmt.Printf("\n[ %s ]\n", school.Code)
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
