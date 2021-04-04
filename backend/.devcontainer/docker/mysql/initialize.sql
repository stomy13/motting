ALTER DATABASE motting default character set 'utf8';
GRANT ALL ON mysql.proc to 'motting'@'%';
GRANT ALL ON *.* to 'motting'@'%';
SET GLOBAL log_bin_trust_function_creators=1;

-- CREATE DATABASE motting default character set 'utf8';
-- create user 'motting' identified by 'motting';
-- grant all on motting.* to 'motting';

CREATE DATABASE webpush default character set 'utf8';
-- create user 'webpush' identified by 'webpush';
-- grant all on webpush.* to 'webpush';
