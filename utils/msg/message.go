package msg

import "errors"

var (
	QueryParamsFail = errors.New("解析请求参数发生错误")
	CreatedSuccess  = errors.New("创建成功")
	CreatedFail     = errors.New("创建失败")
	UpdatedSuccess  = errors.New("更新成功")
	UpdatedFail     = errors.New("更新失败")
	DeletedSuccess  = errors.New("删除成功")
	DeletedFail     = errors.New("删除失败")
	GetSuccess      = errors.New("查询成功")
	GetFail         = errors.New("查询失败")
	NotFound        = errors.New("未找到相关内容或者数据为空")
	TimeOut         = errors.New("操作超时")
	DoNothing       = errors.New("数据未变更")
	DuplicatedData  = errors.New("数据重复, 请检查已有数据")
	ExGTStock       = errors.New("出库数量大于实际库存")
	Copier          = errors.New("copier复制失败")
	LoginSuccess    = errors.New("登录成功")
	LoginFail       = errors.New("帐号密码错误")
)
