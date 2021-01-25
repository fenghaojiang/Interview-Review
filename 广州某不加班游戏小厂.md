---
title: 中东棋牌类游戏小厂
date: 2021-1-25
---

去到科韵路那里的做中东棋牌类游戏的小厂，据说不怎么加班，去到那里感觉装修很舒服，是我喜欢的类型。但是问的题目感觉莫名其妙？我感觉面试官水平也就那样，是不是当时就是百度然后随便问的？  

# 算法题  

1. 10w用户做一个即时排名，不能用stl里面的排序，什么常用的排序不能用。  

   其实我的第一反应是快速排序，10w用户量并不算大，一个int32 4个字节， 10w人40w字节，也就400KB内存，我真的是佛了，后来回来问了个懂哥，懂哥说用休眠排序法，百度了一下，先不说有没有用，发明出这种排序方法的人是天才吧我擦，XD  
   休眠法就是构造n个协程/线程，设置一个单位休眠时间，数组里面的值*休眠时间，醒来就输出他们各自的值，当所有的协程都醒来了，排序就已经结束了，发明出这个排序的简直就是懂王之王，都给他懂完了都  
   到最后我也不知道面试官想问什么，后来说了个分区合并有序链表，奥里给，直接负分  


2. 如何判断链表存在环？  

   这个问题实在没意思，我校招的时候被问到不下10次了。快慢指针，我光速回答，一个指针每次走两步，另外一个每次走1步，如果存在环的话那么两个指针会撞在一起，面试官问我，会不会两个指针在环中永远不会相遇，我说不会，他一脸不相信的样子然后没有问我为什么不会，我尼玛，在环中每次走一轮两个指针就会靠近一步，这样都不会相遇我佛了？  


3. 1-6如何模拟出1-10的随机效果？

   想了一下，我说随机5次然后除5，他说不可以，因为随机5次就是6X6X6X6X6种可能，有没有直接10种可能的，我没有想出，有点失望，回来想到是不是1 3 5的个数等于2 4 6的个数用来模拟二进制，但是1-10也不是刚好4位，难顶，回来百度好像也是差不多是随机5次的说法，有点被面试官坑了的感觉。  


4. golang写个SocketServer

```go
package main

import (
	"fmt"
	"log"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("TCP Connection Establish")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go handleMessage(conn)
	}
}

func handleMessage(conn net.Conn) {
	if conn == nil {
		log.Panic("Not valid Socket Connection")
	}
	info := make([]byte, 512)
	for {
		cnt, err := conn.Read(info)
		if cnt == 0 || err != nil {
			conn.Close()
			break
		}
	}
}

func main() {
	StartServer()
}
```




