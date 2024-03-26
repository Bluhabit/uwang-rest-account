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
	Page       string `json:"page"`
	Size       string `json:"size"`
	TotalData  int `json:"total_data"`
	TotalPage  int `json:"total_page"`
}

func (BaseResponse[T]) Success(data T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		StatusCode: 200,
		Data:       data,
		Message:    message,
	}
}

func (BaseResponseWithPagination[T]) SuccessWithPagination(data T, message string, page string, size string, total_data int, total_page int) BaseResponseWithPagination[T] {

	return BaseResponseWithPagination[T]{
		StatusCode: 200,
		Data:       data,
		Message:    message,
		Page:       page,
		Size:       size,
		TotalData:  total_data,
		TotalPage:  total_page,

	}

}

func (BaseResponseWithPagination[T]) BadRequestPagination(data T, message string, page string, size string, total_data int,total_page int) BaseResponseWithPagination[T] {
	return BaseResponseWithPagination[T]{
		StatusCode: 400,
		Data:       data,
		Message:    message,
		Page:       page,
		Size:       size,
		TotalData:  total_data,
		TotalPage:  total_page,
	}
}

func (BaseResponse[T]) BadRequest(data T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		StatusCode: 400,
		Data:       data,
		Message:    message,
	}
}
