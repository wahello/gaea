package server

import (
	"net/http"

	"moocss.com/gaea/cmd/server/hook"
	"moocss.com/gaea/internal/service"
	"moocss.com/gaea/pkg/twirp"

	blog_v1 "moocss.com/gaea/rpc/blog/v1"
	helloworld_v1 "moocss.com/gaea/rpc/helloworld/v1"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

// var authHooks = twirp.ChainHooks(
// 	hook.NewHeaders(),
// 	hook.NewRequestID(),
// 	hook.NewLog(),
// 	hook.NewAuth(),
// )

func initMux(mux *http.ServeMux, isInternal bool) {
	{
		server := &service.HelloworldServer{}
		handler := helloworld_v1.NewHelloworldServer(server, hooks)
		mux.Handle(helloworld_v1.HelloworldPathPrefix, handler)
	}

	{
		server := &service.PostServer{}
		handler := blog_v1.NewPostServer(server, hooks)
		mux.Handle(blog_v1.PostPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux) {
}
