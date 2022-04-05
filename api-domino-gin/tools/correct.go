package tools

import "sync"

type correct struct {
	m  string
	c  interface{}
	s  uint
	ec uint
}

var responseIns *correct
var once sync.Once

func CorrectIns(msg string) *correct {
	once.Do(func() { responseIns = &correct{m: msg} })
	return responseIns
}

func (cls *correct) Get() map[string]interface{} {
	ret := map[string]interface{}{
		"message":    cls.m,
		"content":    cls.c,
		"status":     cls.s,
		"error_code": cls.ec,
	}
	return ret
}

func (cls *correct) Set(content interface{}, status uint, errorCode uint) *correct {
	cls.c = content
	if status == 0 {
		cls.s = 200
	} else {
		cls.s = status
	}
	cls.ec = errorCode
	return cls
}

func (cls *correct) Ok(content interface{}) (int, map[string]interface{}) {
	if cls.m == "" {
		cls.m = "OK"
	}
	return 200, cls.Set(content, 200, 0).Get()
}

func (cls *correct) Created(content interface{}) (int, map[string]interface{}) {
	if cls.m == "" {
		cls.m = "新建成功"
	}
	return 201, cls.Set(content, 201, 0).Get()
}

func (cls *correct) Updated(content interface{}) (int, map[string]interface{}) {
	if cls.m == "" {
		cls.m = "编辑成功"
	}

	return 202, cls.Set(content, 202, 0).Get()
}

func (cls *correct) Deleted() (int, interface{}) {
	if cls.m == "" {
		cls.m = "删除成功"
	}
	return 204, cls.Set(nil, 204, 0).Get()
}
