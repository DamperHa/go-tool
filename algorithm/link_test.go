package algorithm

import (
	"fmt"
	"testing"
)

type Node struct {
	next *Node
	val  int
}

func constructLinkFromArr(arr []int) *Node {
	dummy := &Node{}
	cur := dummy

	for _, i := range arr {
		node := &Node{
			next: nil,
			val:  i,
		}

		cur.next = node
		cur = cur.next
	}

	return dummy.next
}

func printLink(head *Node) {
	cur := head

	for cur != nil {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
	}
}

// 给链表排序
func sortList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	head2 := split(head)
	sortList(head)
	sortList(head2)
	return merge(head, head2)
}

// 特例，奇数个：只有一个元素时
// 偶数个：这套逻辑也是ok的
func split(head *Node) *Node {
	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	second := slow.next
	slow.next = nil

	return second
}

func merge(head1 *Node, head2 *Node) *Node {
	dummy := &Node{}
	cur := dummy

	for head1 != nil && head2 != nil {
		if head1.val < head2.val {
			cur.next = head1
			head1 = head1.next
		} else {
			cur.next = head2
			head2 = head2.next
		}

		cur = cur.next
	}

	if head1 != nil {
		cur.next = head1
	}

	if head2 != nil {
		cur.next = head2
	}

	return dummy.next
}

func TestConstructLinkFromArr(t *testing.T) {
	arr := []int{1, 3, 2}
	head := constructLinkFromArr(arr)

	printLink(head)
	sortList(head)

	fmt.Println(head)
	printLink(head)
}
