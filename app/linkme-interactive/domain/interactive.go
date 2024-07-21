package domain

type Interactive struct {
	Id           int64
	BizId        int64
	BizName      string
	ReadCount    int64
	LikeCount    int64
	CollectCount int64
	UpdateTime   int64
	CreateTime   int64
	PostId       int64
	DeletedAt    int64
}

type Pagination struct {
	Page int    // 当前页码
	Size *int64 // 每页数据
	Uid  int64
	// 以下字段通常在服务端内部使用，不需要客户端传递
	Offset *int64 // 数据偏移量
	Total  *int64 // 总数据量
}
