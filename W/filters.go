package W

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"time"
)

func PanicFilter(ctx *Context) {
	defer func() {
		if err := recover(); err != nil {
			//L.Panic(errors.New(`Internal Server Error`), ``, err)
			err2, ok := err.(error)
			if !ok {
				err2 = errors.New(fmt.Sprintf("%# v", err))
			}
			err_str := err2.Error()
			L.LOG.Errorf(err_str)
			str := L.StackTrace(2)
			L.LOG.Criticalf("StackTrace: %s", str)
			ctx.Title += ` (error)`
			detail := ``
			if ctx.Engine.DebugMode {
				detail = spew.Sdump(err) + string(L.StackTrace(3))
			} else {
				ref_code := `EREF:` + S.RandomCB63(4)
				L.LOG.Notice(ref_code) // no need to print stack trace, should be handled by PanicFilter
				detail = `Detailed error message not available on production mode, error reference code for webmaster: ` + ref_code
				ctx.Engine.SendDebugMail(ctx.RequestDebugStr() + S.WebBR + S.WebBR + err_str + S.WebBR + detail)
			}
			ctx.Error(500, detail)
		}
	}()
	ctx.Next()(ctx)
}

func LogFilter(ctx *Context) {
	start := time.Now()
	ctx.Next()(ctx)
	L.Trace()
	var codeStr string
	code := ctx.Response.StatusCode()
	if code < 400 {
		codeStr = L.BgGreen(`%s`, color.BlueString(`%3d`, code))
	} else {
		codeStr = L.BgRed(`%3d`, code)
	}
	ones := `nano`
	elapsed := float64(time.Since(start).Nanoseconds())
	if elapsed > 1000000000.0 {
		elapsed /= 1000000000.0
		ones = `sec`
	} else if elapsed > 1000000.0 {
		elapsed /= 1000000.0
		ones = `mili`
	} else if elapsed > 1000.0 {
		elapsed /= 1000.0
		ones = `micro`
	}
	referrer := ctx.RequestCtx.Referer()
	url := ctx.RequestCtx.RequestURI()
	msg := fmt.Sprintf(`[%s] %7d %7.2f %5s | %4s %-40s | %-40s | %15s | %s || %s`,
		codeStr,
		ctx.Buffer.Len(),
		elapsed,
		ones,
		ctx.RequestCtx.Method(),
		url,
		referrer,
		ctx.Session.IpAddr,
		ctx.Session.String(),
		ctx.Posts().String(),
	)
	msg = S.Replace(msg, `%`, `%%`)
	L.LOG.Notice(msg)
}

func SessionFilter(ctx *Context) {
	ctx.Session = &Session{}
	ctx.Session.Load(ctx)
	ctx.Next()(ctx)
	ctx.Session.Save(ctx)
}