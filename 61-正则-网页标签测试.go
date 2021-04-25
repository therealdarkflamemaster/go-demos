package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="Content-Type" content="text/html;charset=utf-8">
<meta name="Referrer" content="unsafe-url">
<meta content="True" name="HandheldFriendly">
<meta name="theme-color" content="#333344">
<meta name="apple-mobile-web-app-capable" content="yes" />
<meta name="mobile-web-app-capable" content="yes" />
<meta name="detectify-verification" content="d0264f228155c7a1f72c3d91c17ce8fb" />
<meta name="p:domain_verify" content="b87e3b55b409494aab88c1610b05a5f0" />
<meta name="alexaVerifyID" content="OFc8dmwZo7ttU4UCnDh1rKDtLlY" />
<meta name="baidu-site-verification" content="D00WizvYyr" />
<meta name="msvalidate.01" content="D9B08FEA08E3DA402BF07ABAB61D77DE" />
<meta property="wb:webmaster" content="f2f4cb229bda06a4" />
<meta name="google-site-verification" content="LM_cJR94XJIqcYJeOCscGVMWdaRUvmyz6cVOqkFplaU" />
<title>V2EX</title>
<link rel="dns-prefetch" href="https://static.v2ex.com/" />
<link rel="dns-prefetch" href="https://cdn.v2ex.com/" />
<link rel="dns-prefetch" href="https://i.v2ex.co/" />
<link rel="dns-prefetch" href="https://www.google-analytics.com/" />
<link rel="stylesheet" type="text/css" media="screen" href="/css/basic.css?v=542005:1618090207:3.9.8.5">
<link rel="stylesheet" type="text/css" media="screen" href="/assets/afe358d048b97e4ae75f3f098200e42346ce6cf4-combo.css">
<link rel="stylesheet" type="text/css" media="screen" href="/css/desktop.css?v=3.9.8.5">
<script>
        const SITE_NIGHT = 0;
    </script>
<link rel="stylesheet" href="/static/css/tomorrow.css?v=3c006808236080a5d98ba4e64b8f323f" type="text/css">
<link rel="icon" sizes="192x192" href="/static/icon-192.png">
<link rel="apple-touch-icon" sizes="180x180" href="/static/apple-touch-icon-180.png?v=91e795b8b5d9e2cbf2d886c3d4b7d63c">
<link rel="shortcut icon" href="/static/favicon.ico" type="image/png">
<link rel="manifest" href="/manifest.webmanifest">
<script>
        const FEATURES = ['search', 'favorite-nodes-sort'];
    </script>
<script src="/assets/39beba1a1982aebe111156eb00777d395f54ba02-combo.js" defer></script>
<meta name="description" content="创意工作者的社区。讨论编程、设计、硬件、游戏等令人激动的话题。">
<link rel="alternate" type="application/atom+xml" href="/index.xml">
<link rel="canonical" href="https://www.v2ex.com/">
<script type="text/javascript">
    document.addEventListener("DOMContentLoaded", function(event) {
        protectTraffic();

        if ('serviceWorker' in navigator) {
    window.addEventListener('load', () =>
        navigator.serviceWorker.register('/sw.js?v0')
            .catch(() => {}) // ignore
    );
}

        $("#nodes-sidebar").sortable();
        $("#nodes-sidebar").disableSelection();
        $("#nodes-sidebar").sortable({
            stop: function(event, ui) {
                let sorted = $("#nodes-sidebar").sortable( "serialize", {key: "n"});
                $.post('/my/nodes/sorted', {sorted: sorted}, function(data) {
        });
            }
        });

        

        const blocked = [];
        const ignored_topics = [];

        $("#TopicsHot").children('.cell').each(function(index) {
            for (i in blocked) {
                if ($(this).hasClass('from_' + blocked[i])) {
                    $(this).css('display', 'none');
                }
            }
            for (i in ignored_topics) {
                let css_class = 'hot_t_' + ignored_topics[i];
                if ($(this).hasClass(css_class)) {
                    $(this).css('display', 'none');
                }
            }
        });
        
    });
</script>
</head>
<title>
</title>
<div>hello regexp</div>
<div>hello 2</div>
<div>hello 3</div>
<div>hello 4</div>
<div>hello 5</div>
<div>
hello 6
换行
换行
</div>
<body>
身体
</body>
`

	//解析，编译正则表达式
	//ret :=regexp.MustCompile(`<div>(.*?)</div>`)
	// ret :=regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	ret := regexp.MustCompile(`<link rel="stylesheet" href="(?s:(.*?))"`)

	// 提取需要的信息
	alls := ret.FindAllStringSubmatch(str, -1)
	fmt.Println(alls)
}
