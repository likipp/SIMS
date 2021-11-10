package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"strconv"
)

type Unit struct {
	BaseModel
	Name string `json:"name" gorm:"comment:'名称'"`
}

type UnitSelect struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Key   string `json:"key"`
}

func (u *Unit) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required.Error("单位名称不能为空")),
	)
	return err
}

func getUnitDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(Unit))
}

func (u *Unit) CreateUnit() error {
	db := getUnitDB()
	err := entity.CheckExist(db, "name", u.Name)
	if err != nil {
		return err
	}
	err = db.Create(&u).Error
	if err != nil {
		return msg.CreatedFail
	}
	return msg.CreatedSuccess
}

func GetUnitSelectList(param string) (err error, list []UnitSelect, success bool) {
	var us UnitSelect
	var usl []UnitSelect
	var ul []Unit
	con := fmt.Sprintf("%s%s%s", "%", param, "%")
	db := getUnitDB()
	if param != "" {
		err = db.Where("name like ?", con).Find(&ul).Error
		if err != nil {
			return msg.GetFail, list, false
		}
	}
	if err = db.Find(&ul).Error; err != nil {
		return msg.GetFail, list, false
	}
	for i := range ul {
		us.Value = strconv.Itoa(ul[i].ID)
		us.Key = strconv.Itoa(ul[i].ID)
		us.Label = ul[i].Name
		usl = append(usl, us)
	}
	return msg.GetSuccess, usl, true
}
