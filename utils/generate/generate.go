package generate

import (
	"SIMS/global"
	"SIMS/models"
	"SIMS/utils"
	"bytes"
	"fmt"
	"time"
)

func GenerateNumber(t string) (n string) {
	if t == "" {
		return ""
	}
	var s models.BillHeader
	var total int64
	var buf bytes.Buffer
	timeNow := time.Now().Format("20060102")
	timeParam := fmt.Sprintf("%s%s%s", "%", timeNow, "%")
	global.GDB.Where("stock_type = ? and number like ?", t, timeParam).Limit(1).Order("number").Find(&s).Count(&total)
	//global.GDB.Raw("select number from bill_headers where stock_type = ? and number like ? order by number desc limit 1", t, timeParam).Scan(&s).Count(&total)
	if total > 0 {
		number := s.Number
		lastString := number[len(number)-2:]
		newNumStr := utils.IntConvJoin(lastString)

		if t == global.Ex {
			buf.WriteString("EX")
			buf.WriteString(timeNow)
			buf.WriteString(newNumStr)
			return buf.String()

		}
		if t == global.In {
			buf.WriteString("IN")
			buf.WriteString(timeNow)
			buf.WriteString(newNumStr)
			return buf.String()
		}
	}
	if t == global.Ex {
		buf.WriteString("EX")
		buf.WriteString(timeNow)
		buf.WriteString("01")
		return buf.String()
	}
	buf.WriteString("IN")
	buf.WriteString(timeNow)
	buf.WriteString("01")
	return buf.String()
}

func GenerateNumberWithYF() (n string) {
	var s models.Payable
	var total int64
	var buf bytes.Buffer
	timeNow := time.Now().Format("20060102")
	timeParam := fmt.Sprintf("%s%s%s", "%", timeNow, "%")
	global.GDB.Where("number like ?", timeParam).Limit(1).Order("number").Find(&s).Count(&total)
	if total > 0 {
		number := s.Number
		lastString := number[len(number)-2:]
		newNumStr := utils.IntConvJoin(lastString)

		buf.WriteString("YF")
		buf.WriteString(timeNow)
		buf.WriteString(newNumStr)
		return buf.String()

	}
	buf.WriteString("YF")
	buf.WriteString(timeNow)
	buf.WriteString("01")
	return buf.String()
}

func NumberWithProduct(parent string) (n string) {
	var p models.Products
	var total int64
	var buf bytes.Buffer
	global.GDB.Joins("JOIN brands on brands.id = products.brand").Where("brands.number = ?", parent).Last(&p).Count(&total)
	if total > 0 {
		NewNumberL := utils.IntConvJoin(p.PNumber[3:])
		buf.WriteString(parent)
		buf.WriteString(NewNumberL)
		return buf.String()
	}
	buf.WriteString(parent)
	buf.WriteString("001")
	return buf.String()
}