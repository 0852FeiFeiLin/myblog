package controllers

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/24 14:05
 * @Description:
 **/
//type FileComController struct {
//	beego.Controller
//}
////简单的文件上传功能
//func (f *FileComController) FileCom() {
//	fmt.Println("hello,filecomm")
//	file, header, err := f.GetFile("file") //返回文件，文件信息头，错误信息
//	if err != nil {
//		f.Ctx.WriteString("File retrieval failure")
//		return
//	}
//	defer file.Close()          //记得延迟关闭文件上传，否则会出现临时文件不清除的情况，
//	filename := header.Filename //将文件头信息赋值给filename变量
//
//	//保存文件，存在static/upload中，（文件名）
//	path := "static/upload/"
//	/* s := []string {"static/upload",filename}
//	err := f.SaveToFile("file", strings.Join(s, "")) //拼接*/
//	err = f.SaveToFile("file", path+filename) //fromFile是提交的时候的html表单中的name
//	if err != nil {
//		f.Ctx.WriteString("File upload failed！")
//		fmt.Println("文件上传失败")
//		return
//	} else {
//		f.Ctx.WriteString("File upload succeed！")
//		fmt.Println("文件上传成功")
//	}
//	f.TplName = "blog.html"
//
//}
