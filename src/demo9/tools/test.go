package tools

import . "fmt"

func Log_int(x int)  {
	Printf("\n值:%#v 类型:%T  指针:%p\n",x,x,&x)
}
