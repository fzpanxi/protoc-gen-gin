package main

var httpTemplate = `
type {{.ServiceName}}HttpServer interface {
{{range .MethodSets}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{end}}
}

type {{$.ServiceName}} struct{
	server {{ $.ServiceName }}HttpServer
	router *gin.Engine
}

func Register{{ $.ServiceName }}HttpServer (r *gin.Engine, srv {{ $.ServiceName }}HttpServer) {
	s := {{.ServiceName}} {
		server: srv,
		router: r,
	}
	s.RegisterService()
}

{{range .Methods}}
func (s *{{$.ServiceName}}) {{ .HandlerName }} (ctx *gin.Context) {
	var in {{.Request}}
{{if .HasRouterParams }}
	if err := binding.MapProto(&in, binding.RestParams(ctx)); err != nil {
		binding.Response(ctx, nil, err)
		return
	}
{{end}}
{{if eq .Method "GET" "DELETE" }}
	if err := binding.MapProto(&in, binding.QueryParams(ctx)); err != nil {
		binding.Response(ctx, nil, err)
		return
	}
{{end}}
{{if eq .Method "POST" "PUT" }}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		binding.Response(ctx, nil, err)
		return
	}
{{end}}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.{{.Name}}(newCtx, &in)
	binding.Response(ctx, out, err)
}
{{end}}

func (s *{{$.ServiceName}}) RegisterService() {
{{range .Methods}}
		s.router.Handle("{{.Method}}", "{{.Path}}", s.{{ .HandlerName }})
{{end}}
}
`
