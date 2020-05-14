package main

import . "fmt"

type Humaner interface {
	//方法 只有声明 没有实现
	sayhi()
}

// 学生
type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	Println("student:", tmp.id, tmp.name)
}

// 老师
type Teacher struct {
	name  string
	id    int
	class int
}

func (tmp *Teacher) sayhi() {
	Println("Teacher:", tmp.id, tmp.name, tmp.class)
}

type Goods interface {
	// 方法只有声明 没有实现 由别的类型实现
	submit()
}

type Junwang struct {
	yggc_order string
	price      int
	goods_id   int
}

func (input *Junwang) submit() {
	Printf("军网推送订单信息：%s, %d", input.yggc_order, input.price)
}

type Yangcai struct {
	yggc_order string
	price      int
	goods_id   int
}

func (input *Yangcai) submit() {
	Printf("央采推送订单信息：%s, %d", input.yggc_order, input.goods_id)
}

//定义一个普通函数 函数的参数为接口类型
// 只有一个函数,但是可以有不同表现，多态
func WhoSupplier(goods Goods) {
	goods.submit()
}

func main() {
	// 定义接口类型的变量
	// var good Goods
	// 只要实现了此接口方法的类型  那么这个类型的变量（接收者类型）就可以给 g 赋值
	junwang := &Junwang{"GW123123", 3425, 1}
	yangcai := &Yangcai{"YC23123", 325, 2}
	WhoSupplier(junwang)
	WhoSupplier(yangcai)

	// 创建一个切片
	qiepian := make([]Goods, 2)
	qiepian[0] = junwang
	qiepian[1] = yangcai

	// 第一个返回下标，第二个返回对应值
	for _, i := range qiepian {
		i.submit()
	}

}

func main1() {
	// 定义接口类型变量
	var i Humaner

	s := &Student{"lzk", 1}
	i = s
	i.sayhi()

	t := &Teacher{"wwww", 1, 10}
	i = t
	i.sayhi()

	// 定义接口类型的变量
	var good Goods
	// 只要实现了此接口方法的类型  那么这个类型的变量（接收者类型）就可以给 g 赋值
	junwang := &Junwang{"GW123123", 3425, 1}
	good = junwang
	good.submit()

	yangcai := &Yangcai{"YC23123", 325, 2}
	good = yangcai
	good.submit()

}
