package timex

import (
	"fmt"
	"time"
)

func main() {
	x,y,z := GetYearMonthDay(TimeCurrString())
	fmt.Println(x,y,z)
}

//****************************************************************************
func GetYearMonthDay(strTime string) (int, int, int) {
	d, err := time.Parse("2006-01-02 15:04:05", strTime)
	if err != nil {
		return -1, -1, -1
	}
	return d.Year(), int(d.Month()), d.Day()
}
func GetYearMonthDayEx(nTime int64) (int, int, int) {
	d, err := time.Parse("2006-01-02 15:04:05", time.Unix(nTime, 0).Format("2006-01-02 15:04:05"))
	if err != nil {
		return -1, -1, -1
	}
	return int(d.Year()), int(d.Month()), int(d.Day())
}

//****************************************************************************
//1599809584(2020-09-11 15:33:04) => 1599753600(2020-09-11 00:00:00)
func GetThisMorning(nTime int64) int64 {
	t := time.Unix(nTime, 0)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

//-------------------------------
//1599809584(2020-09-11 15:33:04) => 1599840000(2020-09-12 00:00:00)
func GetNextMorning(nTime int64) int64 {
	t := time.Unix(nTime, 0)
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location()).Unix()
}

//****************************************************************************
//获取当前时间  2020-09-10 15:30:00
func TimeCurrString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//获取当前时间戳 1599805463
func TimeCurrInt64() (nCurrTimeStamp int64) {
	return time.Now().Unix()
}

//时间戳转时间  1599805463 => 2020-09-10 15:30:00
func TimeInt64ToTimeString(nTimeStamp int64) (strTime string) {
	tm := time.Unix(nTimeStamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//时间转时间戳 2020-09-10 15:30:00 => 1599805463
func TimeStringToTimeInt64(strTime string) (nTimeStamp int64) {
	timestamp, _ := time.Parse("2006-01-02 15:04:05", strTime)
	return timestamp.Unix() - 3600*8
}

//time.Time => int64
func TimetToInt64(tTime time.Time) (nTime int64) {
	return time.Date(tTime.Year(), tTime.Month(), tTime.Day(), tTime.Hour(),
		tTime.Minute(), tTime.Second(), 0, tTime.Location()).Unix()
}

//****************************************************************************
//某个时间之前多久，或者之后多久(精确时间)
//after_time := TimeBeforeOrLater(TimeCurrInt64(), "after", 0, 0, 3, 0, 0, 0)
//三天之后  1582848900(2020-02-28 08:15:00) => 1583108100(2020-03-02 08:15:00)
//-------------------------------
//before_time := TimeBeforeOrLater(TimeCurrInt64(), "before", 0, 0, 3, 0, 0, 0)
//三天之前  1583108100(2020-03-02 08:15:00) => 1582848900(2020-02-28 08:15:00)
func TimeBeforeOrLater(curr_time int64, tag string, year_diff int, month_diff time.Month, day_diff, hour_diff, min_diff, sec_diff int) (target_time int64) {
	tTime := time.Unix(curr_time, 0).Local()
	if tag == "before" {
		return time.Date(tTime.Year()-year_diff, tTime.Month()-month_diff, tTime.Day()-day_diff, tTime.Hour()-hour_diff,
			tTime.Minute()-min_diff, tTime.Second()-sec_diff, 0, tTime.Location()).Unix()
	} else if tag == "after" {
		return time.Date(tTime.Year()+year_diff, tTime.Month()+month_diff, tTime.Day()+day_diff, tTime.Hour()+hour_diff,
			tTime.Minute()+min_diff, tTime.Second()+sec_diff, 0, tTime.Location()).Unix()
	} else {
		return 0
	}
}

//****************************************************************************
func Getconstellation(birthday int64) (star string) {
	d, err := time.Parse("2006-01-02", time.Unix(birthday, 0).Format("2006-01-02"))
	if err != nil {
		fmt.Println("日期解析错误！")
		return ""
	}
	month := int(d.Month())
	day := d.Day()
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "魔蝎座"
	}
	return star
}

