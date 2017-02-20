CREATE TABLE `msgcenter_innodb` (
  `msgid` bigint(20) NOT NULL AUTO_INCREMENT,
  `type` varchar(32) NOT NULL DEFAULT '',
  `initiatorid` bigint(20) NOT NULL DEFAULT 0,
  `initiatorname` varchar(255) NOT NULL DEFAULT '',
  `initiatorportrait` varchar(255) NOT NULL DEFAULT '',
  `consumerid` bigint(20) NOT NULL DEFAULT 0,
  `resource_id` varchar(255) DEFAULT '',
  `extra_info1` varchar(512) DEFAULT '',
  `extra_info2` varchar(512) DEFAULT '',
  `extra_info3` varchar(512) DEFAULT '',
  `extra_info4` varchar(512) DEFAULT '',
  `insert_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`msgid`,`consumerid`),
  KEY `idx_cid_itime` (`consumerid`,`insert_time`)
  KEY `idx_cid_initiator` (`consumerid`,`initiatorid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
PARTITION BY LINEAR HASH(`consumerid`)
PARTITIONS 128;
