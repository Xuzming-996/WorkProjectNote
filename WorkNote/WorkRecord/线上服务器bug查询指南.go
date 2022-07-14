金币项目通过JumpSever堡垒机进行统一管理服务器机器
1. 根据具体的某个平台找到对应其的server中心服的机器
2. 进入data/s...下查看operateconf.conf文件内的mysql密码
3. 登陆mysql，选择使用h1opt中心服表。通过mysql查询指令查找具体服务的server_open列，看是否有合服记录，没有就是id对应的游戏服，指令select * from server_open where serverid = 601;
4. 然后进入跳板机，使用autossh 平台名 服务器名 命令，会跳转到对应所在的的机器服务器
5. 进入data/s....路径再进入对应的服路径
6. cd log
7. 需要下载日志就用sz
8. data/backup/gamedb/下有数据库数据的备份

跳板机跳指定服务器 autossh 平台 区服id
exit 退出当前服务器
select actorname from actor where actorname='司琪华'; //查找玩家是否在该服