package biz

type CreateJobParams struct {
	FilePath string
}

type CreateJobResult struct {
	Jobid string
}

type ScaleParams struct {
	Jobid  string
	Width  int32
	Height int32
}

type ImageResult struct {
	ImageBytes []byte
}

type ClipParams struct {
	Jobid   string
	Width   int32
	Height  int32
	XOffset int32
	YOffset int32
}
