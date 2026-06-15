-- 建库（MySQL 8）
DROP DATABASE IF EXISTS interconnect_peer_pairing;
CREATE DATABASE interconnect_peer_pairing CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE interconnect_peer_pairing;

-- 建表

CREATE TABLE connections (
    connection_uuid VARCHAR(64)  PRIMARY KEY,
    from_device     VARCHAR(128) NOT NULL,
    from_port       VARCHAR(64)  NOT NULL,
    to_device       VARCHAR(128) NOT NULL,
    to_port         VARCHAR(64)  NOT NULL
) ENGINE=InnoDB;

CREATE TABLE interconnects (
    uuid            VARCHAR(64)  PRIMARY KEY,
    connection_uuid VARCHAR(64)  NOT NULL,
    device          VARCHAR(128) NOT NULL,
    ip_address      VARCHAR(64)  NOT NULL,
    CONSTRAINT fk_ic_conn FOREIGN KEY (connection_uuid)
        REFERENCES connections(connection_uuid)
) ENGINE=InnoDB;

-- 示例 1：单链路基本配对

INSERT INTO connections VALUES ('conn-1', 'S6_B', '100GE1/1/1', 'LF6_A', '100GE1/0/1');

INSERT INTO interconnects VALUES
    ('ic-1', 'conn-1', 'S6_B',  '172.16.1.1/30'),
    ('ic-2', 'conn-1', 'LF6_A', '172.16.1.2/30');

-- 示例 2：并行链路——同设备对之间有两根线

INSERT INTO connections VALUES
    ('conn-A', 'LF6_A', '100GE1/0/2', 'LF6_C', '100GE1/0/2'),
    ('conn-B', 'LF6_A', '100GE1/0/3', 'LF6_C', '100GE1/0/3');

INSERT INTO interconnects VALUES
    ('ic-A1', 'conn-A', 'LF6_A', '172.16.10.1/30'),
    ('ic-C1', 'conn-A', 'LF6_C', '172.16.10.2/30'),
    ('ic-A2', 'conn-B', 'LF6_A', '172.16.10.5/30'),
    ('ic-C2', 'conn-B', 'LF6_C', '172.16.10.6/30');

-- 示例 3：孤儿端——对端行缺失（DEV-B 那行故意不插入）

INSERT INTO connections VALUES ('conn-X', 'DEV-A', '10GE1/0/1', 'DEV-B', '10GE1/0/1');

INSERT INTO interconnects VALUES
    ('ic-A', 'conn-X', 'DEV-A', '192.168.0.1/30');
