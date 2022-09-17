CREATE TABLE `xz_auto_server_conf` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`group_zone` varchar(32) NOT NULL COMMENT '大区例如：wanba,changan,aiweiyou,360',
`server_id` int(11) DEFAULT '0' COMMENT '区服id',
`server_name` varchar(255) NOT NULL COMMENT '区服名称',
`open_time` varchar(64) DEFAULT NULL COMMENT '开服时间',
`service` varchar(30) DEFAULT NULL COMMENT '环境,test测试服,formal混服,wb玩吧',
`username` varchar(100) DEFAULT NULL COMMENT 'data管理员名称',
`submit_date` datetime DEFAULT NULL COMMENT '记录提交时间',
`status` tinyint(2) DEFAULT '0' COMMENT '状态,0未处理,1已处理,默认为0',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
