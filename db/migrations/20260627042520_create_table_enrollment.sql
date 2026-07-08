--
-- PostgreSQL database dump
--

\restrict vNmUexbuc9ckc4Ieg8x9WCGhYTdiMUulpKSdHFmIufZMJ4rQMdQY0N9k7PGeEzQ

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
-- Name: enrollment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.enrollment (
    id integer NOT NULL,
    nim character varying(255) NOT NULL,
    course_code character varying(255) NOT NULL,
    semester integer NOT NULL,
    grade "char" DEFAULT 'E'::"char" NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    udpated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.enrollment OWNER TO postgres;

--
-- Name: enrollment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.enrollment ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.enrollment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: enrollment enrollment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.enrollment
    ADD CONSTRAINT enrollment_pkey PRIMARY KEY (id);


--
-- Name: enrollment unique_pair_constraint; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.enrollment
    ADD CONSTRAINT unique_pair_constraint UNIQUE (nim, course_code);


--
-- Name: idx_enrollment_nim; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_enrollment_nim ON public.enrollment USING btree (nim) WITH (deduplicate_items='true');


--
-- Name: enrollment college_course_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.enrollment
    ADD CONSTRAINT college_course_fkey FOREIGN KEY (course_code) REFERENCES public.course(code) ON DELETE CASCADE NOT VALID;


--
-- Name: enrollment college_nim_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.enrollment
    ADD CONSTRAINT college_nim_fkey FOREIGN KEY (nim) REFERENCES public.college(nim) NOT VALID;


--
-- PostgreSQL database dump complete
--

\unrestrict vNmUexbuc9ckc4Ieg8x9WCGhYTdiMUulpKSdHFmIufZMJ4rQMdQY0N9k7PGeEzQ

