-- container内で実行(今の所)
CREATE DATABASE motting default character set 'utf8';
create user 'motting' identified by 'motting';
grant all on motting.* to 'motting';

CREATE DATABASE webpush default character set 'utf8';
create user 'webpush' identified by 'webpush';
grant all on webpush.* to 'webpush';
