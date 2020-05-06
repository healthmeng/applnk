<html>
<head>
<title></title>
<meta charset="UTF-8"/>
<meta name="viewport" content= "width=device-width, initial-scale=1.0,maximum-scale=1.0,user-scalable=0" />
</head>
<body>
<form action="/quickview" method="post">
前缀:<input type="text" name="prefix" value="{{.Prefix}}" /><br>
后缀:<input type="text" name="suffix" value="{{.Suffix}}" /><br>
表格内容:<br>
<textarea name="excel" rows="15" cols="40">{{.Excel}}</textarea> <br>
<input type="submit" value="转换" />
</form>
</body>
</html>
