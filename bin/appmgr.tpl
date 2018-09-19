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
	var s=showModalDialog("/appmgr/editapp",id,"dialogwidth:260px;dialogheight:240px");
	if (s==1){
		window.location.reload(true); 
	}
}
</script>
</head>
<body>
<pre>
<form name="appmgr" action="/appmgr" method="post">
<font style="font-size:20px;vertical-align: top;color: blue">应用管理</font>
</p>
{{AppEdit}}
</p>
<input type="hidden" name="editid" />
</form>
</pre>
</body>
</html>
