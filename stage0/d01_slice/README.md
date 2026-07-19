## 【D1提交】

### 一、基础问题回答
1. [3]int 和 []int 有什么区别？
- go的参数都是值传递，[3]int是数组，大小容量固定，[]int是切片，大小容量可动态扩容，作为参数传递时，会复制 Slice 描述符；复制后的 Slice 通常仍然引用相同的底层数组。
2. 将 Slice 作为函数参数传递时，复制的是什么？
- 复制的是Slice的描述符
3. Slice 本身是否等于底层数组？
- Slice 描述符包含指针、len、cap；
- 指针指向底层数组；
- 底层数组真正存储元素。
4. len 和 cap 分别代表什么？
- len=当前切片的数据长度
- cap 表示从 Slice 起始位置开始，到底层数组末尾最多能够访问或扩展的元素数量。
5. 为什么修改一个 Slice，有时会影响另一个 Slice？
- 通过赋值、切片操作或在原容量范围内 append 得到的 Slice，可能共享同一个底层数组；发生扩容后，新 Slice 通常指向新的底层数组。
6. append 后，原来的 Slice 为什么有时不受影响？
- 可能本次append超出了当前容量，发生了扩容
7. 如何创建一个和原 Slice 完全独立的副本？
- 使用copy函数
  - 先用 make 创建一个长度与原切片相同的新切片。
  - 再用 copy 函数将数据复制过去
- 使用 append 函数
  - copyOfOriginal := append([]int(nil), original...)
- 使用 slices.Clone 函数 (Go 1.21+)
  - copyOfOriginal := slices.Clone(original)


### 二、实验程序
- 每段程序的实际输出；
```
=== array and slice ===
array type=[3]int len=3 value=[10 20 30]
slice type=[]int len=3 cap=3 value=[10 20 30]

=== shared underlying array ===
before: base=[10 20 30 40 50] left=[20 30 40] right=[30 40 50]
after:  base=[10 20 999 40 50] left=[20 999 40] right=[999 40 50]

=== append and capacity ===
before append: base=[10 20] len=2 cap=4 child=[10 20] len=2 cap=4
&base[0]=0xc000098040
&child[0]=0xc000098040
after first append: base=[10 20] child=[10 20 30] len=3 cap=4
&base[0]=0xc000098040
&child[0]=0xc000098040
base=[10 20]
base[:3]=[10 20 30]
after expansion: base=[10 20] child=[10 20 30 40 50] len=5 cap=8
&base[0]=0xc000098040
&child[0]=0xc0000a4080

=== independent copy ===
source=[100 20 30]
target1=[100 20 30]
target2=[10 200 30]
&a: 0xc0000ac098
&b: 0xc0000ac098
a: [100 2]
a len: 2
a cap: 3
&a: 0xc0000ac098
&b: 0xc0000ac098
b: [100 2 3]
b len: 3
b cap: 3
&a0: 0xc0000ac090
&b0: 0xc0000ae330
&a1: 0xc0000ac098
&b1: 0xc0000ae338
a: [100 2]
b: [100 200 3 4]
inside: value=[99 2 3 4] len=4 cap=6
outside: value=[99 2 3] len=3 cap=3

```
- 为什么会出现该输出；
```
=== array and slice ===
因为array创建后是固定长度的，底层也不存在容量

后面的都是因为底层数据的共享和扩容之后不共享导致的

```
- 哪些 Slice 共享底层数组；
```
=== shared underlying array ===
base和left、right是共享的
=== append and capacity ===
after first append时底层数据是共享的，之后就不共享了
=== independent copy ===
target1和source是共享的
```
- 哪一次 append 触发了扩容；
```
=== append and capacity ===
after expansion发生了扩容
```
- target1 和 target2 的区别。
```json
"target1": "target1和source是共享底层数组的",
"target2": "target2和source是不共享底层数组的"
```
### 三、预测题
第一次输出：
a:[100,2]
b:[100,2,3]
原因：
因为a,b底层数组是共享的
第二次输出：
a:[100,2]
b:[100,200,3,4]
原因：
因为a,b底层数组是共享的,但是后面b扩容了，所以b底层数组是不共享，导致第二次修改不会影响a
四、实验结论
1. 数组与Slice的区别：数组是固定长度的，Slice是可变长度的
2. Slice共享底层数组的条件：没有发生动态扩容
3. append触发扩容后的影响：改变底层数组
4. 独立复制Slice的方法：newSlice := append([]int(nil), source...)

### 四、代码
提交代码或仓库链接：https://github.com/DurianToGit/UPUP

### 五、自评
独立完成比例：全部
使用AI的部分：无
最不理解的地方：无
实际学习时间：1小时

### 六、补充实验
#### Slice 参数传递实验
1. 为什么 source[0] 变成了 99？
- 因为source[0]和s[0]共享了同一个底层数组
2. 为什么调用方的 len(source) 仍然是3？
- 因为modifySlice中 append发生了扩容，复制到一个新的地址空间，不再共享同一个底层数组，所以不影响调用方的len(source)
3. 为什么函数内部的 append 没有让调用方长度变成4？
- 因为函数内部的append让s复制到一个新的地址空间，不再使用同一个底层数组了

#### 长度与底层数组实验、
在原实验第一次 append 后增加：

```
fmt.Printf("base=%v\n", base)
fmt.Printf("base[:3]=%v\n", base[:3])
```

解释为什么两个输出不同。
:因为原本的base的len长度只有2，所以显示的时候只能按照2的长度显示[10, 20],但是因为容量是4，所以在第一次append后再没有扩容，又共享了底层数组的情况下，第三个元素是已经被影响了的