<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>Upload</title>
</head>
<body>
    <form action="http://localhost:9090/upload" enctype="multipart/form-data" method="post">
        <input type="file" name="uploadfile">
        <input type="hidden" name="token" value="{{.}}">
        <input type="submit" value="upload">
    </form>
</body>
</html>