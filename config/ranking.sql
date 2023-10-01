# 活动表

CREATE TABLE
    `activity` (
        `id` int(11) NOT NULL AUTO_INCREMENT,
        `name` varchar(100) COLLATE utf8mb4_bin NOT NULL,
        `addTime` int(11) DEFAULT 0 NOT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;

# 参赛选手表

CREATE TABLE
    `player` (
        `id` int(11) NOT NULL AUTO_INCREMENT,
        `aid` int(11) NOT NULL COMMENT "所属活动",
        `ref` varchar(50) DEFAULT "" NOT NULL COMMENT "编号",
        `nickName` varchar(100) DEFAULT "" NOT NULL COMMENT "昵称",
        `declaration` varchar(255) DEFAULT "" NOT NULL COMMENT "宣言",
        `avatar` varchar(255) DEFAULT "" NOT NULL COMMENT "头像",
        `score` int(11) DEFAULT 0 NOT NULL COMMENT "分数",
        `addTime` int(11) DEFAULT 0 NOT NULL COMMENT "添加时间",
        `updateTime` int(11) DEFAULT 0 NOT NULL COMMENT "更新时间",
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;

# 用户表

CREATE TABLE
    `user` (
        `id` int(11) NOT NULL AUTO_INCREMENT,
        `username` varchar(50) NOT NULL COMMENT "用户名",
        `password` varchar(50) DEFAULT "" NOT NULL COMMENT "密码",
        `addTime` int(11) DEFAULT 0 NOT NULL COMMENT "添加时间",
        `updateTime` int(11) DEFAULT 0 NOT NULL COMMENT "更新时间",
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;

# 点赞表

CREATE TABLE
    `vote` (
        `id` int(11) NOT NULL AUTO_INCREMENT,
        `userId` int(11) NOT NULL COMMENT "投票用户id",
        `playerId` int(11) NOT NULL COMMENT "选手id",
        `addTime` int(11) DEFAULT 0 NOT NULL COMMENT "添加时间",
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;