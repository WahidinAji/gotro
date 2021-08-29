package main

import (
	"context"

	"os"

	"github.com/kokizzu/gotro/W2/example/conf"
	"github.com/kokizzu/gotro/W2/example/domain"
)

func cliArgsRunner(args []string) {
	tracerCtx, span := conf.T.Start(context.Background(), args[0])
	defer span.End()

	var (
		vdomain = domain.NewDomain()
	)

	patterns := map[string]map[string]int{
		domain.UserChangeEmail_Url:    {},
		domain.UserChangePassword_Url: {},
		domain.UserConfirmEmail_Url:   {},
		domain.UserForgotPassword_Url: {},
		domain.UserList_Url:           {},
		domain.UserLogin_Url:          {},
		domain.UserLogout_Url:         {},
		domain.UserProfile_Url:        {},
		domain.UserRegister_Url:       {},
		domain.UserResetPassword_Url:  {},
	}
	switch pattern := cliUrlPattern(args[0], patterns); pattern {

	case domain.UserChangeEmail_Url:
		in := domain.UserChangeEmail_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserChangeEmail(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserChangePassword_Url:
		in := domain.UserChangePassword_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserChangePassword(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserConfirmEmail_Url:
		in := domain.UserConfirmEmail_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserConfirmEmail(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserForgotPassword_Url:
		in := domain.UserForgotPassword_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserForgotPassword(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserList_Url:
		in := domain.UserList_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserList(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserLogin_Url:
		in := domain.UserLogin_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserLogin(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserLogout_Url:
		in := domain.UserLogout_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserLogout(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserProfile_Url:
		in := domain.UserProfile_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserProfile(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserRegister_Url:
		in := domain.UserRegister_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserRegister(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	case domain.UserResetPassword_Url:
		in := domain.UserResetPassword_In{}
		in.FromCli(os.Stdin, tracerCtx)
		out := vdomain.UserResetPassword(&in)
		out.ToCli(os.Stdout)
		in.ToCli(os.Stdout, &out)

	}
}