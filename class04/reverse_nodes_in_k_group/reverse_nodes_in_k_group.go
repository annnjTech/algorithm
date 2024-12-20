package reverse_nodes_in_k_group

import "algorithm/class04/model"

// GetKGroupEnd
/*
    k = 3
    ①->②->③->④->⑤->⑥->⑦->⑧->nil
1) start指向①
2) i = 1, start指向②
3) i = 2, start指向③
4) i = 3, 结束循环，返回start（start此时指向③）
*/
func GetKGroupEnd(start *model.Node, k int) *model.Node {
	for i := 1; i < k && start != nil; i++ {
		start = start.Next
	}
	return start
}

// Reverse
/*
    k = 3
    ①->②->③->④->⑤->⑥->⑦->⑧->nil
1) end最开始是指向③，end = end.Next的意思是将end往前移动一位指向④
2) start是指向①的，cur = start的意思是cur指向start指向的位置，即也指向①
3) 当cur != end时，进入循环
4) 将游标next移动到cur的next指向的位置
5) 将cur的next指向pre指向的位置，此时pre指向nil
6) 将pre指向cur指向的位置，此时pre指向①
7) 将cur指向next指向的位置，此时cur指向②
8) 重复步骤3-7，直到cur == end
步骤图解：
一：当cur != end时，进入循环
   next     start                       end
    |        |                           |
   nil       ①   ->   ②   ->    ③  ->  ④->⑤->⑥->⑦->⑧->nil
    |        |
   pre     cur
二：将游标next移动到cur的next指向的位置
           start                        end
             |                           |
   nil       ①   ->   ②   ->    ③  ->  ④->⑤->⑥->⑦->⑧->nil
    |        |         |
   pre      cur       next
三：将cur的next指向pre指向的位置，此时pre指向nil
           start                        end
             |                           |
   nil  <-   ①        ②   ->    ③  ->  ④->⑤->⑥->⑦->⑧->nil
    |        |         |
   pre      cur       next
四：将pre指向cur指向的位置，此时pre指向①
            start                        end
              |                           |
   nil  <-    ①        ②   ->    ③  ->  ④->⑤->⑥->⑦->⑧->nil
             | |        |
           cur pre    next
五：将cur指向next指向的位置，此时cur指向②
            start                        end
              |                           |
   nil  <-    ①        ②   ->    ③  ->  ④->⑤->⑥->⑦->⑧->nil
              |        | |
             pre     cur next
================================================================================
六：当cur != end时，进入循环
           start                         end
             |                            |
   nil  <-    ①        ②   ->    ③  ->  ④->⑤->⑥->⑦->⑧->nil
              |        | |
             pre     cur next
七：将游标next移动到cur的next指向的位置
           start                         end
             |                            |
   nil  <-    ①        ②   ->    ③  ->  ④->⑤->⑥->⑦->⑧->nil
              |        |          |
             pre     cur         next
八：将cur的next指向pre指向的位置，此时pre指向①，所以cur的next指向①
            start                        end
              |                           |
   nil  <-    ①   <-  ②          ③  ->  ④->⑤->⑥->⑦->⑧->nil
              |        |          |
             pre      cur        next
九：将pre指向cur指向的位置，此时pre指向①
            start                        end
              |                           |
   nil  <-    ①   <-  ②          ③  ->  ④->⑤->⑥->⑦->⑧->nil
                      | |         |
                    cur pre      next
十：将cur指向next指向的位置，此时cur指向③
            start                        end
              |                           |
   nil  <-    ①   <-  ②          ③  ->  ④->⑤->⑥->⑦->⑧->nil
                       |         | |
                      pre      cur next
================================================================================
十一：当cur != end时，进入循环
            start                       end
              |                          |
   nil  <-    ①   <-  ②         ③  ->  ④->⑤->⑥->⑦->⑧->nil
                       |        | |
                      pre     cur next
十二：将游标next移动到cur的next指向的位置
            start                       end
              |                          |
   nil  <-    ①   <-  ②         ③  ->  ④->⑤->⑥->⑦->⑧->nil
                       |         |       |
                      pre       cur     next
十三：将cur的next指向pre指向的位置，此时pre指向②，所以cur的next指向②
            start                        end
              |                           |
   nil  <-    ①   <-  ②   <-  ③         ④->⑤->⑥->⑦->⑧->nil
                       |        |         |
	                  pre      cur       next
十四：将pre指向cur指向的位置，此时pre指向③
            start                        end
              |                           |
   nil  <-    ①   <-  ②   <-   ③        ④->⑤->⑥->⑦->⑧->nil
                                | |       |
                              cur pre    next
十五：将cur指向next指向的位置，此时cur指向④
            start                        end
              |                           |
   nil  <-    ①   <-  ②   <-   ③        ④->⑤->⑥->⑦->⑧->nil
                                |        | |
                              pre      cur next
十六：当cur == end时，退出循环
十七：将start.Next指向end，即将①指向④，这里只是完成了①到③的反转，①的next其实不应该指向④的位置，应该指向⑥的位置的
            start ______________________ end
              | ↗                       ↘ |
              ①   <-  ②   <-   ③        ④->⑤->⑥->⑦->⑧->nil
                                |        | |
                              pre      cur next
================================================================================
*/
func Reverse(start, end *model.Node) {
	end = end.Next
	var pre, next, cur *model.Node
	cur = start
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	start.Next = end
}

// ReverseKGroup
/*
延续上述的十八的结果
一：lastEnd指向start，即lastEnd指向①
            start ______________________ end
              | ↗                       ↘ |
              ①   <-  ②   <-   ③        ④->⑤->⑥->⑦->⑧->nil
              |                 |
            lastEnd            head（第一次找到的end就是最终返回的head）
二：lastEnd.Next != nil，进入循环
            start ______________________ end
              | ↗                       ↘ |
              ①   <-  ②   <-   ③        ④->⑤->⑥->⑦->⑧->nil
              |                 |
            lastEnd            head（第一次找到的end就是最终返回的head）
三：start指向lastEnd的next位置，即start指向④
                 ______________________ start end
                ↗                       ↘  |  |
              ①   <-  ②   <-   ③           ④    ->⑤->⑥->⑦->⑧->nil
              |                 |
            lastEnd            head（第一次找到的end就是最终返回的head）
四：end指向GetKGroupEnd(start, k)，即end指向⑦
                 ______________________ start           end
                ↗                       ↘ |              |
              ①   <-  ②   <-   ③        ④  ->  ⑤  ->  ⑥  ->  ⑦  ->  ⑧  ->  nil
              |                 |
            lastEnd            head（第一次找到的end就是最终返回的head）
五：end == nil条件不成立，不进入break
六：Reverse(start, end)，即将④到⑥的反转
                 ______________________ start ________________
                ↗                       ↘ | ↗                  ↘
              ①   <-  ②   <-   ③        ④  <-  ⑤  <-  ⑥      ⑦  ->  ⑧  ->  nil
              |                 |                        |
            lastEnd           head                      end
七：lastEnd.Next = end，即将lastEnd的next指向end，此时lastEnd的next指向⑦
                                head    start ________________
                                 |        | ↗                  ↘
              ①   <-  ②   <-   ③        ④  <-  ⑤  <-  ⑥      ⑦  ->  ⑧  ->  nil
              |  ↘                                     ↗ |
         lastEnd   -----------------------------------  end
八：lastEnd = start，即lastEnd指向④
                               head   lastEnd start ______________________
                                 |           | |   ↗                       ↘
              ①   <-  ②   <-   ③            ④          <-  ⑤  <-   ⑥     ⑦  ->  ⑧  ->  nil
                ↘                                                ↗ |
                 -----------------------------------------------  end
九：lastEnd.Next != nil，继续循环
                               head    lastEnd start _____________________
                                 |           | |   ↗                       ↘
              ①   <-  ②   <-   ③            ④          <-  ⑤  <-   ⑥     ⑦  ->  ⑧  ->  nil
                ↘                                                ↗ |
                 -----------------------------------------------  end
十：start指向lastEnd的next位置，即start指向⑦
                               head    lastEnd    _____________________   start
                                |          |    ↗                       ↘  |
              ①   <-  ②   <-   ③         ④          <-  ⑤  <-   ⑥      ⑦  ->  ⑧  ->  nil
                ↘                                                ↗ |
                 -----------------------------------------------  end
十一：end指向GetKGroupEnd(start, k)，即end指向nil
                               head     lastEnd  ______________________   start
                                |          |    ↗                       ↘  |
              ①   <-  ②   <-   ③         ④          <-  ⑤  <-   ⑥      ⑦  ->  ⑧  ->  nil
                ↘                                                ↗                         |
                 -----------------------------------------------                          end
十二：end == nil条件成立，进入break
*/
/*
	k = 3
    ①->②->③->④->⑤->⑥->⑦->⑧->nil
结果：
    ③->②->①->⑥->⑤->④->⑦->⑧->nil
*/
func ReverseKGroup(head *model.Node, k int) *model.Node {
	start := head
	end := GetKGroupEnd(start, k)
	if end == nil {
		return head
	}
	// 第一组凑齐了
	head = end
	Reverse(start, end)
	// 上一组的结尾节点
	lastEnd := start
	for lastEnd.Next != nil {
		start = lastEnd.Next
		end = GetKGroupEnd(start, k)
		if end == nil {
			break
		}
		Reverse(start, end)
		lastEnd.Next = end
		lastEnd = start
	}
	return head
}
