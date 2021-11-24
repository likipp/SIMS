package models

import (
	"SIMS/global"
	"SIMS/utils/msg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// BillHeader 单据头
type BillHeader struct {
	BaseModel
	StockType    string  `json:"bill_type" gorm:"comment:'单据类型'"`
	Number       string  `json:"bill_number" gorm:"comment:'单号'"`
	Custom       int     `json:"custom" gorm:"comment:'客户'"`
	PayMethod    string  `json:"pay_method"  gorm:"comment:'收款方式'"`
	Status       int     `json:"status" gorm:"comment:'状态'"`
	BillAmount   float32 `json:"bill_amount" gorm:"订单金额"`
	RemainAmount float32 `json:"remain_amount" gorm:"剩余金额"`
}

// ExBillDetail 出库单详情
type ExBillDetail struct {
	Custom     string `json:"c_number"`
	CustomName string `json:"c_name"`
	BillHeader
	Body []BillEntry `json:"body"`
}

// InBillDetail 入库单详情
type InBillDetail struct {
	BillHeader
	Body []BillEntry `json:"body"`
}

func (sh *BillHeader) Validate() error {
	err := validation.ValidateStruct(sh,
		validation.Field(&sh.Number, validation.Required.Error("单号不能为空")),
		validation.Field(&sh.StockType, validation.Required.Error("出入库类型不能为空")),
		validation.Field(&sh.PayMethod, validation.When(sh.StockType == global.Ex, validation.Required.Error("收款方式不能为空"))),
		validation.Field(&sh.Custom, validation.When(sh.StockType == global.Ex, validation.Required.Error("客户不能为空"))),
	)
	return err
}

func (sh *BillHeader) BillLog(sb []BillEntry) (err error, success bool) {
	var billTotal float32
	var total int64
	// 校验字段是否满足条件
	err = validation.Validate(sh, validation.NotNil)
	if err != nil {
		return err, false
	}
	// 创建单据表头信息
	if sh.StockType == "出库单" {
		sh.Status = 1
	} else {
		sh.Status = 0
		sh.RemainAmount = 0
	}
	// 开始数据库事务
	tx := global.GDB.Begin()
	// 先判断单据号是否已经生成, 如果存在, 不允许继续新建
	tx.Model(&BillHeader{}).Where("number = ?", sh.Number).Count(&total)
	if total > 0 {
		return msg.Exists, false
	}
	// 开始创建订单表头信息, 不成功, 回滚不继续执行
	if err = tx.Create(&sh).Error; err != nil {
		return msg.CreatedFail, false
	}
	for i, length := 0, len(sb); i < length; i++ {
		sb[i].HeaderID = sh.ID
		//stock := GetWareHouseQtyWithProduct(sb[i].WareHouse, sb[i].PNumber, tx)
		if sh.StockType == global.In {
			//sb[i].InQTY = stock.QTY + sb[i].InQTY
			err, success = sb[i].InStockLog(tx)
			if !success {
				tx.Rollback()
				return err, false
			}
			// 将明细金额汇总, 填写到表头信息中, 方便与供应商结算
			billTotal += sb[i].Total
			continue
		}
		// 判断出入库类型, 如果类型属于出库, ExStockLog
		err, success = sb[i].ExStockLog(tx)
		if !success {
			tx.Rollback()
			return err, false
		}
	}
	err = tx.Model(&BillHeader{}).Where("number = ?", sh.Number).Update("bill_amount", billTotal).Error
	if err != nil {
		tx.Rollback()
		return msg.CreatedFail, false
	}
	tx.Commit()
	return msg.CreatedSuccess, true
}

func DeleteBillByID(number string) (err error, success bool) {
	var sh BillHeader
	var sb []BillEntry
	err = global.GDB.Model(&BillHeader{}).Where("number = ?", number).Find(&sh).Error
	if err != nil {
		return msg.GetFail, false
	}
	err = global.GDB.Where("header_id = ?", sh.ID).Find(&sb).Error
	if err != nil {
		return msg.GetFail, false
	}

	tx := global.GDB.Begin()
	for _, v := range sb {
		stock := GetWareHouseQtyWithProduct(v.WareHouse, v.PNumber, tx)
		if sh.StockType == global.In {
			if stock.QTY < v.InQTY {
				tx.Rollback()
				return msg.ExGTStock, false
			}
			if err = tx.Delete(&sb).Error; err != nil {
				tx.Rollback()
				return msg.DeletedFail, false
			}
			if stock.QTY-v.InQTY == 0 {
				if err = tx.Delete(&stock).Error; err != nil {
					tx.Rollback()
					return msg.DeletedFail, false
				}
			}
			if err = tx.Model(stock).Update("qty", stock.QTY-v.InQTY).Error; err != nil {
				tx.Rollback()
				return msg.UpdatedFail, false
			}
		}
		if err = tx.Delete(&sb).Error; err != nil {
			tx.Rollback()
			return msg.DeletedFail, false
		}
		if err = tx.Model(stock).Update("qty", stock.QTY+v.ExQTY).Error; err != nil {
			tx.Rollback()
			return msg.UpdatedFail, false
		}
	}
	if err = tx.Delete(&sh).Error; err != nil {
		tx.Rollback()
		return msg.DeletedFail, false
	}
	tx.Commit()
	return msg.DeletedSuccess, true
}

func UpdateDataByIn(tx *gorm.DB, Nsb BillEntry, sh BillHeader, stock Stock, QTY int, billTotal float32) (err error, success bool) {
	if err = tx.Model(Nsb).Updates(Nsb).Error; err != nil {
		tx.Rollback()
		return msg.UpdatedFail, false
	}
	if err = tx.Model(stock).Update("qty", QTY).Error; err != nil {
		tx.Rollback()
		return msg.UpdatedFail, false
	}
	if sh.BillAmount != billTotal {
		if err = tx.Model(sh).Update("bill_amount", billTotal).Error; err != nil {
			tx.Rollback()
			return msg.UpdatedFail, false
		}
	}
	return msg.UpdatedSuccess, true
}

func UpdateDataByEx(tx *gorm.DB, Nsb BillEntry, sh BillHeader, stock Stock, QTY int) (err error, success bool) {
	if err = tx.Model(Nsb).Updates(Nsb).Error; err != nil {
		tx.Rollback()
		return msg.UpdatedFail, false
	}
	if err = tx.Model(stock).Update("qty", QTY).Error; err != nil {
		tx.Rollback()
		return msg.UpdatedFail, false
	}
	fmt.Println(sh, "单据表头")
	if err = tx.Model(sh).Updates(sh).Error; err != nil {
		tx.Rollback()
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}

func UpdateDataHeaderByEx(tx *gorm.DB, sh BillHeader) (err error, success bool) {
	if err = tx.Model(sh).Updates(sh).Error; err != nil {
		tx.Rollback()
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}

// UpdateBillByNumber 待完成
func UpdateBillByNumber(sh BillHeader, sb []BillEntry) (err error, success bool) {
	//var sh BillHeader
	var shOld BillHeader
	var sbOld []BillEntry
	var billTotal float32
	for _, v := range sb {
		billTotal += v.Total
	}
	tx := global.GDB.Begin()
	// 通过ID值获取到对应的表单头部数据
	if err = tx.Where("number = ?", sh.Number).Find(&shOld).Error; err != nil {
		return msg.GetFail, false
	}
	fmt.Println("旧数据", shOld)
	sh.ID = shOld.ID
	// 通过HeaderID获取对应的表单明细数据
	if err = tx.Where("header_id = ?", sh.ID).Find(&sbOld).Error; err != nil {
		return msg.GetFail, false
	}
	// 如果前端传递的明细行数量小于原始数据, 走小于逻辑的判断
	if len(sbOld) >= len(sb) {
		if len(sb) == 0 {
			err, success = UpdateDataHeaderByEx(tx, sh)
			if !success {
				return err, false
			}
		}

		// 判断单据类型
		if sh.StockType == global.In {
			// 循环原始数据
			for i := 0; i < len(sbOld); i++ {
				// 获取明细行产品的库存信息
				stock := GetWareHouseQtyWithProduct(sbOld[i].WareHouse, sbOld[i].PNumber, tx)
				// 如果原始数据与新数据的产品是一样的, 就只查看仓库, 数量, 单价是否变更
				if sbOld[i].PNumber == sb[i].PNumber {
					// 判断仓库是否一样, 如果一样, 更新数量, 单价, 金额
					if sbOld[i].WareHouse == sb[i].WareHouse {
						// 如果原始数据的入库数量大于新数据的入库数量
						if sbOld[i].InQTY > sb[i].InQTY {
							// 如果产品的库存数量小于（原始数据-新数据的数量), 则说明库存不足以修改, 报错返回
							if stock.QTY < sbOld[i].InQTY-sb[i].InQTY {
								return msg.ExGTStock, false
							}
							err, success = UpdateDataByIn(tx, sb[i], sh, *stock, stock.QTY-(sbOld[i].InQTY-sb[i].InQTY), billTotal)
							if !success {
								return err, false
							}
							continue
						}
						err, success = UpdateDataByIn(tx, sb[i], sh, *stock, stock.QTY+(sb[i].InQTY-sbOld[i].InQTY), billTotal)
						if !success {
							return err, false
						}
						continue
					}

					// 判断仓库发生变化
					if stock.QTY < sbOld[i].InQTY {
						return msg.ExGTStock, false
					}
					if stock.QTY == sbOld[i].InQTY {
						if err = tx.Delete(&stock).Error; err != nil {
							return msg.DeletedFail, false
						}
					}
					err, success = UpdateDataByIn(tx, sb[i], sh, *stock, stock.QTY+(sb[i].InQTY-sbOld[i].InQTY), billTotal)
					if !success {
						return err, false
					}
					continue
				}
				if stock.QTY < sbOld[i].InQTY {
					return msg.ExGTStock, false
				}
				if err = tx.Delete(sbOld[i]).Error; err != nil {
					return msg.DeletedFail, false
				}
				if err = tx.Create(&sb[i]).Error; err != nil {
					return msg.CreatedFail, false
				}
				if err = tx.Model(stock).Update("qty", stock.QTY+(sb[i].InQTY-sbOld[i].InQTY)).Error; err != nil {
					return msg.UpdatedFail, false
				}
				if sh.BillAmount != billTotal {
					if err = tx.Model(sh).Update("bill_amount", billTotal).Error; err != nil {
						return msg.UpdatedFail, false
					}
				}
			}
		} else {
			for i := 0; i < len(sbOld); i++ {
				// 获取明细行产品的库存信息
				stock := GetWareHouseQtyWithProduct(sbOld[i].WareHouse, sbOld[i].PNumber, tx)
				fmt.Println(stock, "库存")
				// 如果原始数据与新数据的产品是一样的, 就只查看仓库, 数量, 单价是否变更
				if sbOld[i].PNumber == sb[i].PNumber {
					// 判断仓库是否一样, 如果一样, 更新数量, 单价, 金额
					if sbOld[i].WareHouse == sb[i].WareHouse {
						fmt.Println("到了比较新旧数据")
						// 如果原始数据的入库数量大于新数据的入库数量
						if sb[i].ExQTY > sbOld[i].ExQTY {
							// 如果产品的库存数量小于（原始数据-新数据的数量), 则说明库存不足以修改, 报错返回
							if stock.QTY < sb[i].ExQTY - sbOld[i].ExQTY{
								return msg.ExGTStock, false
							}
							err, success = UpdateDataByEx(tx, sb[i], sh, *stock, stock.QTY-(sb[i].ExQTY - sbOld[i].ExQTY))
							if !success {
								return err, false
							}
							continue
						}
						err, success = UpdateDataByEx(tx, sb[i], sh, *stock, stock.QTY+(sbOld[i].ExQTY-sb[i].ExQTY))
						if !success {
							return err, false
						}
						continue
					}

					// 判断仓库发生变化
					if stock.QTY < sbOld[i].ExQTY {
						return msg.ExGTStock, false
					}
					if stock.QTY == sbOld[i].ExQTY {
						if err = tx.Delete(&stock).Error; err != nil {
							return msg.DeletedFail, false
						}
					}
					err, success = UpdateDataByEx(tx, sb[i], sh, *stock, stock.QTY-sbOld[i].ExQTY)
					if !success {
						return err, false
					}
					continue
				}
				if stock.QTY < sbOld[i].ExQTY {
					return msg.ExGTStock, false
				}
				if err = tx.Delete(sbOld[i]).Error; err != nil {
					return msg.DeletedFail, false
				}
				if err = tx.Create(&sb[i]).Error; err != nil {
					return msg.CreatedFail, false
				}
				if err = tx.Model(stock).Update("qty", stock.QTY-sbOld[i].ExQTY).Error; err != nil {
					return msg.UpdatedFail, false
				}
			}
		}
	}
	//else {
	//	if sh.StockType == global.In {
	//		// 循环原始数据
	//		for i := 0; i < len(sbOld); i++ {
	//			// 获取明细行产品的库存信息
	//			stock := GetWareHouseQtyWithProduct(sbOld[i].WareHouse, sbOld[i].PNumber, tx)
	//			// 如果原始数据与新数据的产品是一样的, 就只查看仓库, 数量, 单价是否变更
	//			if sbOld[i].PNumber == sb[i].PNumber {
	//				// 判断仓库是否一样, 如果一样, 更新数量, 单价, 金额
	//				if sbOld[i].WareHouse == sb[i].WareHouse {
	//					// 如果原始数据的入库数量大于新数据的入库数量
	//					if sbOld[i].InQTY > sb[i].InQTY {
	//						// 如果产品的库存数量小于（原始数据-新数据的数量), 则说明库存不足以修改, 报错返回
	//						if stock.QTY < sbOld[i].InQTY-sb[i].InQTY {
	//							return msg.ExGTStock, false
	//						}
	//						err, success = UpdateDataByIn(tx, sb[i], sh, *stock, stock.QTY-(sbOld[i].InQTY-sb[i].InQTY), billTotal)
	//						if !success {
	//							return err, false
	//						}
	//						continue
	//					}
	//					err, success = UpdateDataByIn(tx, sb[i], sh, *stock, stock.QTY+(sb[i].InQTY-sbOld[i].InQTY), billTotal)
	//					if !success {
	//						return err, false
	//					}
	//					continue
	//				}
	//
	//				// 判断仓库发生变化
	//				if stock.QTY < sbOld[i].InQTY {
	//					return msg.ExGTStock, false
	//				}
	//				if stock.QTY == sbOld[i].InQTY {
	//					if err = tx.Delete(&stock).Error; err != nil {
	//						return msg.DeletedFail, false
	//					}
	//				}
	//				err, success = UpdateDataByIn(tx, sb[i], sh, *stock, stock.QTY+(sb[i].InQTY-sbOld[i].InQTY), billTotal)
	//				if !success {
	//					return err, false
	//				}
	//				continue
	//			}
	//			if stock.QTY < sbOld[i].InQTY {
	//				return msg.ExGTStock, false
	//			}
	//			if err = tx.Delete(sbOld[i]).Error; err != nil {
	//				return msg.DeletedFail, false
	//			}
	//			if err = tx.Create(&sb[i]).Error; err != nil {
	//				return msg.CreatedFail, false
	//			}
	//			if err = tx.Model(stock).Update("qty", stock.QTY+(sb[i].InQTY-sbOld[i].InQTY)).Error; err != nil {
	//				return msg.UpdatedFail, false
	//			}
	//			if sh.BillAmount != billTotal {
	//				if err = tx.Model(sh).Update("bill_amount", billTotal).Error; err != nil {
	//					return msg.UpdatedFail, false
	//				}
	//			}
	//		}
	//		for i := 0; i < len(sb) - len(sbOld); i++ {
	//			fmt.Println("会到此处来")
	//			if err = tx.Create(&sb[len(sbOld)+i]).Error; err != nil {
	//				return msg.CreatedFail, false
	//			}
	//		}
	//	}
	//}

	tx.Commit()
	return msg.UpdatedSuccess, true
}

// GetExBillDetail 获取出库订单详情信息
func GetExBillDetail(number string) (error, ExBillDetail, bool) {
	var header BillHeader
	var body []BillEntry
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

// GetInBillDetail 获取采购订单详情
func GetInBillDetail(number string) (error, InBillDetail, bool) {
	var header BillHeader
	var body []BillEntry
	var b InBillDetail
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
	b.Body = body
	return msg.GetSuccess, b, true
}
