<html>
<head>
<title></title>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
</head>
<body>
<form action="/quickview" method="post">
起始日期:&emsp;&emsp;<input type="text" name="from" value="{{.From}}"><br>结束日期:&emsp;&emsp;<input type="text" name="to" value="{{.To}}"><br>
门店关键字:&emsp;<input type="text" name="store" value="{{.Store}}"><br>应用关键字:&emsp;<input type="text" name="app" value="{{.App}}"> <input type="submit" value="搜索"><br>共 {{.Total}} 条下载记录<br>
</form>
</body>
</html>
