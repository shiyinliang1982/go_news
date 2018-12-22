package main

import (
	"github.com/astaxie/beego"
	_ "news/models"
	_ "news/routers"
)

func main() {
	/*关联上一页功能视图函数和后台函数
	* 第一个参数是上一页功能的视图函数
	* 第二个参数是上一页功能的后台函数
	 */
	beego.AddFuncMap("prePage", ShowPrePage)

	/*关联下一页功能视图函数和后台函数
	* 第一个参数是下一页功能的视图函数
	* 第二个参数是下一页功能的后台函数
	 */
	beego.AddFuncMap("nextPage", ShowNextPage)

	beego.Run()
}

/*上一页功能的后台函数
* 参数为pageIndex 类型为int
 */
func ShowPrePage(pageIndex int) int {
	//避免页码值小于1
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

/*下一页功能的后台函数
* 参数为pageIndex,pageCount 类型为int
 */
func ShowNextPage(pageIndex, pageCount int) int {
	//避免页码值大于总页码
	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1
}
