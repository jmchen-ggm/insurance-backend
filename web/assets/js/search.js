function main_search_switch(str) {
    var main = $("#main_page");
    var srch = $("#search_page");
    if(str == 'main'){
        main.slideDown();
        srch.hide();
    } else if (str == 'search'){
        main.slideUp();
        srch.show();
    } else {
        location.replace('/apt');
    }
}

function search_render(array) {
    for (var i = 0; i < array.length && i < 50; i ++){
        console.log(json_data_to_html(array[i]));
        $("#search_result").append(
            json_data_to_html(array[i])
        )
    }

    if(array.length > 0){
        message('info','查询成功');
    } else {
        message('info','无结果');
    }
}

function search() {
    var key_word = $("#search_box");
    var json = {
        demand : key_word.val()
    };
    $.ajax({
        url: '/apt/homepage/search_info',
        type: "POST",
        contentType: 'application/json; charset=utf-8', // 很重要
        dataType: "json",
        data: JSON.stringify(json),
        success: function (data) {
            console.log(data);
            if (data.state) {
                var array = data.element;
                search_render(array);
            } else {
                // error
                message('error', '服务器内部错误');
            }
        },
        error: function () {
            console.log('error network');
            message('error', '通信错误！')
        }
    })
}