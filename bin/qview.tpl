<html>
<head>
<title></title>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
</head>
<body>
<form action="/quickview" method="post">
起始日期:<input type="text" name="from" value="{{.From}}">  结束日期:<input type="text" name="to" value="{{.To}}"><br>
门店名称:<input type="text" name="store" value="{{.Store}}">  应用名称:<input type="text" name="app" value="{{.App}}"> <input type="submit" value="搜索"><br>共 {{.Total}} 条下载记录<br>
</form>
</body>
</html>
