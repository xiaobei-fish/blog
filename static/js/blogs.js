$(document).ready(function () {
    //添加文章的表单
    $("#write-art-form").validate({
        rules: {
            title: "required",
            tags: "required",
            short: {
                required: true,
                minlength: 2
            },
            content: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            title: "请输入标题",
            tags: "请输入标签",
            short: {
                required: "请输入简介",
                minlength: "简介内容最少两个字符"
            },
            content: {
                required: "请输入文章内容",
                minlength: "文章内容最少两个字符"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/write";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/mine"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    //修改信息的表单
    $("#message-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            phone: {
                required: true,
                rangelength: [11, 11]
            },
            userId : {
                required: true
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            phone: {
                required: "请输入电话号码",
                rangelength: "电话必须是11位"
            },
            userId: {
                required: "请输入用户Id"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/selfhome";
            alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status);
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/selfhome"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status);
                }
            });
        }
    });

    //修改博客
    $("#update-art-form").validate({
        rules: {
            title: "required",
            tags: "required",
            short: {
                required: true,
                minlength: 2
            },
            content: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            title: "请输入标题",
            tags: "请输入标签",
            short: {
                required: "请输入简介",
                minlength: "简介内容最少两个字符"
            },
            content: {
                required: "请输入文章内容",
                minlength: "文章内容最少两个字符"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/update";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/mine"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    //收藏
    $("#blog-collect-form").validate({
        submitHandler: function (form) {
            var urlStr = "/blog_collect";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    //点赞
    $("#blog-like-form").validate({
        submitHandler: function (form) {
            var id = $("#like-article-id").val();
            var urlStr = "/like";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                    window.location.href = "/detail?id=" + id
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    //添加好友
    $("#add-friend-form").validate({
        rules: {
            phone: {
                required: true,
                rangelength: [11, 11]
            },
            userId : {
                required: true
            }
        },
        messages: {
            phone: {
                required: "请输入电话号码",
                rangelength: "电话必须是11位"
            },
            userId: {
                required: "请输入用户Id"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/add";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/friend"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message);
                }
            });
        }
    });

    //留言
    $("#note-form").validate({
        rules: {
            content: {
                required: true,
                minlength: 2
            },
            noted_id: {
                required: true
            }
        },
        messages: {
            content: {
                required: "请输入留言内容",
                rangelength: "最少输入2个字符"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/note";
            var notedId = $("#noted_id").val();
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/note?id=" + notedId
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message);
                }
            });
        }
    });

    //文件
    $("#album-upload-button").click(function () {
        var filedata = $("#album-upload-file").val();
        if (filedata.length <= 0) {
            alert("请选择文件!");
            return
        }
        //文件上传通过Formdata去储存文件的数据
        var data = new FormData()
        data.append("upload", $("#album-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/img"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert("data:" + data.message);
                if (data.code == 1) {
                    setTimeout(function () {
                        window.location.href = "/selfhome"
                    }, 1000)
                }
            },
            error: function (data, status) {
                alert("err:" + data.message + ":" + status)
            }
        })
    })

});