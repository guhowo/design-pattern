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
`type Finance interface{
    Stock() //股票
    Futures()   //期货
    Fund()  //基金
}`

如果这时候有个basic网点，此时想加一个第三方业务咋办？
最好不要试图去修改已有的接口，而是通过增加扩展去做。
`type ThirdPart interface {
    Business()
}`

只需要增加一个将原有的struct多绑定一个方法就可以了。