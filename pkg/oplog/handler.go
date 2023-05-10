//nolint
package oplog

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	// appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	constant "github.com/NpoolPlatform/oplog-middleware/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	EntID       *string
	AppID       string
	UserID      *string
	Path        string
	Method      basetypes.HTTPMethod
	Arguments   string
	NewValue    *string
	Result      *basetypes.Result
	FailReason  *string
	StatusCode  *int32
	ReqHeaders  *string
	RespHeaders *string
	Offset      int32
	Limit       int32
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

func WithEntID(ctx context.Context, id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(id); err != nil {
			return err
		}
		h.EntID = &id
		return nil
	}
}

func WithAppID(ctx context.Context, id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		/*
			exist, err := appmwcli.ExistApp(ctx, id)
			if err != nil {
				return err
			}
			if !exist {
				return fmt.Errorf("invalid app_id")
			}
		*/
		h.AppID = id
		return nil
	}
}

func WithUserID(ctx context.Context, id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.UserID = id
		return nil
	}
}

func WithPath(ctx context.Context, path string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if !strings.HasPrefix(path, "/api/") {
			return fmt.Errorf("invalid path %v", path)
		}
		h.Path = path
		return nil
	}
}

func WithMethod(ctx context.Context, method basetypes.HTTPMethod) func(context.Context, *Handler) error {
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

func WithArguments(ctx context.Context, args string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		var _args map[string]interface{}
		if err := json.Unmarshal([]byte(args), &_args); err != nil {
			return err
		}
		h.Arguments = args
		return nil
	}
}

func WithNewValue(ctx context.Context, value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		var _args map[string]interface{}
		if err := json.Unmarshal([]byte(*value), &_args); err != nil {
			return err
		}
		h.NewValue = value
		return nil
	}
}

func WithResult(ctx context.Context, result *basetypes.Result) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
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

func WithFailReason(ctx context.Context, reason *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.FailReason = reason
		return nil
	}
}

func WithStatusCode(ctx context.Context, statusCode *int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StatusCode = statusCode
		return nil
	}
}

func WithReqHeaders(ctx context.Context, headers *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ReqHeaders = headers
		return nil
	}
}

func WithRespHeaders(ctx context.Context, headers *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RespHeaders = headers
		return nil
	}
}

func WithOffset(ctx context.Context, offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(ctx context.Context, limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
