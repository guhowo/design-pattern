##include<stdout>

class SplitterFactory {
public:
    virtual ISplitter* createSplitter()=0;
    virtual ~SplitterFactory(){}
};

//文件分割器
class ISplitter {
public:
    virtual void split()=0;
    virtual ~ISplitter(){}
};

// 具体的实现类：1
class BinarySplitter: public ISplitter {
public:
    virtual void split(){
        std::out<<"split binary"<<std::endl;
    }
};

// 具体的实现类：2
class TextSplitter: public ISplitter {
public:
    virtual void split(){
        std::out<<"text binary"<<std::endl;
    }
};

//具体工厂类
class BinarySplitterFactory: public SplitterFactory {
public:
    virtual ISplitter* createSplitter() {
        return new(BinarySplitter);
    }
};

class TextSplitterFactory: public SplitterFactory {
public:
    virtual ISplitter* createSplitter() {
        return new(TextSplitter);
    }
};

class MainForm {
    SplitterFactory* factory;
public:
    MainForm(SplitterFactory* factory):factory(factory){}

    void Do(){
        ISplitter* splitter = this->factory.createSplitter();
        splitter->split();
    }
};







