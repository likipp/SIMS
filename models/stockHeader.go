package models

import (
	"SIMS/global"
	"SIMS/utils"
	"SIMS/utils/msg"
	"bytes"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
	"time"
)

type StockHeader struct {
	BaseModel
	StockType string `json:"type" gorm:"comment:'单据类型'"`
	Number    string `json:"number" gorm:"comment:'单号'"`
	Custom    int    `json:"custom" gorm:"comment:'客户'"`
	//Courier   int    `json:"courier" gorm:"comment:'供应商'"`
	Discount  int    `json:"discount"  gorm:"comment:'折扣'"`
	PayMethod string `json:"pay_method"  gorm:"comment:'收款方式'"`
}

type ExBillDetail struct {
	Custom     string `json:"c_number"`
	CustomName string `json:"c_name"`
	StockHeader
	Body []StockBody `json:"body"`
}

func (sh *StockHeader) Validate() error {
	err := validation.ValidateStruct(sh,
		validation.Field(&sh.Number, validation.Required.Error("单号不能为空")),
		validation.Field(&sh.StockType, validation.Required.Error("出入库类型不能为空")),
		validation.Field(&sh.PayMethod, validation.Required.Error("收款方式不能为空")),
		validation.Field(&sh.Custom, validation.When(sh.StockType == global.Ex, validation.Required.Error("客户不能为空"))),
	)
	return err
}

func (sh *StockHeader) StockHeaderLog(sb []StockBody) (err error, success bool) {
	n := BillNumber(sh.StockType)
	sh.Number = n
	tx := global.GDB.Begin()
	if err = tx.Create(&sh).Error; err != nil {
		tx.Rollback()
		return msg.CreatedFail, false
	}
	for i := range sb {
		sb[i].HeaderID = sh.ID
		if sh.StockType == global.In {
			err, success = sb[i].InStockLog(tx)
			if !success {
				return err, false
			}
			continue
		}
		err, success = sb[i].ExStockLog(tx)
		if !success {
			return err, false
		}
	}
	tx.Commit()
	return err, true
}

func GetExBillDetail(number string) (error, ExBillDetail, bool) {
	var header StockHeader
	var body []StockBody
	var b ExBillDetail
	var c Custom
	err := global.GDB.Where("number = ?", number).Find(&header).Error
	if err != nil {
		return msg.GetFail, b, false
	}
	global.GDB.Where("id = ?", header.Custom).Find(&c)
	if err != nil {
		return msg.GetFail, b, false
	}
	err = global.GDB.Where("header_id = ?", header.ID).Find(&body).Error
	if err != nil {
		return msg.GetFail, b, false
	}
	if err = copier.Copy(&b, &header); err != nil {
		return msg.Copier, b, false
	}
	b.Custom = c.CNumber
	b.CustomName = c.CName
	b.Body = body
	return msg.GetSuccess, b, true
}

func BillNumber(t string) (n string) {
	var s StockHeader
	var total int64
	var buf bytes.Buffer
	timeNow := time.Now().Format("20060102")
	timeParam := fmt.Sprintf("%s%s%s", "%", timeNow, "%")
	global.GDB.Raw("select number from stock_headers where stock_type = ? and number like ? order by number desc limit 1", t, timeParam).Scan(&s).Count(&total)
	if total > 0 {
		number := s.Number
		lastString := number[len(number)-2:]
		newNumStr := utils.IntConvJoin(lastString)

		if t == global.Ex {
			buf.WriteString("EX")
			buf.WriteString(timeNow)
			buf.WriteString(newNumStr)
			return buf.String()
			//return fmt.Sprintf("%s%s%s", "EX", timeNow, newNumStr)

		}
		if t == global.In {
			buf.WriteString("IN")
			buf.WriteString(timeNow)
			buf.WriteString(newNumStr)
			return buf.String()
			//return fmt.Sprintf("%s%s%s", "IN", timeNow, newNumStr)
		}
	}
	if t == global.Ex {
		buf.WriteString("EX")
		buf.WriteString(timeNow)
		buf.WriteString("01")
		return buf.String()
		//return fmt.Sprintf("%s%s%s", "EX", timeNow, "01")
	}
	buf.WriteString("IN")
	buf.WriteString(timeNow)
	buf.WriteString("01")
	return buf.String()
	//return fmt.Sprintf("%s%s%s", "In", timeNow, "01")
}
