CREATE USER test_user PASSWORD 'postgrespw';

CREATE DATABASE course_manager;
GRANT ALL PRIVILEGES ON DATABASE course_manager TO test_user;

CREATE DATABASE course_manager_test;
GRANT ALL PRIVILEGES ON DATABASE course_manager_test TO test_user;
