CREATE DATABASE
IF
    NOT EXISTS blog_service DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

use blog_service;

CREATE TABLE `blog_tag` (
    `id` INT(16) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) DEFAULT '' COMMENT '标签名',

    `created_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '创建时间',
    `created_by` VARCHAR(10) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '删除时间',
    `is_deleted` TINYINT(16) UNSIGNED DEFAULT 0 COMMENT '是否删除 0 未删除 1 已删除',

    `state` TINYINT(4) UNSIGNED DEFAULT 0 COMMENT '状态 0 启用， 1 禁用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

CREATE TABLE `blog_article` (
    `id` INT(16) UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
    `desc` VARCHAR(255) DEFAULT '' COMMENT '文章简述',
    `cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面图片地址',
    `content` LONGTEXT DEFAULT '' COMMENT '文章内容',

    `created_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '创建时间',
    `created_by` VARCHAR(10) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '删除时间',
    `is_deleted` TINYINT(16) UNSIGNED DEFAULT 0 COMMENT '是否删除 0 未删除 1 已删除',
    `state` TINYINT(4) UNSIGNED DEFAULT 0 COMMENT '状态 0 启用， 1 禁用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

CREATE TABLE `blob_article_tag` (
    `id` INT(16) UNSIGNED NOT NULL AUTO_INCREMENT,
    `article_id` INT(16) UNSIGNED NOT NULL COMMENT '文章 ID',
    `tag_id` INT(16) UNSIGNED NOT NULL COMMENT '标签 ID',

    `created_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '创建时间',
    `created_by` VARCHAR(10) DEFAULT '' COMMENT '创建人',
    `modified_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '修改时间',
    `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` INT(16) UNSIGNED DEFAULT 0 COMMENT '删除时间',
    `is_deleted` TINYINT(16) UNSIGNED DEFAULT 0 COMMENT '是否删除 0 未删除 1 已删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';