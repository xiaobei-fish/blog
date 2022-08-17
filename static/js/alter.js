$("#alter").click(
    function() {
        layer.open({
            type: 1,
            skin: 'alter',
            title: ['修改密码', 'font-size:25px'],
            area: ['500px', '480px'],
            shadeClose: true,
            content: '\<\div class="layui-form" style="padding:40px;">' +
                '<form id="al" class="form_1 layui-form" action="http://127.0.0.1:8080/alter" method="post">' +
                '<div class="layui-form-item">' +
                '<div class=""><input id="p1" class="log_input" type="text" required name="oldPassword" placeholder="请输入旧密码"/>\<\/div>' +
                '<div class=""><input class="log_input" type="password" lay-verify="newPassword" required name="newPassword" placeholder="请输入新密码"/>\<\/div>' +
                '<div class=""><input lay-filter="aaa" class="log_submit" type="submit" value="修改"/>\<\/div>' +
                '</form>\<\/div>\<\/div>' //这里是修改密码表单
        })
        //转json
        var o1 = {
            url: '/alter', //数据发送地址，就是action
            type: 'post',
            dataType: 'json',
            success: function(data, status) {
                alert("修改状态:" + data.message + "   请您重新登录")
                if (data.code == 1) {
                    setTimeout(function() {
                        window.location.href = "/home"
                    }, 1000)
                }
            },
            error: function(data, status) {
                alert("修改状态:" + data.message)
            }
        }
        layui.use('form', function() {
            var form = layui.form;
            form.verify({
                newPassword: function(value, item) {
                    if ($("#p1")[0].value == value) {
                        return '两次密码不能一致'
                    }
                    if (!/^[\S]{6,12}$/.test(value)) {
                        return '密码必须在6-12位且无空格'
                    }
                }
            })
            form.on('submit(aaa)', function(data) {});
        })
        $("#al").submit(function() {
            $(this).ajaxSubmit(o1);
            return false
        })
    }
)