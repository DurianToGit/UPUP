## 【D1提交】

### 一、基础问题回答
1. [3]int是数组，大小容量固定，值传递，[]int是切片，大小容量可动态扩容，指针传递
2. 复制的是地址值
3. Slice不等于底层数组，他的底层是由指针+len+cap组成
4. len=当前切片的数据长度，cap=当前切片的容量，是大于等于len的
5. 因为另一个切片可能是从当前切片或者裁切过去的，再没有发生动态扩容之前是共享底层数组的
6. 可能本次append超出了当前容量，发生了扩容
7. 不知道


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
&base[0]=0xc000018180
&child[0]=0xc000018180
after first append: base=[10 20] child=[10 20 30] len=3 cap=4
&base[0]=0xc000018180
&child[0]=0xc000018180
after expansion: base=[10 20] child=[10 20 30 40 50] len=5 cap=8
&base[0]=0xc000018180
&child[0]=0xc00000e500

=== independent copy ===
source=[100 20 30]
target1=[100 20 30]
target2=[10 200 30]
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

四、代码
提交代码或仓库链接：https://github.com/DurianToGit/UPUP

五、自评
独立完成比例：全部
使用AI的部分：无
最不理解的地方：无
实际学习时间：1小时