package auth

//=========================================================================
//                               Const
//=========================================================================

const (
	ServiceName string = "auth"
)

//=========================================================================
//                               Enum
//=========================================================================

//=========================================================================
//                               Entity
//=========================================================================

//=========================================================================
//                               Mapping
//=========================================================================

//========== For Swaggo ===========

// RequestPagination Struct
type RequestPagination struct {
	Page     int `json:"page"      default:"1"  swaggertype:"integer"`
	PageSize int `json:"page_size" default:"10" swaggertype:"integer"`
}

// ResponseDataList struct
type ResponseDataList struct {
	Success            bool                `json:"success"              swaggertype:"boolean"`
	Message            string              `json:"message"              swaggertype:"string"`
	Data               interface{}         `json:"data,omitempty"       swaggertype:"object"`
	ResponsePagination *ResponsePagination `json:"pagination,omitempty" swaggertype:"object"`
	ResponseError      *ResponseError      `json:"error,omitempty"      swaggertype:"object"` // errors don't define JSON marshaling`
}

// ResponsePagination Struct
type ResponsePagination struct {
	Page       int `json:"page"         swaggertype:"integer"`
	PageSize   int `json:"page_size"    swaggertype:"integer"`
	PageCount  int `json:"page_count"   swaggertype:"integer"`
	ItemCount  int `json:"item_count"   swaggertype:"integer"`
	TotalCount int `json:"total_count"  swaggertype:"integer"`
}

// ResponseError Struct
type ResponseError struct {
	Message string `json:"message" swaggertype:"string"`
}

// ===================================

//=========================================================================
//                               Error
//=========================================================================
