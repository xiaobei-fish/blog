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
                '<em title="点击收藏" class="iconfont_a btn" style="float: right;">&#xe605;</em></p>' +
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
    $(".btn").click(function(e) {
        e.stopPropagation(); //阻止事件冒泡
        console.log("正确触发");
        if ($("li").is("#userId")) {

            var userId = $("#userId").html();
            var bookId = $(this).siblings().html();
            console.log("登录状态确定：" + userId + '//' + userId);
            /*
                添加收藏夹功能
            */
            $.ajax({
                url: "/result", //操作地址
                type: "post",
                data: {
                    userId: userId,
                    bookId: bookId
                },
                dataType: "json",
                success: function(data) {
                    if (data.code == 1) {
                        alert("添加成功！");
                    } else {
                        alert("已在收藏夹中");
                    }
                }
            })
        } else {
            alert("请先登录");
        }
    })
}
var page_num = 1;

$("#if").submit(function(e) {
    $("#out").html("搜索结果");
    e.preventDefault();
    var data_put = '';
    console.log($("#if").serialize() + '&page=' + 1);
    data_put = $("#if").serialize();
    $.ajax({
        type: "GET",
        url: "/search/?page=1", //提交地址
        data: data_put,
        async: false,
        success: function(data) {
            if (data.code == 1) {
                $("#out").html("未找到匹配书籍，为您推荐一下结果")
            }
            page_num = data.num;
            render(data.novel);
            console.log("全局变量page_num : " + page_num + " , 后台数据 data .num : " + data.num);
        }
    });
    layui.use('laypage', function() {
        var laypage = layui.laypage;
        console.log("全局变量page_num : " + page_num);
        laypage.render({
            elem: 'h_h',
            count: page_num, //co, //数据总数
            limit: 8, //每页限制数据数量
            jump: function(obj, first) {
                if (!first) {
                    $.get({ url: "/search/?page=" + obj.curr }, data_put, function(data) {
                        render(data.novel);
                    })
                }
            }
        })
    });
    return false;
})