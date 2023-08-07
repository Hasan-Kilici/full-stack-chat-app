-- Database: ChatAPP

-- DROP DATABASE IF EXISTS "ChatAPP";

CREATE DATABASE "ChatAPP"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Turkish_Turkey.1254'
    LC_CTYPE = 'Turkish_Turkey.1254'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

BEGIN;


CREATE TABLE IF NOT EXISTS public.rooms
(
    token text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    ownerid text COLLATE pg_catalog."default" NOT NULL
);

CREATE TABLE IF NOT EXISTS public.users
(
    token text COLLATE pg_catalog."default" NOT NULL,
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    perm text COLLATE pg_catalog."default" NOT NULL
);
END;
