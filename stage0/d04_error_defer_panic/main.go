package main

func main() {
	// 实验一：Error 包装与 errors.Is
	// experimentErrorIs()
	// 实验二：自定义 Error 与 errors.As
	// experimentErrorAs()
	// 实验三：Defer 顺序与求值时机
	// experimentDeferEvaluation()
	// 实验四：Defer 修改命名返回值
	// ExperimentNamedReturn()
	// 实验五：循环中的 Defer
	ids := []int{1, 2, 3}
	// processResourcesBad(ids)
	processResourcesGood(ids)
}
