package main

import "github.com/DurianToGit/UPUP/stage0/d03_method_interface/experiments"

func main() {
	// 实验一：值接收者与指针接收者
	// experiments.ExperimentReceiver()
	// // 实验二：方法集与接口实现
	// experiments.ExperimentMethodSet()
	// // 实验三：Interface 的动态类型和值
	// experiments.ExperimentInterfaceValue()
	// 实验四：经典 Typed Nil 错误
	experiments.ExperimentTypedNil()
	// 实验五：类型断言
	// experiments.ExperimentTypeAssertion()
	// // 实验六：接口解耦
	// experiments.InjectNotifier()
	// experiments.ExperimentNotifyFunc()
}
