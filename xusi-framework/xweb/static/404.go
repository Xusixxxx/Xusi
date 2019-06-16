package static

const PAGE_404 string = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta charset="UTF-8" http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>Xusi Framework xweb</title>
    <style type="text/css">
        .content{
            padding-top: 50px;
            position: relative;
            margin: auto;
            width: 100%;
            height: 100%;
            overflow: hidden;
            text-align: center;
        }
        h1{
            color: midnightblue;
        }
        h2{
            color: firebrick;
        }
        h5{
            padding-top: 70px;
            color: dimgray;
        }
        h5 a{
            margin: 10px;
        }
		.anim{
            position:fixed;
            animation:anim 5s infinite;
            -moz-animation:anim 5s infinite; /* Firefox */
            -webkit-animation:anim 5s infinite; /* Safari and Chrome */
            -o-animation:anim 5s infinite; /* Opera */
        }
        @keyframes anim
        {
            0%   {left:-20px;width:10px; height:10px; background:midnightblue; }
            50%  {left:100%;width:10px; height:10px; background:midnightblue;border-radius:100%; }
            100% {left:-20px;width:10px; height:10px; background:midnightblue; }
        }
        @-moz-keyframes anim /* Firefox */
        {
            0%   {left:0px;width:0px; height:0px; background:midnightblue; }
            50% {left:100%;width:20px; height:20px; background:midnightblue;border-radius:100%; }
            100%   {left:0px;width:0px; height:0px; background:midnightblue; }
        }

        @-webkit-keyframes anim /* Safari 和 Chrome */
        {
            0%   {left:0px;width:0px; height:0px; background:midnightblue; }
            50% {left:100%;width:20px; height:20px; background:midnightblue;border-radius:100%; }
            100%   {left:0px;width:0px; height:0px; background:midnightblue; }
        }

        @-o-keyframes anim /* Opera */
        {
            0%   {left:0px;width:0px; height:0px; background:midnightblue; }
            50% {left:100%;width:20px; height:20px; background:midnightblue;border-radius:100%; }
            100%   {left:0px;width:0px; height:0px; background:midnightblue; }
        }
    </style>
</head>
<body>
<div class="content">
    <img src="https://s2.ax1x.com/2019/06/16/VTbJ1S.png" />
    <h1>Sorry, not found page!</h1>
    <h2>error code : 404</h2>
	<div class="anim"></div>
    <h5>
        <a href="https://github.com/Xusixxxx/Xusi">Github</a>
        <a href="https://gitee.com/Xusixxxx/Xusi">Gitee</a>
        <p style="font-size: 12px">
            <a href="xusixxxx@gmail.com">xusixxxx@gmail.com</a>
        </p>
    </h5>
</div>

</body>
</html>

`
