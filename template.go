package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type serviceDesc struct {
	ServiceName string
	FullName    string
	FilePath    string
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}
type methodDesc struct {
	// method
	Name    string
	Num     int
	Vars    []string
	Forms   []string
	Request string
	Reply   string
	// http_rule
	Path         string
	Method       string
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute(pbFilename string) string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	bufStr := string(buf.Bytes())
	return bufStr
}

// HandlerName for gin handler name
func (m *methodDesc) HandlerName() string {
	return fmt.Sprintf("%s_%d", m.Name, m.Num)
}

// HasRouterParams
func (m *methodDesc) HasRouterParams() bool {
	paths := strings.Split(m.Path, "/")
	for _, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			return true
		}
	}
	return false
}

// replaceRouterParams replace {xxxx} to :xxxx
func (m *methodDesc) replaceRouterParams() {
	paths := strings.Split(m.Path, "/")
	for i, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			paths[i] = ":" + p[1:len(p)-1]
		}
	}
	m.Path = strings.Join(paths, "/")
}
