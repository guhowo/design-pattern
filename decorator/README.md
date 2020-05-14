# decorator design pattern

装饰器模式：在装饰器模式中，装饰器类附加的是跟原始类相关的增强功能。

## 模式解决的问题
装饰器模式解决的问题是，当一个基类有很多子类，而子类又要派生出很多子类的时候，这时候类的数量是爆炸的，也不便于维护。此时，用组合的方式来替代继承，不同的子类通过组合多个不同的基类是的类的层次清晰，也解决类的数量爆炸的问题。装饰器模式很好的展示了组合优于继承的理念。

## demo的场景介绍
有一个stream基类，声明了read,write操作。然后有多个子类：FileStream、NetworkStream和MemoryStream。此时这些子类的read方法分别要实现read后的加密功能，再写入。如果此时在用一个crypt类继承的话，就需要再写三个类，如果不仅要加密，还要增加带缓冲的功能呢？显然用继承的话会导致类的数量爆炸。

## 思考
golang所谓的用匿名函数实现继承，用interface实现运行时的多态。不知道对不对