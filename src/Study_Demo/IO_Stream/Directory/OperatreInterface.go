package Directory

//目录操作接口

// DirectoryOperate 目录操作接口
type DirectoryOperate interface {

	// MakeDirectory 创建单个目录
	MakeDirectory(path string)

	//MakeAllDirectory 创建层级目录
	MakeAllDirectory(path string)

	//RemoveDirectory 删除单个目录
	RemoveDirectory(path string)

	//RemoveAllDirectory 删除层级目录
	RemoveAllDirectory(path string)
}

// FileOperate 文件操作接口
type FileOperate interface {
	FileCreate(path string)
	FileDelete(path string)
}
