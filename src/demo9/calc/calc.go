package calc  //自定义包名可以不和文件夹同名，但是建议同名操作

//大写表示共有的方法，只有共有的才导入

func Add(x,y int) (sum int) {
	sum=x+y
	return
}
func Sub(x,y int) int {
	return  x-y
}

type User struct {
	Name string
	Age int
	Sex string
	Id int
	CreateTime string
}