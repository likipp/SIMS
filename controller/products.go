package controller

import (
	"SIMS/global"
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils"
	"SIMS/utils/generate"
	"SIMS/utils/msg"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"os"
	"path"
	"strings"
	"time"
)

func CCreateProduct(c *gin.Context) {
	var product *models.Products
	err := gins.ParseJSON(c, &product)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	err = validation.Validate(product, validation.NotNil)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}

	err, data, success := services.SCreateProduct(product)
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetProductsSelectList(c *gin.Context) {
	var products []models.ProductsSelect
	param := c.DefaultQuery("param", "")
	if param == "{}" {
		param = ""
	}
	err, products, success := services.SGetProductsSelectList(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(products, err, 1, true, c)
}

func CGetProductsList(c *gin.Context) {
	var params models.ProductQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	err, list, total, success := new(models.Products).GetProductsList(params)
	if !success {
		msg.Result(nil, err, 1, success, c)
		return
	}
	msg.ResultWithPageInfo(list, err, 0, success, total, params.Current, params.PageSize, c)
}

func CUpdateProduct(c *gin.Context) {
	var product models.Products
	err := gins.ParseJSON(c, &product)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	// 获取到UUID, 只有uuid有值时才能更新成功
	id := c.Param("id")
	product.ID = utils.StringConvInt(id)
	err, success := product.UpdateProduct()
	if !success {
		msg.Result(nil, err, 1, success, c)
		return
	}
	msg.Result(nil, err, 0, true, c)
}

func CUploadPicture(c *gin.Context) {
	data := make(map[string]string)
	timeNow := time.Now().Format("20060102150405")
	file, err := c.FormFile("file")
	if err != nil {
		msg.Result(nil, msg.PictureUploadFailed, 1, false, c)
		return
	}
	brand := c.PostForm("brand")
	fileExt := strings.ToLower(path.Ext(file.Filename))
	filName := fmt.Sprintf("%s%s", timeNow, fileExt)
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		msg.Result(nil, msg.PictureExtFailed, 1, false, c)
		return
	}
	paths := fmt.Sprintf("%s/%s/", global.ImagePath, brand)
	if _, err = os.Stat(paths); err != nil {
		err = os.MkdirAll(paths, 0711)
		if err != nil {
			msg.Result(nil, msg.CreateFolderFailed, 1, false, c)
			return
		}
	} else {
		fmt.Println("文件夹已存在")
	}
	filePath := fmt.Sprintf("%s%s", paths, filName)
	filePathFS := fmt.Sprintf("%s%s", fmt.Sprintf("%s/%s/", global.ImageFS, brand), filName)
	fileUrlPath := fmt.Sprintf("%s/%s/%s", "http:/", c.Request.Host, filePathFS)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		msg.Result(nil, msg.PictureUploadFailed, 1, false, c)
		return
	}
	data["image_url"] = fileUrlPath
	data["image_path"] = filePath
	msg.Result(data, msg.PictureUploadSuccess, 0, true, c)
}

func CGenerateProductNumber(c *gin.Context) {
	p := c.DefaultQuery("parent", "")
	fmt.Println(p, "前端信息")
	number := generate.NumberWithProduct(p)
	msg.Result(number, errors.New("创建编号成功"), 0, true, c)
}

func CDeleteProduct(c *gin.Context) {
	fmt.Println("删除产品")
	msg.Result(nil, errors.New("创建编号成功"), 0, true, c)
}
