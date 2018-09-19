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
}

function oncancel(){
	window.returnValue=0;
	window.close();
}

function onok(){
	window.returnValue=1;
	document.forms["editapp"].submit();
	window.close();
}

</script>
</head>
<body>
<form name="editapp" action="/appmgr/editapp" method="post">
<font style="font-size:20px;vertical-align: top;color: blue">应用详情</font>
</p>
编号&emsp;<input type="text" name="appid" id="appid" readonly>
名称&emsp;<input type="text" name="appname">
链接&emsp;<input type="text" name="appurl">
图标&emsp;<input type="text" name="appicon">
状态&emsp;<select>
<option value="online" selected>上线</option>
<option value="offline">下线</option>
</select>
</p>
<div style="text-align:center">
<input type="button" value="取 消" onclick="oncancel()" />&emsp;<input type="button" value="确 定" onclick="onok()" />
</div>
<input type="hidden" id="editid" name="editid" />
</form>
</body>
</html>
