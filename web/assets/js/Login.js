function user_action(ac) {
    console.log(ac);
    var name = $("#signin-name").val();
    var pawd = $("#signin-password").val();
    var email = $("#signin-email").val();
    switch (ac){
        case 'sign_up':

            if (!IsMatchingAddress(email)){
                message('error','邮箱格式错误');
                return
            }

            if(name.length < 4 ){
                message('info',"用户名过短");
                return
            } else if(name.length > 15){
                message('info',"用户名过长");
                return
            }

            // if (!IsMatchingUserName(name)){
            //     message('error',"用户名格式错误");
            //     return
            // }

            if(pawd.length < 6 ){
                message('info',"密码过短");
                return
            } else if(pawd.length > 15){
                message('info',"密码过长");
                return
            }

            if (!IsMatchingUserPassword(pawd)){
                message('error',"密码中不允许特殊字符");
                return
            }

            user_request(sign_up_success(),sign_up_fail(),'sign_up');

            break;
        case 'login':
            if(name.length < 4 ){
                message('info',"用户名过短");
                return
            } else if(name.length > 15){
                message('info',"用户名过长");
                return
            }

            // if (!IsMatchingUserName(name)){
            //     message('error',"用户名格式错误");
            //     return
            // }

            if(pawd.length < 6 ){
                message('info',"密码过短");
                return
            } else if(pawd.length > 15){
                message('info',"密码过长");
                return
            }

            if (!IsMatchingUserPassword(pawd)){
                message('error',"密码中不允许特殊字符");
                return
            }
            user_request(login_success,login_failed,'login');
            break;
    }
}

function sign_up_fail() {
    message('error','注册失败');
}

function sign_up_success() {
    message('success','注册成功');
    sign_up_switch('login');
    message('success','请登录');
}

function login_success() {
    location.replace('/apt');
}

function login_failed() {
    message('error', '用户名或密码错');
}

function IsMatchingAddress(str){
    var myRegExp = /\w@\w*\.\w/ ;
    return myRegExp.test(str)
}


function IsMatchingUserName(str){
    var myRegExp = /[a-z0-9A-Z-]{4,15}/ ;
    return myRegExp.test(str)
}


function IsMatchingUserPassword(str){
    var myRegExp = /[a-z0-9A-Z-]{6,15}/ ;
    return myRegExp.test(str)
}


function user_request(success, error,ac) {
    var name = $("#signin-name").val();
    var pawd = $("#signin-password").val();
    var email = $("#signin-email").val();


    var json = {
        action : ac,
        user_name : name,
        pass_word : pawd,
        email : email
    };

    $.ajax({
        url : '/apt/user/action',
        type: "POST",
        contentType: 'application/json; charset=utf-8', // 很重要
        dataType: "json",
        data: JSON.stringify(json),
        success : function (data) {
            console.log(data);
            if(data.state){
                // log in success
                success();
                console.log('log in success');
            } else {
                error();
            }
        },
        error : function () {
            console.log('error network');
            message('error','通信错误！')
        }
    })

}




function sign_up_switch(acc) {
    switch (acc){
        case 'login':
            $("#log_in_").hide();
            $("#email_box").hide();

            $("#sign_up_").show();
            $("#log_in_button").show();
            $("#sign_up_button").hide();

            break;
        case 'sign_up':
            $("#log_in_").show();
            $("#email_box").show();

            $("#sign_up_").hide();
            $("#log_in_button").hide();
            $("#sign_up_button").show();


            break;
    }
}

function message(context, message) {

    toastr.options.timeOut = 3000;
    toastr.options.closeButton = true;
    $context = context;
    $message = message;
    $position = 'top-center';

    if($context == '') {
        $context = 'info';
    }

    if($position == '') {
        $positionClass = 'toast-left-top';
    } else {
        $positionClass = 'toast-' + $position;
    }

    toastr.remove();
    toastr[$context]($message, '' , { positionClass: $positionClass });
}