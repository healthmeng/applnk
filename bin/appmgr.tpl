<html>
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
<title>
应用管理
</title>
<link href="http://imgsrc.baidu.com/imgad/pic/item/f2deb48f8c5494ee4982ea3f26f5e0fe99257e70.jpg" rel="shortcut icon">
<script type="text/javascript">
function oneditapp(name,id)
{
  /*  alert(name+"ID:"+id);
    document.forms["appmgr"].editid.value=id;
    document.forms["appmgr"].submit();
*/
	var s=window.showModalDialog("/appmgr/editapp?appid="+id,id,"dialogwidth:260px;dialogheight:305px;resizable:no");
	if (s==1){
		window.location.reload(true);
	}
}

function ondelapp(name,id)
{
    var ret=confirm("确定要删除应用\""+name+"\"的所有信息吗？\n(如果不希望应用出现在下载列表中，可以暂时通过让该应用下线来实现)");
    if (ret==true){
       	window.location.href="/appmgr/delapp?appid="+id;
    }
}

var editwin;
var timer;

function addapp()
{
//    var s=window.showModalDialog("/appmgr/addapp","自动分配","dialogwidth:260px;dialogheight:305px;resizable:no");
	if (editwin!=undefined && editwin.closed==false){
		editwin.focus();
		return;
	}
	editwin=window.open('/appmgr/addapp','','modal=yes,depended=yes,z-look=yes,height=330,width=260,top=70,left=50,toolbar=no,menubar=no,scrollbars=yes,resizable=no,status=no');
	//window.open('/appmgr/addapp','自动分配','height=305,width=260,toolbar=no,menubar=no,scrollbars=yes,resizable=no,location=no,status=no');
	//editwin=window.open('/appmgr/addapp?controlName=btnOpenCompleted&',"自动分配",'modal=yes,height=305,width=260,top='+(screen.height-305)/2+',left='+(screen.width-260)/2+',toolbar=no,menubar=no,scrollbars=yes,resizable=no,location=no,status=no');
   // if (s==1)
     //   window.location.reload();
	timer = window.setInterval("IfWindowClosed()", 1000);
}

function IfWindowClosed() {
// if (editwin.closed == true) {
  //	window.clearInterval(timer);
	if (editwin!=undefined ){
		if(editwin.closed==true){
  			window.clearInterval(timer);
			editwin=undefined;
			appmgr.retval.value=="";
		}else{
			if (appmgr.retval.value=="ok"){
  				window.clearInterval(timer);
				editwin.close();
				appmgr.retval.value="";
				window.location.reload(true);
			}
		}
	}
}

</script>
</head>
<body>
<pre>
<form name="appmgr" action="/appmgr" method="post">
<font style="font-size:28px;vertical-align: middle;color: blue">应用管理</font>&nbsp;&nbsp;<img src="/localapp/img/add.jpg" alt="添加应用" style="height:32px; width:32px;vertical-align:middle" onclick="javascript:addapp();" />
</p>
{{AppEdit}}
</p>
<input type="hidden" name="editid" />
<input type="hidden" name="retval" value="cancel" />
</form>
</pre>
</body>
</html>
