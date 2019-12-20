# Design-pattern
开闭原则
对扩展开放，对修改关闭。
意思就是通过增加代码来扩展功能，而不是通过修改来扩展功能。
感觉和S（单一原则）是结合的。
比如一个公司有三个业务，那么三个业务。都放在一个interface里可以的
比如：
type Basic interface{
    Save()  //  存款业务
    Withdraw()  //取款业务
    Transfer()  //转账业务
}

如果随着公司扩大，有了股票、理财业务等大量业务。如果把所有方法定义在一个接口，那是很危险的。业务的耦合度太高。
应该拆分成若干个interface，然后组合起来。
type Finance interface{
    Stock() //股票
    Futures()   //期货
    Fund()  //基金
}

然后组合起来
type Banker interface{
    Basic
    Finance
}

好处：
这样做的好处是，如果这公司扩张，开分店，有了很多实例，有些实例只有基本业务，有些则只做金融业务，有些全都做。隔离型就做的很好。小网点也不用为了实现接口，去实现所有金融加上基本业务。