package models

type BaseResponse[T any] struct {
	StatusCode uint16 `json:"status_code"`
	Data       T      `json:"data"`
	Message    string `json:"message"`
}
type BaseResponseWithPagination[T any] struct {
	StatusCode uint16 `json:"status_code"`
	Data       T      `json:"data"`
	Message    string `json:"message"`
	Page       uint32 `json:"page"`
	Size       uint32 `json:"size"`
	TotalData  uint32 `json:"total_data"`
}

func (BaseResponse[T]) Success(data T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		StatusCode: 200,
		Data:       data,
		Message:    message,
	}
}

func (BaseResponseWithPagination[T]) SuccessWithPagination(data T, message string, page uint32, size uint32, total_data uint32) BaseResponseWithPagination[T] {
	return BaseResponseWithPagination[T]{
		StatusCode: 200,
		Data:       data,
		Message:    message,
		Page:       page,
		Size:       size,
		TotalData:  total_data,
	}
}

func (BaseResponse[T]) BadRequest(data T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		StatusCode: 400,
		Data:       data,
		Message:    message,
	}
}
