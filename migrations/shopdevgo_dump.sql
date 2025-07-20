--
-- PostgreSQL database dump
--

-- Dumped from database version 14.18 (Ubuntu 14.18-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.18 (Ubuntu 14.18-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: crm_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.crm_users (
    id integer NOT NULL,
    email character varying(30) DEFAULT ''::character varying NOT NULL,
    phone character varying(15) DEFAULT ''::character varying NOT NULL,
    username character varying(30) DEFAULT ''::character varying NOT NULL,
    password character varying(32) DEFAULT ''::character varying NOT NULL,
    created_at integer DEFAULT 0,
    updated_at integer DEFAULT 0,
    create_ip character varying(15) DEFAULT ''::character varying NOT NULL,
    last_login_at integer,
    last_login_ip character varying(15) DEFAULT NULL::character varying,
    login_time integer DEFAULT 0 NOT NULL,
    user_status smallint DEFAULT 0 NOT NULL
);


ALTER TABLE public.crm_users OWNER TO postgres;

--
-- Name: crm_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.crm_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.crm_users_id_seq OWNER TO postgres;

--
-- Name: crm_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.crm_users_id_seq OWNED BY public.crm_users.id;


--
-- Name: crm_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.crm_users ALTER COLUMN id SET DEFAULT nextval('public.crm_users_id_seq'::regclass);


--
-- Data for Name: crm_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.crm_users (id, email, phone, username, password, created_at, updated_at, create_ip, last_login_at, last_login_ip, login_time, user_status) FROM stdin;
\.


--
-- Name: crm_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.crm_users_id_seq', 1, false);


--
-- Name: crm_users crm_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.crm_users
    ADD CONSTRAINT crm_users_pkey PRIMARY KEY (id);


--
-- Name: idx_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_email ON public.crm_users USING btree (email);


--
-- Name: idx_phone; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_phone ON public.crm_users USING btree (phone);


--
-- Name: idx_username; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_username ON public.crm_users USING btree (username);


--
-- PostgreSQL database dump complete
--

