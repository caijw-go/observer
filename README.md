# observer
go基于观察者模式的事件调度

## Installation
```
go get -u github.com/caijw-go/observer
```

## Usage
1. 定义事件结构体
```
package event
type TestEvent struct {
    Field string
}
```

2. 定义监听器
```
package listener

import (
    events "xxx/event" //这里引入的是上面创建事件的包，需要自行修改
)

type testListener struct {
}

func (l testListener) Listen() []interface{} {
    return []interface{}{
        //如果希望在事件处理过程中修改事件数据，可以监听事件指针
        //但是事件和事件指针是两个类型，分发【事件】不会触发只监听了【事件指针】的监听器
        &events.TestEvent{},
        events.TestEvent{},
    }
}

func (l testListener) Process(event interface{}) {
    switch event.(type) {
    case *events.TestEvent:
        e := event.(*events.TestEvent)
        e.Field = "xxx" //指针类型可以修改事件属性，并重新赋值给event
        event = e
        fmt.Println(e)
    case events.TestEvent:
        e := event.(events.TestEvent)
        fmt.Println(e)
    }
}
```

3. 注册监听器
```
package listener

import "github.com/caijw-go/observer"

func Init() {//在监听器中只暴露一个Init即可，所有的监听器均可为小写
    observer.RegisterListener(testListener{})
}
```

```
package main

func main() {
    listener.Init()//程序启动时注册事件
}

```

4. 分发事件，在项目中使用以下代码进行事件分发
```
import "github.com/caijw-go/observer"

//事件：不需要修改
observer.Dispatch(event.TestEvent{})

//事件指针：会动态修改事件的值
e:= &event.TestEvent{}
observer.Dispatch(e)
fmt.Println(e)
```


