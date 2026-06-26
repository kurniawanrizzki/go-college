--
-- PostgreSQL database dump
--

\restrict 564YcL4pHA2LbGiBHJSIHsRfmEtpAfPlmfBUIz4ePW46lJUDrHB3IoHvA5GujE8

-- Dumped from database version 18.4 (Homebrew)
-- Dumped by pg_dump version 18.4 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: course; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.course (
    code character varying NOT NULL,
    name character varying(255) NOT NULL,
    sks integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.course OWNER TO postgres;

--
-- Name: course course_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course
    ADD CONSTRAINT course_pkey PRIMARY KEY (code);


--
-- PostgreSQL database dump complete
--

\unrestrict 564YcL4pHA2LbGiBHJSIHsRfmEtpAfPlmfBUIz4ePW46lJUDrHB3IoHvA5GujE8

