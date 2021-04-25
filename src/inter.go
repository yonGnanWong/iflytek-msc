package src

type Msc interface {
	/**
		预处理.包括文件名,应用id等参数上传
	 */
	Prepare()
	/**
		分片上传文件
	 */
	Upload()
	/**
		服务端文件合并
	 */
	Merge()
	/**
		获取处理进程
	 */
	GetProgress()
	/**
		获取处理结果
	 */
	GetResult()
}

func NewMsc() Msc {
	return newClient()
}
