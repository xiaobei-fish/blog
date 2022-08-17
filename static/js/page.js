function render(count) {
    var text = '';
    var content = count;
    console.log(content);
    $.each(content, function(i, n) {
        console.log(n);
        if (n.Introduction.length > 66) {
            n.Introduction = n.Introduction.slice(0, 66) +
                '<span style="color:rgb(180, 180, 180);">...</span>'
        }
        if (n.Id != "") {
            text = text +
                '<div class="books">' +
                '<img class="b_img" src="' + n.Picture + '" alt="">' +
                '<div class="b_p">' +
                '<p class="b_name">' + n.Name + '<em style="display: none;">' + n.Id + '</em>' +
                '<em title="点击移除" class="iconfont_a" style="float: right;">&#xe61a;</em></p>' +
                '<p class="b_s"><span class="writer">' +
                n.Writer + '&nbsp;著</span>' +
                '<span class="point">&nbsp;&nbsp;' + n.Genres + '</span></p>' +
                '<p class="b_t">' + n.Introduction + '</p>' +
                '</div></div>'
        }
    })
    $("#cont").html(text);
    $(".books").click(function() {
        var id = $(this).find("em:eq(0)").html();
        location.href = '/book/' + id;
    })
}

function bookRemove() {
    $(".iconfont_a").click(function(e) {
        e.stopPropagation(); //阻止事件冒泡
        var id = $(this).siblings().first().html(); //被删除于指定存储位置的书籍
        console.log(id);
        /*
        跨域问题，options请求时查看当前接口 都支持哪些请求方式得 结果在 Response Headers 下 Allow中体现;
解决： 前端设置下 请求头 withCredentials = true; 后端要改成支持跨域形式的
*/
        $.ajax({
            type: "POST",
            url: "/page", //处理此操作的服务器地址
            data: {
                id:id
            } //被删除的书籍id
        })
        $(this).parent().parent().parent().remove();
    })
}
var page_num = 1;

$.ajax({
    type: "GET",
    url: "/collect/?page=1",
    async: false,
    success: function(data) {
        page_num = data.num;
        render(data.novel);
        bookRemove();
        console.log("全局变量page_num : " + page_num + " , 后台数据 data .num : " + data.num);
    }
})


layui.use('laypage', function() {
    var laypage = layui.laypage;
    console.log("全局变量page_num : " + page_num);
    laypage.render({
        elem: 'h_h',
        count: page_num, //100, //co, //数据总数
        limit: 8, //每页限制数据数量
        jump: function(obj, first) {
            if (!first) {
                $.get({ url: "/collect/?page=" + obj.curr }, function(data) {
                    render(data.novel);
                    bookRemove();
                })
            }
        }
    })
})