//登录弹窗
var p1;
var p2;

$(".login").click(
    function() {
        layer.open({
                type: 1,
                skin: 'login-w',
                title: ['登录', 'font-size:25px'],
                area: ['500px', '480px'],
                shadeClose: true,
                content: '\<\div class="layui-form" style="padding:40px;">' +
                    '<form id="l" class="form_1 layui-form" action="http://127.0.0.1:8080/login" method="post">' +
                    '<div class="layui-form-item">' +
                    '<div class=""><input class="log_input" type="text" required name="username" placeholder="请输入用户昵称或电话号码"/>\<\/div>' +
                    '<div class=""><input class="log_input" type="password" required name="password" placeholder="请输入密码"/>\<\/div>' +
                    '<div class=""><input class="log_submit" type="submit" value="登录"/>\<\/div>' +
                    '<div class=""><a href="/forget" style="float: right;font-size:20px;">忘记密码?</a>\<\/div>' +
                    '</form>\<\/div>\<\/div>' //这里是登录表单
            })
            //转json
        var o1 = {
            url: '/login', //数据发送地址，就是action
            type: 'post',
            dataType: 'json',
            success: function(data, status) {
                alert("登录状态:" + data.message)
                if (data.code == 1) {
                    setTimeout(function() {
                        window.location.href = "/"
                    }, 1000)
                }
            },
            error: function(data, status) {
                alert("登录状态:" + data.message)
            }
        }
        $("#l").submit(function() {
            $(this).ajaxSubmit(o1);
            return false
        })
    }
)
    //注册弹窗
$(".sign").click(
    function() {
        layer.open({
                type: 1,
                skin: 'sign-w',
                title: ['注册', 'font-size:25px'],
                area: ['500px', '580px'],
                shadeClose: true,
                content: '\<\div class="layui-form" style="padding:40px;">' +
                    '<form id="s" class="form_1 layui-form" action="http://127.0.0.1:8080/register" method="post">' +
                    '<div class="layui-form-item">' +
                    '<div class=""><input class="log_input" type="text" required name="username" placeholder="请输入用户昵称"/>\<\/div>' +
                    '<div class=""><input class="log_input" type="phone" lay-verify="phone" required name="phone" placeholder="请输入电话号码"/>\<\/div>' +
                    '<div class=""><input id="p1" class="log_input" type="password" lay-verify="pass" required name="password" placeholder="请设置密码"/>\<\/div>' +
                    '<div class=""><input id="p2" class="log_input" type="password" lay-verify="re_pass" required name="repassword" placeholder="请再次输入密码"/>\<\/div>' +
                    '<div class=""><input lay-submit lay-filter="aaa" class="log_submit" type="submit" value="注册"/>\<\/div>' +
                    '</form>\<\/div>\<\/div>' //这里是注册表单
            })
            //转json
        var o1 = {
            url: '/register', //数据发送地址，就是action
            type: 'post',
            dataType: 'json',
            success: function(data, status) {
                alert("注册状态:" + data.message)
                if (data.code == 1) {
                    setTimeout(function() {
                        window.location.href = "/"
                    }, 1000)
                }
            },
            error: function(data, status) {
                alert("注册状态:" + data.message)
            }
        }
        $("#s").submit(function() {
            $(this).ajaxSubmit(o1);
            return false
        })
        $("#p1").keydown(function(event) {
            if (event.keyCode === 13) {
                p1 = $("#p1")[0].value;
            }
        })
        $("#p2").click(
            function() {
                p1 = $("#p1")[0].value;
            }
        )

    }
)

//表单验证

layui.use('form', function() {
    var form = layui.form;
    form.verify({
        pass: [
            /^[\S]{6,12}$/, '密码必须6到12位，且不能出现空格'
        ],
        re_pass: function(value, item) {
            if (p1 != value) {
                return '两次密码不一致'
            }
        }
    })
    form.on('submit(aaa)', function(data) {
        console.log(data.elem) //被执行事件的元素DOM对象，一般为button对象
        console.log(data.form) //被执行提交的form对象，一般在存在form标签时才会返回
        console.log(data.field) //当前容器的全部表单字段，名值对形式：{name: value}
            //阻止表单跳转。如果需要表单跳转，去掉这段即可。
    });
})



//搜索按钮
$(".ii").click(
    function() {
        $("#in").trigger('click');
    }
)

//下拉菜单

$(document).ready(function() {
    $(".bc").hide()
    $(".lc").mouseover(
        function() {
            $(".bc").show();
        }
    )
    $(".lc").mouseleave(
        function() {
            $(".bc").hide();
        }
    )
})

//轮播图组件启动
layui.use('carousel', function() {
    var carousel = layui.carousel;
    carousel.render({
        elem: "#slideshow",
        width: '620px',
        height: '320px',
        arrow: 'always',
        trigger: 'mouseover'
    });
});

//个人页面布局
layui.use('element', function() {
    var element = layui.element;
    element.render();
});