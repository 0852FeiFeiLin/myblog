/*
* 该函数用于统计页面输入的字数
* */
function Words(){
    var nums = 0;
    var title =  document.getElementById("title").value.length;
    var tags =  document.getElementById("tags").value.length;
    var short =  document.getElementById("short").value.length;
    var content =  document.getElementById("content").value.length;
    nums = nums + title + tags + short + content;
    document.getElementById("words").innerHTML= nums
}

/*
$(function(){
    $("#content").focus(function () {
        $('textarea').each(function () {
            this.setAttribute('style', 'height:' + (this.scrollHeight) + 'px;overflow-y:hidden;');
        }).on('input', function () {
            this.style.height = 'auto';
            this.style.height = (this.scrollHeight) + 'px';
        });
    })
});*/
