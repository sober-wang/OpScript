# 策略模式

## 意图
定义一系列的算法，把他们一个个封装起来，并是他们可以互相替换。本模式是的算法可独立于使用他的客户而变化

## 动机
许多算法可对一个文本流进行分行。将这些算法硬编近使用他们的类中是不可取的。原因如下：
- 需要换行功能的客户程序如果直接包含换行算法代码的话会变得复杂，这使得客户程序庞大且难以维护，尤其当其不需要支持多种换行算法时问题会更加严重
- 不同的时候需要使用不同的算法，我们不想支持我们并不适用得换行算法
- 当换行功能是客户程序得一个难以分割得成分时，增加新的换行算法或改变现有算法将十分困难

我们可以定义一些类来封装不同得换行算法，从而避免这些问题。这种封装方法叫策略（Strategy）

## 适用性
以下情况使用策略模式：
- 许多相关得类仅仅是行为有异，策略提供了一种多个行为中得一个行为来配置一个类得方法
- 需要使用一个算法得不同变体。例如：二叉树的遍历分为，前，中，后序遍历方法；排序算法分，冒泡，快速，堆排序
- 算法使用客户不应该知道的数据。可使用策略模式以避免暴漏复杂的，与算法有关的数据结构
- 一个类定义了多种行为，并且这些行为在这个类的操作中以多个条件语句的形式出现。将相关的条件分支移入他们各自的类中以替代这些条件语句

## 参与者
- Strategy： 定义所有支持算法的公共接口。Context 使用这个接口来调用 ConcreteStrategy 定义的算法
- ConcreteStrategy: 具体策略，例如：遍历查找，二分查找。
- Context 上下文
    - 用一个 ConcreteStrategy 对象来配置
    - 维护一个对 Strategy 对象的引用
    - 可顶一个一个接口来让 Strategy 访问它的数据

## 协作
- Strategy 和 Context 湘湖作用以实现选定的算法。
    - 当算法被调用时，Context 可以将该算法所需要的所有数据都传递给该 Strategy 
    - Context可以将自己作为一个参数传递给 Strategy 操作。这就让 Stragey 需要时可以回调 Context 
- Context 将客户的请求转发给它的 Strategy。
    - 客户通常创建一个 ConcreteStrategy 对象给 Context ，这样客户经与 Context 交互
    - 通常有一系列的 ConcreteStrategy 类可共客户从中选择

## 效果
### 优点
1. 相关算法系列： Strategy 类层次为 Context 定义了一系列可供复用的算法或行为。
1. 一个代替继承的方法：继承虽然可以实现类似的功能，但在类的多维度继承中难免针对父类方法进行修改，这使得算法的实现与 Conntext 混合在一起，变得难以维护，难以拓展。算法独立封装在 Strategy 中使我们可以独立于 Context 而改变它，它易于切换，易于扩展，易于理解
1. 消除一些条件语句：当多种算法堆砌在一个类中，难免要使用条件语句选择，而基于不同的 Context , Strategy 可以自动选择算法。例如：查找算法中集合大小作为 Context ，就可以自动选择不同的算法
1. 实现的选择：策略模式提供相同行为的不同实现。

### 缺点
1. 客户必须了解不同的 Strategy：客户要选择一个适合的就要知道这些 Strategy 到底有何不同，此时就不得不暴漏算法和相关数据结构。因此仅当这些不同行为的变体与客户相关时，才需要使用用策略模式
1. Strategy 和 Context 之间的通信开销： 无论各个具体的算法实现是简单还是复杂，它们都共享 Strategy 定义的接口。因此有一些简单算法可能接收到一些用不到的参数，这会导致内存浪费，Context 可能创建和初始化一些永远用不到的参数，这就需要 Context 和 Strategy 更紧耦合
1. 增加了对象数目： Strategy 增加了一个应用中的对象的数目。例如一个排序算法，如果你的数据规模是一尘不变的，Strategy 管理的排序算法可能永远无法实现。

## 实现
Golang 
```golang
package main

import "fmt"

// Calculate 名为计算的结构题。策略模式中的 Context
type Calculate struct{}

func (cl *Calculate) Plus() {
    fmt.Println("加法")
}

func (cl *Calculate) Subtration() {
    fmt.Println("减法")
}

func (cl *Calculate) CalculusEquation(ce Equation) {
    ce.Answer()

}

// Equation 方程接口，里面有个 Answer() 解答方法；策略模式中的：Stategy
type Equation interface {
    Answer()
}

type Differential struct{}

// Answer() 基于微分方程的具体实现。 策略模式中 ConcreteStrategy
func (*Differential) Answer() {
    fmt.Println("微分方程")
}

type Integral struct{}

func (*Integral) Answer() {
    fmt.Println("积分方程")
}

func main() {
    hard := &Calculate{}
    hard.Plus()
    hard.Subtration()
    hard.CalculusEquation(new(Differential))
    hard.CalculusEquation(new(Integral))

}

```
