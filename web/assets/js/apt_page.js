
function get_update() {
    var start = 0;
    var tem = sessionStorage.getItem('apt_start');
    if (tem == null || tem == 'NaN') {
        sessionStorage.setItem('apt_start', 0);
    } else {
        start = parseInt(tem);
    }

    message('info', '获取中');
    var end = start + 10;
    end = parseInt(end);
    request_apt(start, end, apt_render);
    start++;
    sessionStorage.setItem('apt_start', start);
}

function go_to_view(String) {
    $("#ly_monitor").show();
    $("#ly_monitor").attr('src',String);
}

function apt_render(array) {
    for (var i = 0; i < array.length ; i ++){
        $("#ly_apt_container").append(
            "\t\t\t\t\t\t\t<div onclick=\"go_to_view('"+array[i]+"')\" class=\"panel ly_panel\">\n" +
            "\t\t\t\t\t\t\t\t<div class=\"panel-heading\">\n" +
            "\t\t\t\t\t\t\t\t\t<h1 class=\"panel-title\">主题爬虫 链接</h1>\n" +
            "\t\t\t\t\t\t\t\t</div>\n" +
            "\n" +
            "\t\t\t\t\t\t\t\t<div class=\"panel-body\">\n" +
            "\t\t\t\t\t\t\t\t\t<p><a href='"+array[i]+"'>"+array[i].substr(0,30)+"</a></p>\n" +
        "\t\t\t\t\t\t\t\t</div>\n" +
        "\t\t\t\t\t\t\t</div>"
         )
    }

}


function request_apt(st, ed, callback) {
    var json = {
        start : st,
        end : ed
    };
    $.ajax({
        url: '/apt/get/apt',
        type: "POST",
        contentType: 'application/json; charset=utf-8', // 很重要
        dataType: "json",
        data: JSON.stringify(json),
        success: function (data) {
            console.log(data);
            if (data.state) {
                var array = data.element;
                callback(array);
            } else {
                // error
                if (data.errMsg == 'error request') {
                    message('error', '您还未登录！');
                    location.replace('/apt/page-login.html')
                } else {
                    message('error', '服务器内部错误');
                }
            }
        },
        error: function () {
            console.log('error network');
            message('error', '通信错误！')
        }
    })
}