package model

type FileCheckReq struct {
	RequestId string `json:"request_id"`
	Path      string `json:"path"`
}

type Md5JsonFile struct {
	Folder string            `json:"folder"`
	Files  []Md5JsonSubFiles `json:"files"`
}

type Md5JsonSubFiles struct {
	Name string `json:"name"`
	Md5  string `json:"md5"`
}
