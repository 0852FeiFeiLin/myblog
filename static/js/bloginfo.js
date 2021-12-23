function input_disabled() {
    //先获取到所有及input集合
    inputs = document.getElementsByTagName("input");
    trexts = document.getElementsByTagName("textarea")
    // 一个一个获取到input元素，然后改变他们的属性
    for (i = 0; i < inputs.length; i++) {
        $(inputs[i]).attr("disabled", true)
    }
    for (i = 0; i < trexts.length; i++) {
        $(trexts[i]).attr("disabled", true)
    }
    console.log("我把所有的input变成了不可编辑状态，点击编辑就能编辑了")
    $("#sub").hide();
    $("#end").hide();

}

$(function () {
    $("#edit").click(function () {
        console.log('我点击了编辑，现在都能编辑了');
        /*
            一个一个获取到input元素，然后改变他们的属性，然后就是设置disabled = false(可用)、true(不可用)
            readonly" = true(只读)"readonly" = false(非只读)*/
        $("#title").attr("disabled", false);
        $("#tags").attr("disabled", false);
        $("#short").attr("disabled", false);
        $("#content").attr("disabled", false);
        $("#sub").attr("disabled", false);

        //把删除博客,编辑隐藏，看上去在编辑博客页面。其实并没有跳转页面
        $("#del").hide();
        $("#edit").hide();
        //只显示发布和页脚
        $("#sub").show();
        $("#end").show();
    })

})

function Words1(){
    var nums = 0;
    var title =  document.getElementById("title").value.length;
    var tags =  document.getElementById("tags").value.length;
    var short =  document.getElementById("short").value.length;
    var content =  document.getElementById("content").value.length;
    nums = nums + title + tags + short + content;
    document.getElementById("words").innerHTML= nums
}
