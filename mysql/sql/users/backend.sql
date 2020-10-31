CREATE USER 'safe-events-backend'@'%' IDENTIFIED BY 'backendpwd';
GRANT ALL ON 'safe-events-db'.* TO 'safe-events-backend'@'%';