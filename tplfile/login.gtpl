<html>
<head>
<title>test</title>
</head>

<body>
<form action="http://127.0.0.1:8080/login" method="post">
  用戶名:<input type="text" name="username">
  <br>
  密    碼:<input type="password" name="password">
  <br><br><br>
  <input type="radio" name="gender" value="1">男</input>
  <input type="radio" name="gender" value="2">女</input>
  <br><br><br>
  <input type="checkbox" name="interest" value="program">編程</input>
  <input type="checkbox" name="interest" value="game">游戲</input>
  <input type="checkbox" name="interest" value="football">足球</input>
  <br><br><br>
  <input type="hidden" name="token" value="{{.}}">
<input type="submit" value="login">
</form>

</body>
</html>