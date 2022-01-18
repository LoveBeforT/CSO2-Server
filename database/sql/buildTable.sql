-- membercode - memberconfirmation - member code and verification code for the benefit of users -
Create Table If Not Exists UserData(
    uid INT(10) NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL DEFAULT 'unkown_user',
    gamename VARCHAR(64) NOT NULL DEFAULT 'unkown_user',
	membercode VARCHAR(64) NULL DEFAULT '0',
	memberconfirmation VARCHAR(64) NULL,
    password VARCHAR(64) NOT NULL DEFAULT 'unkown_passwd',
    mail VARCHAR(64) NULL DEFAULT 'unkown_mail',
    data text,
    PRIMARY KEY(uid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;