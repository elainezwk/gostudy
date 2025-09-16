package main

import (
	"fmt"
	deflog "study/logprint"
	"time"
)

// //=================结构体嵌套
// 当数据类型比较复杂时，可以设置为多个结构体，之后可以再某个结构体的一个成员变量中去声明另外一种类型的变量
// 多个结构体有相互依赖关系时，可以使用结构体嵌套。
// 如果嵌套时，成员变量有名称重复时，就不能直接赋值，要用全路径赋值 例如 info1.Phone.Name="小米"
type Phone struct {
	Name  string
	Mode  string  // 型号
	Price float32 // 价格
}
type Person struct {
	Name    string
	Age     int
	Address string
	Gender  string
	Mobile  Phone
}

// 组合的嵌套结构体
type Info struct {
	Person
	Phone
}

// //=================接口
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

/*
	type 接口名称 interface {
		方法名称(传参) 返回值
		方法名称(传参) 返回值
	}
*/
type DBCommon interface {
	Insert(string) error
	Update(string) error
	Delete(string) error
}

// 定义一个类型去实现这个接口
type MySQL struct {
	config  DBConfig
	charSet string
}

func (m MySQL) Insert(data string) error {
	fmt.Println("插入数据到MySQL:", data)
	return nil
}
func (m MySQL) Update(data string) error {
	fmt.Println("更新MySQL数据:", data)
	return nil
}
func (m MySQL) Delete(data string) error {
	fmt.Println("删除MySQL数据:", data)
	return nil
}

type PostGreSQL struct {
	config  DBConfig
	charSet string
}

func (m PostGreSQL) Insert(data string) error {
	fmt.Println("插入数据到PostGreSQL:", data)
	return nil
}
func (m PostGreSQL) Update(data string) error {
	fmt.Println("更新PostGreSQL:", data)
	return nil
}
func (m PostGreSQL) Delete(data string) error {
	fmt.Println("删除PostGreSQL:", data)
	return nil
}

type EmptyInterface interface{}

func dealData(data interface{}) {
	t, ok := data.(string)
	if ok {
		fmt.Println("当前类型是string,变量的值是:", data)
	} else {
		fmt.Println("data不是string")
		fmt.Println("当前t的值是:", t)
	}
}

func getType(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("当前值为int类型，值：", t)
	case string:
		fmt.Println("当前值为string类型，值：", t)
	case bool:
		fmt.Println("当前值为bool类型，值：", t)
	default:
		fmt.Println("当前值类型不在处理范围内，值类型为：", t)
	}
}

// //=================协程
func makeBuns(filling string, buns chan string) {
	//startTime := time.Now()
	//fmt.Printf("%s馅开始的时间:%s\n", filling, startTime)
	fmt.Printf("开始做%s馅的包子\n", filling)
	fmt.Printf("开始剁%s馅的包子\n", filling)
	fmt.Println("开始擀皮儿……")
	//time.Sleep(time.Second) // 停1秒钟
	fmt.Printf("开始包%s馅的包子\n", filling)
	fmt.Printf("开始蒸%s馅的包子\n", filling)
	//cost := time.Since(startTime)
	// 使用协程后，耗费时间是主程序创建协程并执行部分协程内函数内容，然后协程进入后台运行，主程序就退出了
	//fmt.Printf("包%s馅儿的包子, 共耗费的时间:%d\n\n", filling, cost)
	time.Sleep(time.Second)
	fmt.Printf("%s馅的包子已经蒸好了,可以上菜了\n", filling)

	buns <- filling
}

func cookDish(chef, disName string, c chan string) {
	fmt.Printf("厨师:%s正在做:%s\n", chef, disName)
	time.Sleep(time.Second * 2)
	c <- disName
}

func screw(c chan int) {
	i := 1
	for {
		fmt.Printf("正在处理第:%d个事情\n", i)
		c <- i
		i++
		time.Sleep(time.Second)
	}
}

func main() {
	/*
				// Hello world
				fmt.Println("Hello world! Hello Go!")

				// 结构体嵌套
				var p Person
				p.Name = "张小凡"
				p.Age = 20
				p.Address = "北京"
				p.Gender = "男"
				p.Mobile.Mode = "HUAWEI"
				p.Mobile.Price = 5699
				fmt.Println("用户信息：", p)
				fmt.Printf("用户：%s，使用的手机是：%s\n", p.Name, p.Mobile.Mode)

				var mobile Phone = Phone{"苹果", "iPhone17", 10000.00}
				var p2 Person
				p2.Name = "秦傲雪"
				p2.Age = 16
				p2.Mobile = mobile
				fmt.Printf("用户：%s，使用的手机是：%s\n", p2.Name, p2.Mobile.Mode)

				var info1 Info
				info1.Person.Name = "青云志"
				info1.Age = 18
				info1.Mode = "xiaomi"
				info1.Phone.Name = "小米"
				info1.Price = 3999
				fmt.Printf("用户：%s，使用的手机是：%s\n", info1.Person.Name, info1.Mode)
				fmt.Println("用户：", info1)

				// 接口
				db := DBConfig{"root", "pwd", "1270.0.01", "3306", "interface_test"}
				var dbCommonInterface DBCommon
				var m MySQL
				m.config = db
				m.charSet = "utf-8"
				dbCommonInterface = m
				dbCommonInterface.Insert("insert xxx data")
				dbCommonInterface.Update("update yyy data")
				dbCommonInterface.Delete("delete zzz data")

				dbType := "pg"
				var dbCommonInterface2 DBCommon
				if dbType == "mysql" {
					var m MySQL
					dbCommonInterface2 = m
				} else {
					var pg PostGreSQL
					dbCommonInterface2 = pg
				}
				dbCommonInterface2.Insert("insert")
				dbCommonInterface2.Update("update")

			// 空接口
			var ei EmptyInterface
			var ei2 interface{}
			s1 := "这是一个字符串"
			i1 := 12345
			ei = s1
			ei2 = i1
			fmt.Println("空接口ei=", ei)
			fmt.Println("空接口ei2=", ei2)
			ei2 = s1
			fmt.Println("空接口ei2=", ei2)
			dealData(s1)
			getType(i1)


		buns := make(chan string, 5)
		// 关闭通道
		defer close(buns)

		// 协程
		fillings := []string{"韭菜鸡蛋", "牛肉大葱", "西葫芦鸡蛋"}
		startTime := time.Now()
		for _, v := range fillings {
			// 使用协程
			go makeBuns(v, buns)
		}
		for i:= 0; i < len(fillings); i++ {
			// 如果通道内没有数据，就会一直在此处阻塞，直到取到数据为止
			bun := <-buns
			fmt.Printf("上菜: %s, 上菜时间:%s\n\n", bun, time.Now())
		}
		cost := time.Since(startTime)
		// 使用协程后，耗费时间是主程序创建协程并执行部分协程内函数内容，然后协程进入后台运行，主程序就退出了
		fmt.Println("共耗费的时间:", cost)

		time.Sleep(time.Second * 3) // 等待3秒，会打印出使用协程的方法中剩余部分内容，仅仅用于验证。开发时程序中不能用time.Sleep去等待！！！
		// 通道超时处理
		chef1 := make(chan string)
		chef2 := make(chan string)
		go cookDish("chef1", "烤鸭", chef1)
		go cookDish("chef2", "佛跳墙", chef2)
		select {
		case dish := <- chef1:
			fmt.Printf("厨师chef1已经做好了:%s\n", dish)
		case dish := <- chef2:
			fmt.Printf("厨师chef2已经做好了:%s\n", dish)
		case <-time.After(time.Second * 3):
			// 对通道进行超时处理
			fmt.Println("你们做饭太慢了，我不吃了，拜拜")
		}
		close(chef1)
		close(chef2)

	screwChan := make(chan int, 20)
	stop := make(chan bool)
	go screw(screwChan)
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("工作时间结束，下班了")
		stop <- true
	}()
	for {
		select {
		case <-stop:
			fmt.Println("时间到")
			// return之后就退出程序了
			return
		case s := <-screwChan:
			fmt.Printf("第%d个事情已完成\n", s)
		}
	}
	*/
	
	deflog.Debug("这是一个debug日志，用来描述程序的运行日志")

}
