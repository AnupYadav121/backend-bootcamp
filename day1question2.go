package main

import "fmt"

type Node struct {
    val string
	left *Node
	right *Node
}


func getPreOrder(pre string ,node *Node) string{
    pre += node.val
	if(node.left != nil){
	    getPreOrder(pre,node.left)
	}
	if(node.right != nil){
	    getPreOrder(pre,node.right)
	}
	
	return pre
}

func getPostOrder(pre string,node *Node) string{
    
	if(node.left != nil){
	    getPreOrder(pre,node.left)
	}
	if(node.right != nil){
	    getPreOrder(pre,node.right)
	}
	pre += node.val
	
	return pre
	
}



func main() {
    n1 := Node{val:"a",left:nil,right:nil}
	n2 := Node{val:"b",left:nil,right:nil}
	n3 := Node{val:"c",left:nil,right:nil}
	n4 := Node{val:"-",left:&n2,right:&n3}
	var ans string 
	root := &Node{val:"+",left:&n1,right:&n4}
	fmt.Println(getPreOrder(ans,root))
	fmt.Println(getPostOrder(ans,root))
	
	
}
