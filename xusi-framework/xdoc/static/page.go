// Copyright 2019 Xusixxxx Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package static

const PAGE_DOC = `
<!DOCTYPE html>
<html lang="ch-zn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Xusi General Development Suite</title>
	<link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
	<link rel="icon" href="/favicon.ico" type="image/x-icon">
    <style rel="stylesheet" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
		` + BOOTSTRAP_CSS + `
	</style>
	<script>
		` + JQUERY + `
	</script>
    <script integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous">
		` + BOOTSTRAP_JS + `
	</script>
    <style type="text/css">
        .header{
			z-index: 999;
            position: fixed;
            width: 100%;
            height: 50px;
			background-color: #FFFFFF;
			-moz-box-shadow:2px 2px 5px #B5B5B5; -webkit-box-shadow:2px 2px 5px #B5B5B5; box-shadow:2px 2px 5px #B5B5B5;
        }
		.header li{
			float:left;
			line-height:50px;
		}
		.header .logo{
			margin-right: 130px;'
		}
		.header .i-li{
			min-width:140px;
		}
        .content{
            margin-top: 50px;
        }
        .left{
			padding-top: 40px;
            float: left;
            width: 20%;
			position: fixed;
			overflow: auto;
			height:calc(100% - 50px);
			
        }
		.scroll-bar::-webkit-scrollbar {/*滚动条整体样式*/
       		width: 5px;     /*高宽分别对应横竖滚动条的尺寸*/
        	height: 1px;
		}
		.scroll-bar::-webkit-scrollbar-thumb {/*滚动条里面小方块*/
			border-radius: 10px;
			 -webkit-box-shadow: inset 0 0 5px rgba(0,0,0,0.2);
			background: #535353;
		}
		.scroll-bar::-webkit-scrollbar-track {/*滚动条里面轨道*/
			-webkit-box-shadow: inset 0 0 5px rgba(0,0,0,0.2);
			border-radius: 10px;
			background: #EDEDED;
		}
        .right{
			padding-left: 20px;
			margin-left: 20%;
			height: auto;
            float: left;
            width: 80%;
        }
		.keyword{
			color: #006633;
		}
		.func-name{
			color: #9966FF;
		}
		.type-name{
			color: #9966FF;
		}
    </style>
</head>
<body class="scroll-bar">
<div class="container-fluid">
    <div class="row header">
		<ul style="list-style:none;">
			<li class="logo"><a href="/"><img src="/logo" height="50px" /></a></li>
			<li class="i-li"><a href="/">开发文档</a></li>
			<li class="i-li"><a href="/api">API文档</a></li>
		</ul>
    </div>
    <div class="row content">
        <div class="left scroll-bar">{menu-left}</div>
        <div class="right">
			<div class="page-header">
				<h1>⚡️Xusi General Development Suite</h1></br>
			</div>
			<blockquote>
				<p>xusi-framework</p>
				<p>Xusi框架项目开发文档，文档由xdoc生成，随代码更新而更新。</p>
			</blockquote>
			<blockquote>
			  <p>目录</p>
			  <footer>xusi-framework开发内容清单</footer>
			</blockquote>
			<blockquote>
				{menu-content}
			</blockquote>
		</div>
    </div>
</div>
</body>
<script>
$(function () {
  $('.param-tip').tooltip()
})
</script>
</html>
`
