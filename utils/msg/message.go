package msg

import "errors"

var (
	QueryParamsFail      = errors.New("解析请求参数发生错误")
	CreatedSuccess       = errors.New("创建成功")
	CreatedFail          = errors.New("创建失败")
	UpdatedSuccess       = errors.New("更新成功")
	UpdatedFail          = errors.New("更新失败")
	DeletedSuccess       = errors.New("删除成功")
	DeletedFail          = errors.New("删除失败")
	GetSuccess           = errors.New("查询成功")
	GetFail              = errors.New("查询失败")
	NotFound             = errors.New("未找到相关内容或者数据为空")
	TimeOut              = errors.New("操作超时")
	DoNothing            = errors.New("数据未变更")
	DuplicatedData       = errors.New("数据重复, 请检查已有数据")
	ExGTStock            = errors.New("出库数量大于实际库存")
	Copier               = errors.New("copier复制失败")
	LoginSuccess         = errors.New("登录成功")
	LoginFail            = errors.New("帐号密码错误")
	SaveSessionSuccess   = errors.New("保存session成功")
	SaveSessionFail      = errors.New("保存session失败")
	GetSessionSuccess    = errors.New("获取session成功")
	GetSessionFail       = errors.New("获取session失败")
	DeleteSessionSuccess = errors.New("删除session成功")
	DeleteSessionFail    = errors.New("删除session失败")
	LoginOutSuccess      = errors.New("退出成功")
	SessionTimeout       = errors.New("session已过期, 请重新登录")
	GetTokenFail         = errors.New("获取Token失败")
	GetTokenSuccess      = errors.New("获取Token成功")
	ExpectationFailed    = errors.New("未登录或token已过期")
	PaginationFailed     = errors.New("分页失败")
	AmountSuccess        = errors.New("货款已结清")
	PictureUploadFailed  = errors.New("图片上传失败")
	PictureUploadSuccess = errors.New("图片上传成功")
	CreateFolderFailed   = errors.New("创建文件夹失败")
	PictureExtFailed     = errors.New("上传失败!只允许png,jpg,gif,jpeg文件")
	Exists               = errors.New("数据已经存在")
)
