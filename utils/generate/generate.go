package generate

import (
	"SIMS/global"
	"SIMS/models"
	"SIMS/utils"
	"bytes"
	"fmt"
	"time"
)

type ProductNumber struct {
	Name        string     `json:"name"`
	Number      string     `json:"number"`
	PNumber     string     `json:"p_number"`
	ID          int        `json:"id"`
}

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
	//var p models.Products
	var pn ProductNumber
	var total int64
	var buf bytes.Buffer
	var NewNumberL string
	global.GDB.Select("brands.number, brands.id, brands.name, products.p_number").Model(&models.Products{}).Joins("left join brands on brands.id = products.brand").Where("brands.id = ?", parent).Last(&pn).Count(&total)
	//fmt.Println("total:", total)
	//if total > 0 {
	//}
	//if parent == "B" {
	//	buf.WriteString("00001")
	//} else {
	//	buf.WriteString("0001")
	//}
	//buf.WriteString(parent)
	NewNumberL = utils.IntConvJoinByProduct(len(pn.Number), pn.PNumber[len(pn.Number):])
	buf.WriteString(pn.Number)
	buf.WriteString(NewNumberL)
	return buf.String()
}
