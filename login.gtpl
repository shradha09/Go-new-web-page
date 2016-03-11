<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    Username:<input type="text" name="username">
    Password:<input type="password" name="password">
    
    <input type="Submit" value="submit"  onclick="myfunction()">
</form>
<script>
function myfunction() {
document.getElementById("login").submit();
}
</script>
</body>
</html>

