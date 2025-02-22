# Go 格式化占位符说明

- `%v`：默认格式打印值  
  例：fmt.Printf("Value: %v", value)

- `%+v`：打印结构体时带字段名  
  例：fmt.Printf("Struct: %+v", myStruct)

- `%#v`：输出值的 Go 语法表示  
  例：fmt.Printf("Go syntax: %#v", value)

- `%T`：打印值的类型  
  例：fmt.Printf("Type: %T", value)

- `%s`：打印字符串  
  例：fmt.Printf("String: %s", str)

- `%d`：打印整数（十进制）  
  例：fmt.Printf("Decimal: %d", num)

- `%f`：打印浮点数  
  例：fmt.Printf("Float: %f", f)

更多占位符请参考官方文档 [fmt包](https://pkg.go.dev/fmt)
