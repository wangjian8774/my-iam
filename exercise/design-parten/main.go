package main

import "time"

//1 创建型模式 创建对象，隐藏创建细节
//1.1 单例模式
//1.1.1 饿汉模式 开始创建时就初始化了

/*type Hunger struct {
	name string
}

var man *Hunger = &Hunger{"hunger-man"}

func GetInstance() *Hunger {
	return man
}

func main() {
	newMan := GetInstance()
	print(newMan.name)
}*/

//1.1.2 懒汉模式 开始的时候创建，在使用到的时候进行初始化

/*type singleton struct {
}

var sing *singleton

func getInstance() *singleton {
	if sing == nil {
		sing = &singleton{}
	}
	return sing

}

// 以上这种方法为非并发安全的，需要加锁，或者使用once.do
var mu sync.Mutex

func getInstanceMutex() *singleton {
	if sing == nil {
		mu.Lock()
		if sing == nil {
			sing = &singleton{}
		}
		mu.Unlock()
	}
	return sing
}

func getInstanceOnce() *singleton {
	once := sync.Once{}
	if sing == nil {
		once.Do(func() {
			sing = &singleton{}
		})
	}
	return sing
}*/

// 1.2 工厂模式
// 1.2.1 简单工厂模式 直接返回一个接口体
/*type Car struct {
	name  string
	price float64
}

func NewCar(name string, price float64) *Car {
	return &Car{
		name:  "BMW",
		price: 100,
	}
}*/

// 1.2.2 抽象工厂模式 返回的是一个接口，接口可以使用其方法，不同结构体的实现是不同的
/*type ICar interface {
	run()
}

type BMW struct {
	name string
}

func (car *BMW) run() {
	println(car.name)
}


func newCar(name string) ICar {
	return &BMW{name: "BMW"}
}*/

// 1.2.3 工厂方法 返回一个方法，用这个方法来生产对象
/*type ICar interface {
	run()
}

type Benz struct {
	name string
}

func newCar(name string) func(price float64) ICar {
	return func(price float64) ICar {
		return &Benz{name: name}
	}
}

func (*Benz) run() {
	fmt.Println("Benz run")
}*/

// 2 结构型模式 关注类和对象的组合
// 2.1 策略型模式 就是接口有多个不同的实现struct， 不同struct的function实现是不同的
/*type IStrategy interface {
	do(int, int) int
}

type Add struct {
}

func (*Add) do(a, b int) int {
	return a + b
}

type Reduce struct {
}

func (*Reduce) do(a, b int) int {
	return a - b
}

type OperateStrategy struct {
	IStrategy
}*/

// 2.2 模版模式 “子类”对于“父类”中实现特定的方法(当然Go中没有面向对象的概念，Go中使用的是组合，比如类型嵌入)
/*type ICooker interface {
	fire()
	cooke()
	offFire()
}

type CookMenu struct {
}

func (*CookMenu) fire() {
	println("burning")
}

func (*CookMenu) cooke() {}

func (*CookMenu) offFire() {
	println("extinguish")
}

type DongPoPork struct {
	CookMenu
}

func (*DongPoPork) cooke() {
	println("cooke dongpopork")
}

type VineFish struct {
	CookMenu
}

func (*VineFish) cooke() {
	println("cooke VinegarFish")
}

func doCook(cooker ICooker) {
	cooker.fire()
	cooker.cooke()
	cooker.offFire()
}*/

// 3 行为型模式 针对于对象间的通信
// 3.1 代理模式 代理一个结构体提供服务
/*type Seller interface { // 为什么这里需要一个接口，感觉并没有什么用
	sell(name string)
}

type Station struct {
	stock int
}

func (ser *Station) sell(name string) {
	if ser.stock > 0 {
		ser.stock--
		fmt.Printf("%s buy a ticket, there are %d remaing tickets\n", name, ser.stock)
	} else {
		println("ticket is sold out")
	}
}

type ProxyStation struct {
	*Station
}

func (pSer *ProxyStation) sell(name string) {
	if pSer.stock > 0 {
		pSer.stock--
		fmt.Printf("%s buy a ticket, there are %d remaing tickets\n", name, pSer.stock)
	} else {
		println("ticket is sold out")
	}
}*/

// 3.2 选项模式 WithXXX 用于复杂结构体的初始化，默认值参数就是直接传，可选参数通过选项函数
type Connection struct {
	addr    string
	TimeOut time.Time
	cache   bool
}

type Option interface {
	apply(*options)
}

type options struct {
	TimeOut time.Time
	cache   bool
}

type OptionFunc func(*options)

func (opFunc OptionFunc) apply(op *options) {
	opFunc(op)
}

func WithTimeOut(time time.Time) Option {
	return OptionFunc(func(o *options) {
		o.TimeOut = time
	})
}

func WithCache(cache bool) Option {
	return OptionFunc(func(o *options) {
		o.cache = cache
	})
}

func NewConnection(addr string, opts ...Option) (*Connection, error) {
	op := &options{}

	for _, opt := range opts {
		opt.apply(op)
	}

	return &Connection{
		addr:    addr,
		TimeOut: op.TimeOut,
		cache:   op.cache,
	}, nil
}

// 我个人感觉写法还是有点复杂，不是很好理解，当然在多参数的时候这种方法应该会带来一些好处
