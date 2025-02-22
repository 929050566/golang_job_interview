package main

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}

// # 问题分析：map 中存储循环变量的地址问题

// 在函数 pase_student 中：
// - `for _, stu := range stus` 中的变量 `stu` 是在整个循环中复用的。
// - `&stu` 总是获取相同的地址，导致所有 map 的值都指向最后一个学生。

// 解决方法：
// - 使用索引方式遍历切片，然后取该元素的地址，如：
//   ```go
//   for i := range stus {
//       m[stus[i].Name] = &stus[i]
//   }
//   ```
