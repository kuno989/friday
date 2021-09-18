package rabbitmq

type ResponseObject struct {
	MinioObjectKey string `json:"minio_object_key"`
	Sha256         string `json:"sha256"`
}
