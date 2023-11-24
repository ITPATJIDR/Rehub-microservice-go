package fileModel

type FileSturct struct {
	Filepath string `json:"Filepath" binding:"required"`
}

type UploadFileSturct struct {
	Foldername        string `json:"Foldername" binding:"required"`
	Physicalrawreport string `json:"Physicalrawreport" binding:"required"`
}
