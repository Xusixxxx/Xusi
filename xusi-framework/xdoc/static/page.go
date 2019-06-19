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
    <title>⚡️Xusi General Development Suite</title>
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
            position: fixed;
            width: 100%;
            height: 50px;
        }
        .content{
            margin-top: 50px;
        }
        .left{
            min-height: 500px;
            float: left;
            width: 20%;
        }
        .right{
            min-height: 500px;
            float: left;
            width: 80%;
        }
    </style>
</head>
<body>
<div class="container-fluid">
    <div class="row header">
		<img src="/logo" height="50px" />
    </div>
    <div class="row content">
        <div class="left">{menu-left}</div>
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
			{content}
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
