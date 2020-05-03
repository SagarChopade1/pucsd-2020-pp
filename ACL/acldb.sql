CREATE DATABASE IF NOT EXIST acl;

DROP TABLE  IF NOT EXIST users;
DROP TABLE  IF NOT EXIST groups;
DROP TABLE  IF NOT EXIST whoGroupCreated;
DROP TABLE  IF NOT EXIST userGroupMap;
DROP TABLE IF NOT EXIST Content;
DROP TABLE IF NOT EXIST permission;
DROP TABLE  IF NOT EXIST session

CREATE TABLE IF NOT EXIST users
(
    uName VARCHAR(30) NOT NULL,
    userId  VARCHAR(30) NOT NULL,
    password  VARCHAR NOT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    userType VARCHAR(30) NOT NULL,
    PRIMARY KEY (userId)
);
INSERT INTO
CREATE TABLE IF NOT EXIST session
(
    userId VARCHAR(30) NOT NULL,
    sessionId VARCHAR NOT NULL,
    logintime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    PRIMARY KEY (userId,sessionId)
);

CREATE TABLE IF NOT EXIST groups
(
    groupId  INT NOT NULL AUTO_INCREMENT,
    groupName  VARCHAR(30) NOT NULL,
    groupCreationDate  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    groupDescription VARCHAR,
    PRIMARY KEY (groupId)
);


CREATE TABLE IF NOT EXIST whoGroupCreated
(
    userId VARCHAR(30),
    groupId INT,    
    FOREIGN KEY (userId) REFERENCES users(userId),
    FOREIGN KEY (groupId) REFERENCES groups(groupId),
    PRIMARY KEY(groupId,userId)
);

CREATE TABLE IF NOT EXIST userGroupMap
(
    groupId  INT,
    userId  VARCHAR(30) NOT NULL,
    FOREIGN KEY (groupId) REFERENCES groups(groupId),
    FOREIGN KEY (userId) REFERENCES users(userId),
    PRIMARY KEY(groupId,userId)
);

CREATE TABLE IF NOT EXISTS content(
    contentId int PRIMARY KEY
    contentName varchar(20),
    contentInfo varchar(2)
);

CREATE TABLE IF NOT EXISTS permission(
     permissionValue varchar(20) PRIMARY KEY
 );

CREATE TABLE IF NOT EXISTS userPermission(
    userId int PRIMARY KEY,
    contentId int ,
    permissionValue varchar(20) PRIMARY KEY
    FOREIGN KEY (userId) REFERENCES user(userId),
    FOREIGN KEY (contentId) REFERENCES content(contentId),
    FOREIGN KEY (permissionValue) REFERENCES permission(permissionValue),
);

CREATE TABLE IF NOT EXISTS groupsPermission(
    groupId int PRIMARY KEY,
    contentId int ,
    permissionValue varchar(20) PRIMARY KEY
    FOREIGN KEY (groupId) REFERENCES groups(groupId),
    FOREIGN KEY (contentId) REFERENCES content(contentId),
    FOREIGN KEY (permissionValue) REFERENCES permission(permissionValue),
);

