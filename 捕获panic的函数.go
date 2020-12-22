//写一个可以捕获所有goroutine panic的函数
import (
	"context"
	"runtime/debug"
	"strings"
)

func HandlePanic(ctx context.Context, f func(interface{})) func(interface{}) {
	reqID := RequestIDFromContext(ctx)
	return func(arg interface{}) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(reqID, "panic: ", err)
				stack := strings.Join(strings.Split(string(debug.Stack()), "\n")[2:], "\n")
				logger.Error(reqID, "stack: ", stack)
			}
		}()

		f(arg)
	}
}