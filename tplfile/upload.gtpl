<html>
<head>
<title>上傳文件</title>
</head>

<body>
<form enctype="multipart/form-data" action="http://127.0.0.1:8080/upload" method="post">
  <input type="file" name="uploadfile" />
  <br><br><br>
  <input type="hidden" name="token" value="{{.}}">
<input type="submit" value="upload">
</form>

</body>
</html>