package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var total struct {
	sync.Mutex
	value int
}
var m2 = sync.Map{}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		total.Lock()
		total.value += 1
		total.Unlock()
	}
}
func f1(ch1 chan int) {

	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1, ch2 chan int) {
	for ret := range ch1 {
		ch2 <- ret * ret
	}
	close(ch2)
}

func set(key int) {
	m2.Store(key, key+100)
}
func get(key int) interface{} {
	ret, _ := m2.Load(key)
	return ret
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			set(i)
			fmt.Println("i :", get(i))
			wg.Done()
		}(i)
		wg.Wait()
	}

	/*wg.Add(2)
	    go worker(&wg)
		go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)*/
	//ch1 := make(chan int, 100)
	//ch2 := make(chan int, 100)
	//
	//go f1(ch1)
	//go f2(ch1, ch2)
	//
	//for ret := range ch2{
	//	fmt.Println(ret)
	//}
}

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }() //回溯过程
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	func() {

	}()
	dfs(root, targetSum)
	return
}
