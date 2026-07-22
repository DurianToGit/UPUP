# D3：方法、Interface、指针接收者与值接收者
## 一、基础问题

### 1. 方法和普通函数有什么区别；
回答：
- 普通函数独立存在，方法绑定在特定的类型（接收者）上。
- 方法调用时，会自动将接收者作为第一个参数传入，而函数直接通过函数名调用。
- 函数无法被接口约束，不能用于接口实现；方法可以被接口约束，用于接口实现。

### 2. 值接收者方法中修改字段，是否会改变调用者？
回答：不会。
   因为值接收者在调用时，Go 会拷贝一份完整的接收者副本传入方法。方法内部修改的是副本的字段，原始变量不受影响。

### 3. 指针接收者方法中修改字段，是否会改变调用者？
回答：会。
因为指针接收者传递的是变量的内存地址，方法内部通过解引用操作修改的是同一块内存数据。
### 4. 为什么下面代码能够调用成功？
```go
type User struct {
    Name string
}

func (u *User) Rename(name string) {
    u.Name = name
}

func main() {
    user := User{Name: "Tom"}
    user.Rename("Jerry")
}
```
回答：这是 Go 提供的语法糖。编译器在编译时，自动将 user.Rename("Jerry") 隐式转换为 (&user).Rename("Jerry")

### 5. 如果一个接口要求方法 Save()，而类型只定义了： `func (u *User) Save()`, 那么`User`和`*User`谁实现了接口？
回答：只有 *User 实现了该接口。
因为 Go 的方法集规则：指针类型的方法集包含所有指针接收者和值接收者的方法；而值类型的方法集仅包含值接收者的方法。这里的接收者是指针，所以 User 类型缺失该方法。

### 6. 如果类型定义的是：`func (u User) Save()`, 那么User和*User谁实现了接口？
回答：User 和 *User 都实现了。
原因同上：*User 的方法集包含值接收者的方法（Go 允许自动取地址或解引用，指针可以调用值方法）。所以无论接口变量接收的是 User 还是 *User，都能满足接口。

### 7. Interface变量中通常保存什么信息？
回答：接口变量本质上是一个结构体，保存着动态类型和动态值

### 8. nil Interface和“内部保存nil指针的Interface”是否相同？
回答: 不相同
- nil 接口：接口的 类型 和 值 都为 nil，判断 iface == nil 为 true
- 内部存 nil 指针的接口：接口的 类型 不为 nil（比如类型是 *int），只是 值 为 nil。此时判断 iface == nil 为 false。
### 9. 类型断言失败会发生什么？
回答：取决于你使用单返回值还是双返回值形式：

- 单返回值（v := i.(T)）：如果断言失败，会直接触发 panic，程序崩溃。

- 双返回值（v, ok := i.(T)）：不会 panic，断言失败时 ok 为 false，v 为类型 T 的零值。

### 10. 类型断言的comma ok形式解决什么问题？
回答：解决 “避免 panic” 和 “安全判断” 的问题

## 二、代码实验
### 实验一：值接收者与指针接收者
问题：
- 为什么值接收者没有修改原对象；
- 为什么指针接收者修改了原对象；
- 调用 user.RenameByPointer() 时，编译器大致做了什么；
- User 不可寻址时，是否还能自动取地址？
回答：
- 因为值接收者在调用时，Go 会拷贝一份完整的接收者副本传入方法。方法内部修改的是副本的字段，原始变量不受影响。
- 因为指针接收者传递的是变量的内存地址，方法内部通过解引用操作修改的是同一块内存数据。
- 当编译器发现 RenameByPointer 的接收者是 *User，而调用者是类型为 User 的 user 变量时，它会自动插入取地址符，将代码： `user.RenameByPointer("Pointer Jerry")` 等价改写为：`(&user).RenameByPointer("Pointer Jerry")`
- 不能，编译器会直接报错；Go 的隐式取地址语法糖仅适用于“可寻址（addressable）”的操作数

### 实验二：方法集与接口实现
- 错误（`printSaved(document)`）记录：stage0\d03_method_interface\experiments\2.go:24:13: cannot use document (variable of type Document) as Saver value in argument to printSaved: Document does not implement Saver (method Save has pointer receiver)
- 恢复后正常执行（`printSaved(&document)`）结果：saved: report
- 增加类型后运行结果：
  - saved image:photo.png
  - saved image:photo.png
- 总结：

| 类型定义的方法            | `T` 是否实现接口 | `*T` 是否实现接口 |
| ---------------------- | ---------- | ----------- |
| `func (t T) Method()`  | 是          | 是           |
| `func (t *T) Method()` | 否          | 是           |

### 实验三：Interface 的动态类型和值
只有动态类型和动态值都为空时，Interface 才等于 nil。
- 在 ExperimentInterfaceValue 函数内进行类型判断 `fmt.Println(userPointer2 == nil) 是 true`，因为不产生动态;因为 userPointer2 的静态类型是 *User，它是一个具体的指针变量。
  此时，编译器只检查这个变量存储的内存地址值。它没有被赋予任何有效地址，存储的就是 0x0（零值），所以比较结果为 true
- 为什么 inspect(userPointer2) 中 nil=false
  - 因为当执行 value any = userPointer2 时，发生了一次隐式的接口装箱（Boxing），在go运行时，接口变量 value 实际上就是一个动态类型和动态值组成的结构体了，这时候动态类型为 *User，动态值为nil，所以整体非nil

### 实验四：经典 Typed Nil 错误
#### 该问题在线上可能造成什么影响。
回答：会造成 err == nil时无法正确判断，导致会以为报错了（因为 err != nil），但实际读取错误信息时却读到空指针，引发 panic。

### 实验五：类型断言
不带 ok 的断言失败时会引发 panic。

### 实验六：接口解耦
#### 1. OrderService 为什么不需要知道具体通知方式？
回答：这是依赖倒置原则（DIP，Dependency Inversion Principle）的体现。
#### 2. 接口定义在调用方还是实现方附近更合理？
回答：绝对应该定义在“调用方”（使用方）附近，而不是实现方附近。
#### 3. 接口是不是越大越好？
回答：绝对不是，接口应该是越小越好（go倡导“小接口”）。
#### 4. 这种设计如何方便测试？
回答：这种设计最大的红利之一就是单元测试极其轻量且稳定。

由于 OrderService 依赖的是接口，在测试时，我们不需要真的去连接邮件服务器或短信网关，只需在测试文件中定义一个“模拟（Mock）实现”。

## 四、结论
### 1. 方法与普通函数有什么区别？
- 归属不同：函数是独立的代码块；方法必须绑定到特定的类型（receiver）。
- 调用方式不同：函数通过 函数名(参数) 调用；方法通过 实例.方法名(参数) 调用。

### 2. 值接收者和指针接收者分别复制什么？
- 值接收者：复制的是接收者变量的副本，方法内部修改不影响原对象
- 指针接收者：复制的是接收者变量的指针地址，方法内部修改直接影响原对象

### 3. 什么是方法集？
个类型所拥有的、可以被调用的全部方法的集合。
它决定了该类型能赋给什么样的接口变量，以及通过 interface 调用哪些方法

### 4. T和*T的方法集有什么区别？
- T的方法集：仅包含T的所有方法。
- *T的方法集：包含*T和T的所有方法。

### 5. 如何判断一个类型是否实现接口？
常用编译期检查写法（推荐）：
> var _ Interface = (*MyStruct)(nil) // 确保 *MyStruct 实现了 Interface

如果未实现，编译器会报错：missing method XXX。
### 6. Interface为什么可能不等于nil？
- 因为要满足动态类型和动态值都为nil时，Interface才等于nil。

### 7. Typed Nil有什么实际风险？
- 错误判断失效：像上面那样，if err != nil 明明非空却进入错误分支，导致程序逻辑混乱或 panic。

- 空指针解引用：虽然 err 接口变量不为 nil，但其内部的动态值（指针）是 nil，如果调用该指针的方法，会引发 panic。

### 8. 类型断言失败会怎样？
回答：取决于你使用单返回值还是双返回值形式：

- 单返回值（v := i.(T)）：如果断言失败，会直接触发 panic，程序崩溃。

- 双返回值（v, ok := i.(T)）：不会 panic，断言失败时 ok 为 false，v 为类型 T 的零值。
### 9. Type Switch解决什么问题？
解决多类型判断的场景。当你有一个 interface{} 或大接口，需要根据其具体类型执行不同业务逻辑时，用 type switch 比一堆 if-else + 断言更清晰、安全。

### 10. Go为什么提倡小接口？
- 解耦：只需定义我需要的行为（如 io.Reader 只有一个方法）。

- 组合优于继承：大接口可由多个小接口组合而成。

- 便于 Mock 和测试：小接口只需要实现 1~2 个方法，Mock 代码量极少，极大降低单元测试成本。

### 11. 接口应该由调用方还是实现方定义？
必须由调用方（消费者）定义。
这是 Go 与 Java 等语言最大的不同。实现方只负责提供具体结构体，而接口定义在使用该接口的函数/模块旁边。标准库典范：io.Reader 定义在调用它数据的包中，而非实现数据读取的包中。

### 12. 如何通过接口进行单元测试隔离？

利用依赖注入（DI） + 小接口替换真实依赖。

步骤：

- 定义小接口（只包含业务需要的方法），比如 type UserGetter interface { Get(id int) error }。

- 生产代码注入真实实现（如 *UserRepo）。

- 测试代码注入 Mock 实现（如下方示例）。

测试代码示例：

```go
type mockUserGetter struct { fail bool }

func (m *mockUserGetter) Get(id int) error {
    if m.fail { return errors.New("mock err") }
    return nil
}

func TestBusiness(t *testing.T) {
    m := &mockUserGetter{fail: true}
    err := DoBusiness(m) // 注入 mock
    if err == nil { t.Error("期望错误但未返回") }
}
```
> 这种模式下，测试不再依赖数据库或网络，执行速度极快（配合 go test -short 可跳过集成测试）。

## 五、自评
独立完成比例：≈10%
AI参与部分：≈90%
最不理解的地方：T和*T的方法集有什么区别？
实际学习时间：> 3小时