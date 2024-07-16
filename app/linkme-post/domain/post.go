package domain

const (
	Draft     = "Draft"     // 草稿状态
	Published = "Published" // 发布状态
	Withdrawn = "Withdrawn" // 撤回状态
	Deleted   = "Deleted"   // 删除状态
)

type Post struct {
	ID         int64
	Title      string
	Content    string
	CreateAt   int64
	UpdatedAt  int64
	DeletedAt  int64
	Deleted    bool // 是否删除
	UserID     int64
	Status     string
	LikeNum    int64
	CollectNum int64
	ViewNum    int64
	Plate      Plate
}

type Plate struct {
	ID          int64
	Name        string
	Description string // 板块描述
	Uid         int64  // 操作人
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
	Deleted     bool
}

type Pagination struct {
	Page int    // 当前页码
	Size *int64 // 每页数据
	Uid  int64
	// 以下字段通常在服务端内部使用，不需要客户端传递
	Offset *int64 // 数据偏移量
	Total  *int64 // 总数据量
}
