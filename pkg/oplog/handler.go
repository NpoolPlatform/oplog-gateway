package oplog

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	constant "github.com/NpoolPlatform/oplog-middleware/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID               *uint32
	EntID            *string
	AppID            string
	UserID           *string
	Path             string
	Method           basetypes.HTTPMethod
	Arguments        string
	NewValue         *string
	Result           *basetypes.Result
	FailReason       *string
	StatusCode       *int32
	ReqHeaders       *string
	RespHeaders      *string
	ElapsedMillisecs *uint32
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithAppID(id string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppID = id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.UserID = id
		return nil
	}
}

func WithPath(path string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if !strings.HasPrefix(path, "/api/") {
			return fmt.Errorf("invalid path %v", path)
		}
		h.Path = path
		return nil
	}
}

func WithMethod(method basetypes.HTTPMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch method {
		case basetypes.HTTPMethod_GET:
		case basetypes.HTTPMethod_HEAD:
		case basetypes.HTTPMethod_POST:
		case basetypes.HTTPMethod_PUT:
		case basetypes.HTTPMethod_DELETE:
		case basetypes.HTTPMethod_CONNECT:
		case basetypes.HTTPMethod_OPTIONS:
		case basetypes.HTTPMethod_TRACE:
		case basetypes.HTTPMethod_PATCH:
		default:
			return fmt.Errorf("invalid method")
		}
		h.Method = method
		return nil
	}
}

func WithArguments(args string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		var _args map[string]interface{}
		if err := json.Unmarshal([]byte(args), &_args); err != nil {
			return err
		}
		h.Arguments = args
		return nil
	}
}

func WithNewValue(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid newvalue")
			}
			return nil
		}
		var _args map[string]interface{}
		if err := json.Unmarshal([]byte(*value), &_args); err != nil {
			return err
		}
		h.NewValue = value
		return nil
	}
}

func WithResult(result *basetypes.Result, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			if must {
				return fmt.Errorf("invalid result")
			}
			return nil
		}
		switch *result {
		case basetypes.Result_Success:
		case basetypes.Result_Fail:
		default:
			return fmt.Errorf("invalid result")
		}
		h.Result = result
		return nil
	}
}

func WithFailReason(reason *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if reason == nil {
			if must {
				return fmt.Errorf("invalid failreason")
			}
			return nil
		}
		h.FailReason = reason
		return nil
	}
}

func WithStatusCode(statusCode *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if statusCode == nil {
			if must {
				return fmt.Errorf("invalid statuscode")
			}
			return nil
		}
		h.StatusCode = statusCode
		return nil
	}
}

func WithReqHeaders(headers *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if headers == nil {
			if must {
				return fmt.Errorf("invalid reqheaders")
			}
			return nil
		}
		h.ReqHeaders = headers
		return nil
	}
}

func WithRespHeaders(headers *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if headers == nil {
			if must {
				return fmt.Errorf("invalid respheaders")
			}
			return nil
		}
		h.RespHeaders = headers
		return nil
	}
}

func WithElapsedMillisecs(secs uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ElapsedMillisecs = &secs
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
