<html>
<head>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
<title>
应用详情
</title>
<script type="text/javascript">
window.onload=function (){
	document.getElementById("editid").value=dialogArguments;
	document.getElementById("appid").value=dialogArguments;
	if(dialogArguments=="自动分配"){
		document.getElementsByName("appname")[0].value="";
		document.getElementsByName("appurl")[0].value="";
		document.getElementsByName("appicon")[0].value="";
	}
}

function oncancel(){
//	window.returnValue=0;
	
	window.close();
}

function onok(){
//	window.returnValue=1;
//	document.forms["editapp"].submit();
//	window.close();
	var appname,url;
	appname=document.getElementsByName("appname")[0].value;
	url=document.getElementsByName("appurl")[0].value;
	if(appname=="" || appname==null){
//		alert("1");
		alert("应用名称为空");
		return ;
	}
	if(url=="" || url==null){
		alert("下载地址为空");
		return ;
	}
	var pwin = window.opener;
	document.forms['addapp'].submit();
//	addapp.submit();
//	document.forms["addapp"].submit();
	pwin.appmgr.retval.value="ok";
	//pwin.ReloadList();
	return; 
}

</script>
</head>
<body>
<form name="addapp" action="/appmgr/addapp" method="post">
<!--onsubmit="return onok()"> -->
<font style="font-size:20px;vertical-align: top;color: blue">应用详情</font>
</p>
编号&emsp;<input type="text" name="appid" id="appid" readonly /></p>
名称&emsp;<input type="text" name="appname" value={{.APPNAME}} /></p>
链接&emsp;<input type="text" name="appurl" value={{.APPURL}} /></p>
图标&emsp;<input type="text" name="appicon" value={{.APPICON}} /></p>
状态&emsp;<select name="status">
<option value="online" {{.SELON}} >上线</option>
<option value="offline" {{.SELOFF}} >下线</option>
</select>
</p>
<div style="text-align:center">
<input type="button" value="取 消" onclick="oncancel()" />&emsp;<input type="button" value="确 定" onclick="onok()" />
</div>
<input type="hidden" id="editid" name="editid" />
</form>
</body>
</html>
