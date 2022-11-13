- 校验controller接口是否是ControllerActivator实体
```
controller interface{}
if c, ok := controller.(*ControllerActivator); ok {
		return c
	}
```


- 判断类是否实现了BaseController接口。
```
//controller 为实体类
typ := reflect.TypeOf(controller)

type BaseController interface {
	BeginRequest(*context.Context)
	EndRequest(*context.Context)
}
var baseControllerTyp = reflect.TypeOf((*BaseController)(nil)).Elem()

func isBaseController(ctrlTyp reflect.Type) bool {
	return ctrlTyp.Implements(baseControllerTyp)
}
```

- 校验类是否实现了特定接口

```
if before, ok := controller.(interface {
		BeforeActivation(BeforeActivation)
	}); ok {
		before.BeforeActivation(c)
	}
```

- 校验是否为v是否为reflect.Value

```
if val, ok := v.(reflect.Value); ok {
		// check if it's already a reflect.Value.
		return val
	}
```



- 获取类的全名称, 结果为包名+类名

```
func NameOf(v interface{}) string {
	elemTyp := indirectType(reflect.ValueOf(v).Type())

	typName := elemTyp.Name()
	pkgPath := elemTyp.PkgPath()
	fullname := pkgPath[strings.LastIndexByte(pkgPath, '/')+1:] + "." + typName

	return fullname
}

func indirectType(typ reflect.Type) reflect.Type {
	switch typ.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return typ.Elem()
	}
	return typ
}
```

- 解析类方法

```
// Type  reflect.Type
func (c *ControllerActivator) parseMethods() {
	n := c.Type.NumMethod()
	for i := 0; i < n; i++ {
		m := c.Type.Method(i)
		c.parseMethod(m)
	}
}
```

- 反射调用方法

```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := new(Test)
	of := reflect.TypeOf(t)
	method, ok := of.MethodByName("A")
	if ok {
		inputs := make([]reflect.Value, method.Type.NumIn())
		//inputs 第一个参数为，对象本身.
		inputs[0] = reflect.ValueOf(t)
		method.Func.Call(inputs)
	}
}

type Test struct {
}

func (t *Test) A() {
	fmt.Print("AAAAA")
}
```

