CREATE USER 'wb_user'@'localhost' IDENTIFIED BY '031501';
GRANT ALL PRIVILEGES ON * . * TO 'wb_user'@'localhost';
FLUSH PRIVILEGES;

SET GLOBAL sql_mode='';