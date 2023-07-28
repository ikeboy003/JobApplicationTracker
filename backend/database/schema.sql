BEGIN TRANSACTION;

DROP TABLE IF EXISTS jobsapplied;
DROP SEQUENCE IF EXISTS job_id_seq;

CREATE SEQUENCE job_id_seq 
	START WITH 1000
	INCREMENT BY 1;
	
CREATE TABLE jobsapplied (
	job_id int DEFAULT nextval('job_id_seq') PRIMARY KEY,
	job_name VARCHAR(200) NOT NULL,
	job_company VARCHAR(200) NOT NULL,
	app_status VARCHAR(20) CHECK (app_status IN ('applied', 'rejected', 'interviewed'))
);
COMMIT;