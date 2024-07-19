package domain

const (
	Draft     = "Draft"     // 草稿状态
	Published = "Published" // 发布状态
	Withdrawn = "Withdrawn" // 撤回状态
	Deleted   = "Deleted"   // 删除状态
)

type Check struct {
	ID        int64  // 审核ID
	PostID    int64  // 帖子ID
	Content   string // 审核内容
	Title     string // 审核标签
	UserId    int64  // 提交审核的用户ID
	Status    string // 审核状态
	Remark    string // 审核备注
	CreatedAt int64  // 创建时间
	UpdatedAt int64  // 更新时间
}

type Pagination struct {
	Page int    // 当前页码
	Size *int64 // 每页数据
	Uid  int64
	// 以下字段通常在服务端内部使用，不需要客户端传递
	Offset *int64 // 数据偏移量
	Total  *int64 // 总数据量
}
