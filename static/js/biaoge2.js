//JavaScript代码区域
layui.use('element', function() {
    var element = layui.element;

});
layui.use('table', function() {
    var table = layui.table;
    table.render({
        elem: '#demo',
        height: 700,
        url: '/admin_user',
        page: true, //开启分页
        limits: [10],
        limit: 10,
        cols: [
            [ //表头
                {
                    field: 'Id',
                    title: 'ID',
                    width: 100,
                    sort: true,
                    fixed: 'left'
                }, {
                    field: 'Username',
                    title: '用户名',
                    width: 150,
                }, {
                    field: 'Genre',
                    title: '类型(1: 普通用户  2: 管理员)',
                    width: 300,
                    sort: true
                }, {
                    fixed: 'right',
                    width: 150,
                    align: 'center',
                    toolbar: '#barDemo'
                }
            ]
        ]
    });

    table.on('tool(users)', function(obj) {
        var data = obj.data; //获得当前行数据
        var layEvent = obj.event; //获得 lay-event 对应的值（也可以是表头的 event 参数对应的值）
        var tr = obj.tr; //获得当前行 tr 的 DOM 对象（如果有的话）
        console.log(data);
        console.log(obj.tr);
        if (layEvent === 'del') { //删除
            obj.del();
            console.log(data.id);
            $.post({
                url: "/admin/user/delete" //向该地址的服务程序发送需要从数据库删除的用户id
            }, {
                remove_id: data.Id
            }, function(data) {
                if (data.code == 1) {
                    alert("删除成功")
                }
            });
        }
    });
});