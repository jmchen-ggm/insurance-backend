function message(context, message) {
    toastr.options.timeOut = 3000;
    toastr.options.closeButton = true;
    $context = context;
    $message = message;
    $position = 'top-center';

    if ($context == '') {
        $context = 'info';
    }

    if ($position == '') {
        $positionClass = 'toast-left-top';
    } else {
        $positionClass = 'toast-' + $position;
    }
    toastr.remove();
    toastr[$context]($message, '', {positionClass: $positionClass});
}


function getUserInfo() {
// ajax to get user info
    $.ajax({
        url: '/apt/user/info',
        type: "POST",
        contentType: 'application/json; charset=utf-8', // 很重要
        dataType: "json",
        success: function (data) {
            console.log(data);
            if (data.state) {
                // log in success
                //{"user_name":"liyu","email":"sdf@ds.com"}
                var user = data.element;
                $("#user_name").html(user.user_name);
                sessionStorage.setItem('user_info', user);
                console.log('log in success');
            } else {
                // error
                if (data.errMsg == 'not logged') {
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

function json_data_to_html(json) {

    var t_content = json.content;
    t_content = t_content.substr(0, 300) + '...';
    var content = "" +
        "<div class=\"col-md-9 col-lg-6 col-sm-12\">\n" +
        "        <div onclick='toDetail(this)' data-id='" + json.data_id + "' class=\"panel ly_panel\">\n" +
        "        <div data-id='" + json.data_id + "' class=\"panel-heading\">\n" +
        "        <h1 class=\"panel-title\">" + json.title + "</h1>\n" +
        "    <p  class=\"panel-note\">Anthor ： " + json.author + ",</p>\n" +
        "\n" +

        "\n" +
        "    </div>\n" +
        "\n" +
        "    <div data-id='" + json.data_id + "' class=\"panel-body\">\n" +
        "\n" +
        "        <p>" + t_content + "</p>\n" +
        "    <div class=\"panel-note\">\n" +
        "        crawl time : " + json.c_time + "\n" +
        "        </div>\n" +
        "        </div>\n" +
        "\n" +
        "        </div>\n" +
        "        </div>";

    return content;
}

function toDetail(target) {
    target = $(target);
    var id = target.data('id');
    location.replace('/apt/page_detail.html?data_id=' + encodeURI(id));
}



function daily_render(array) {
    var news_container = $("#daily_news_start");
    for (var i = 0; i < array.length; i++) {
        var json = array[i];
        news_container.append(
            json_data_to_html(json)
        )
    }

    if (array.length == 0) {
        message('info', '没有更多了');
    } else {
        message('info', '刷新成功');
    }
}


function get_daily() {
    var start = parseInt(0);

    var tem = sessionStorage.getItem('home_start');
    if (tem == null || tem == 'NaN') {
        sessionStorage.setItem('home_start', 0);
    } else {
        start = parseInt(tem);
    }

    message('info', '获取中');
    var end = start + 10;
    get_data_offset(start, end, daily_render);
    start++;
    sessionStorage.setItem('home_start', start);
}


function get_data_offset(start, end, callback) {
    var json = {
        start: start,
        end: end
    };

    $.ajax({
        url: '/apt/get/data/offset',
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

function toggle_collape(target) {
    $(target).clickToggle(
        function(e) {
            e.preventDefault();

            // if has scroll
            if( $(this).parents('.panel').find('.slimScrollDiv').length > 0 ) {
                affectedElement = $('.slimScrollDiv');
            }

            $(this).parents('.panel').find(affectedElement).slideUp(300);
            $(this).find('i.lnr-chevron-up').toggleClass('lnr-chevron-down');
        },
        function(e) {
            e.preventDefault();

            // if has scroll
            if( $(this).parents('.panel').find('.slimScrollDiv').length > 0 ) {
                affectedElement = $('.slimScrollDiv');
            }

            $(this).parents('.panel').find(affectedElement).slideDown(300);
            $(this).find('i.lnr-chevron-up').toggleClass('lnr-chevron-down');
        }
    );
}