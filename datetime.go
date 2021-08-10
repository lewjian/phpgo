package phpgo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultDateTimeFormatTpl = "2006-01-02 15:04:05"
	DefaultDateFormatTpl     = "2006-01-02"
	DefaultTimeFormatTpl     = "15:04:05"
)

// StrToTime 模拟PHP的strtotime函数
// 支持format有："+1 year"、"-1 year"、"+5 months"、"1 year +5 month -1 day +10 hours -23 minutes +5 seconds"等
// relativeTimestamp是时间戳，单位秒，之所以设为可变参数，是为了和php一致，使用时可以省略，省略则默认当前时间
// relativeTimestamp传多个值没用，只用到relativeTimestamp[0]
func StrToTime(format string, relativeTimestamp ...int64) (time.Time, error) {
	// 标识是否为负号，1为正；-1为负
	symbol := 1
	// 预处理一下format，首先将多个空格全都替换为一个空格，再去掉收尾空格，得到格式如：+1 year -1 month
	reg, err := regexp.Compile(`\s+`)
	if err != nil {
		return time.Time{}, err
	}
	format = reg.ReplaceAllString(format, " ")
	reg, err = regexp.Compile(`\s+`)
	if err != nil {
		return time.Time{}, err
	}
	var ts int64
	if len(relativeTimestamp) > 0 {
		ts = relativeTimestamp[0]
	} else {
		ts = time.Now().Unix()
	}
	data := reg.Split(format, -1)
	base := time.Unix(ts, 0)
	num := -1
	allowTimePeriod := []string{"year", "month", "day", "hour", "minute", "second"}
	for i, str := range data {
		if (i+1)%2 == 1 {
			symbol = 1
			// 去掉前后空格
			if str[0] == '-' {
				symbol = -1
				str = str[1:]
			} else if str[0] == '+' {
				str = str[1:]
			}
			// 奇数，该值必须是一个数字
			n, err := strconv.Atoi(str)
			if err != nil {
				return time.Time{}, fmt.Errorf("format参数格式不正确")
			}
			num = n
		} else {
			// 偶数，必须是year/month/day/hour/minute/second/
			// 尝试去掉最后一个s，如果存在
			timePeriod := strings.TrimRight(str, "s")
			if !InArrayString(timePeriod, allowTimePeriod) || num < 0 {
				return time.Time{}, fmt.Errorf("format参数格式不正确, %s不支持", str)
			}
			count := symbol * num
			switch timePeriod {
			case "year":
				base = base.AddDate(count, 0, 0)
			case "month":
				base = base.AddDate(0, count, 0)
			case "day":
				base = base.AddDate(0, 0, count)
			case "hour":
				base = base.Add(time.Duration(count) * time.Hour)
			case "minute":
				base = base.Add(time.Duration(count) * time.Minute)
			case "second":
				base = base.Add(time.Duration(count) * time.Second)
			}
			// 重置
			num = -1
		}
	}
	return base, nil
}

// Date PHP的date函数
// timestamp 时间戳，传了时间戳则以此为准，不传默认当前时间
// 用法参考PHP函数：https://www.php.net/manual/zh/function.date.php
// 但是不是每一个标志都实现了，具体可以参考下面的map对照表
func Date(format string, timestamp ...int64) string {
	phpGoDateMap := map[int32]string{
		'Y': "2006",
		'y': "06",
		'm': "01",
		'M': "Jan",
		'F': "January",
		'd': "02",
		'D': "Mon",
		'j': "2",
		'H': "15",
		'h': "03",
		'g': "3",
		'i': "04",
		's': "05",
		'u': ".000",
		'a': "pm",
		'A': "PM",
		'e': "MST",
		'O': "-0700",
		'P': "-07:00",
	}
	var goFormat strings.Builder
	for _, letter := range format {
		if goLetter, ex := phpGoDateMap[letter]; ex {
			goFormat.WriteString(goLetter)
		} else {
			goFormat.WriteByte(byte(letter))
		}
	}
	t := time.Now()
	if len(timestamp) > 0 {
		t = time.Unix(timestamp[0], 0)
	}
	return t.Format(goFormat.String())
}

// Time 获取当前时间戳，单位：秒
func Time() int64 {
	return time.Now().Unix()
}
