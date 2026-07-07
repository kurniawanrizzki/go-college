--
-- PostgreSQL database dump
--

\restrict P7kxcJk5Vk0oXvfaPW0IzO0nE5Z3ppp5YU8aZ1xc33Izaab8C1OQAjKac0ZePGn

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
-- Name: college; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.college (
    nim character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    semester integer NOT NULL,
    sks integer NOT NULL,
    active boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.college OWNER TO postgres;

--
-- Name: college college_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.college
    ADD CONSTRAINT college_pkey PRIMARY KEY (nim);


--
-- Name: idx_college_nim; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_college_nim ON public.college USING btree (nim) WITH (deduplicate_items='true');


--
-- PostgreSQL database dump complete
--

\unrestrict P7kxcJk5Vk0oXvfaPW0IzO0nE5Z3ppp5YU8aZ1xc33Izaab8C1OQAjKac0ZePGn

