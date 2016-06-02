# moilicms
基于Beego的内容管理系统(CMS)
待支持功能:
一期目标:
1.生成静态HTML,减少服务器压力.
2.支持多语言切换
3.支持离线发布内容
4.支持多说评论
5.支持七牛云存储静态文件
二期目标:
6.支持微信分享到功能.
7.支持二维码生成,并显示文章内容.
10.文章内容转发统计.
11.支持微信公众号绑定.
12.支持微信关键词回复,获得文章数字列表.
13.支持文章页面底部广告推荐.
14.支持发布到微信公众号
15.采集微信文章

http://localhost:8000
帐号：admin
密码：admin888
moilicms (MOLI CMS)魔力内容管理系统


看门狗脚本开机自动启动
vi /etc/rc.local 
在文件最后增加一行
nohup /usr/local/molicms/watchdogs.run 2>&1 >> info.log 2>&1 /dev/null &
这样子,在开机的时候就能自动启动看门狗脚本了.程序在也不怕挂掉了.
watchdogs.sh  看门狗脚本
run.sh 主程序启动脚本
setup.sh 程序安装脚本
pack.sh 程序打包脚本