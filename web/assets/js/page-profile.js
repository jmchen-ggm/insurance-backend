//加载用户收藏信息
$(document).ready(
    function(){
        var json={"signal":"InfoOfCollection"};
        $.ajax({
            url:"/apt/homepage/usercollection",
            data:JSON.stringify(json),
            type:"get",
            dataType:"json",
            success:function(data){
                console.log(data);
                console.log(data.errMg);
                if(!data.state){
                    alert("收藏信息查询失败");
                }
                else{
                    var len=data.element.length;
                    var info=data.element;
                    console.log("num of collenction is:"+len);
                    for(var i=0;i<len;i++){
                        $("#InfoOfCollection").append(
                            "<tr>\n" +
                            "\t\t\t\t\t\t\t\t\t\t\t\t\t<td><a href=\"/apt/page_detail.html?data_id="+info[i].data_id+"\">"+info[i].title+"</a></td>\n" +
                            "\t\t\t\t\t\t\t\t\t\t\t\t\t<td>"+info[i].author+"</td>\n" +
                            "\t\t\t\t\t\t\t\t\t\t\t\t</tr>"
                        )
                    }
                    $("#NumOfCollection").text(len);
                }
            },
            error:function(){
                alert("collection后台通信错误");
            }
        })
    }
)

//加载浏览历史信息
// $(document).ready(
//     function(){
//         var json={"signal":"InfoOfHistory"};
//         $.ajax({
//             url:"",
//             data:JSON.stringify(json),
//             type:"get",
//             dataType:"json",
//             success:function(data){
//                 console.log(data);
//                 console.log(data.errMg);
//                 if(!data.state){
//                     alert("浏览历史查询失败");
//                 }
//                 else{
//                     var len=data.element.length;
//                     var info=data.element;
//                     console.log("num of history is:"+len);
//                     for(var i=0;i<len;i++){
//                         $("#InfoOfHistory").append(
//                             "<tr>\n" +
//                             "\t\t\t\t\t\t\t\t\t\t\t\t\t\t<td><a href=\"/apt/page_detail.html?data_id="+data[i].data_id+"\">"+info[i].title+"</a></td>\n" +
//                             "\t\t\t\t\t\t\t\t\t\t\t\t\t\t<td>"+info[i].author+"</td>\n" +
//                             "\t\t\t\t\t\t\t\t\t\t\t\t\t</tr>"
//                         )
//                      }
//                      $("#NumOfHistory").text(len);
//                 }
//             },
//             error:function(){
//                 alert("history后台通信错误");
//             }
//         })
//     }
// );