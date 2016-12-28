--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: blocks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE blocks (
    block_number text NOT NULL,
    block_hash text,
    block_number_id bigint NOT NULL
);


ALTER TABLE blocks OWNER TO postgres;

--
-- Name: COLUMN blocks.block_number_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN blocks.block_number_id IS 'Makes lookups easier, as block_number is an hex text';


--
-- Name: logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE logs (
    id bigint NOT NULL,
    log_transaction_hash text,
    data text,
    log_index text,
    mined text
);


ALTER TABLE logs OWNER TO postgres;

--
-- Name: logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE logs_id_seq OWNER TO postgres;

--
-- Name: logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE logs_id_seq OWNED BY logs.id;


--
-- Name: topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE topics (
    log_id bigint,
    content text
);


ALTER TABLE topics OWNER TO postgres;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transactions (
    transaction_hash text NOT NULL,
    tx_block_number text,
    transaction_index text,
    tx_from text,
    tx_to text
);


ALTER TABLE transactions OWNER TO postgres;

--
-- Name: logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY logs ALTER COLUMN id SET DEFAULT nextval('logs_id_seq'::regclass);


--
-- Name: blocks blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY blocks
    ADD CONSTRAINT blocks_pkey PRIMARY KEY (block_number_id);


--
-- Name: logs logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (transaction_hash);


--
-- Name: block_hash_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX block_hash_idx ON blocks USING btree (block_hash);


--
-- Name: block_number_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX block_number_idx ON blocks USING btree (block_number);


--
-- Name: content_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX content_idx ON topics USING btree (content);


--
-- Name: log_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX log_id_idx ON topics USING btree (log_id);


--
-- Name: log_transaction_hash_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX log_transaction_hash_idx ON logs USING btree (log_transaction_hash);


--
-- Name: tx_block_number_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX tx_block_number_idx ON transactions USING btree (tx_block_number);


--
-- Name: tx_from_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX tx_from_idx ON transactions USING btree (tx_from);


--
-- Name: tx_to_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX tx_to_idx ON transactions USING btree (tx_to);


--
-- Name: topics log_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY topics
    ADD CONSTRAINT log_id_fk FOREIGN KEY (log_id) REFERENCES logs(id) MATCH FULL;


--
-- Name: logs log_transaction_hash_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY logs
    ADD CONSTRAINT log_transaction_hash_fk FOREIGN KEY (log_transaction_hash) REFERENCES transactions(transaction_hash) MATCH FULL;


--
-- Name: transactions tx_block_number_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transactions
    ADD CONSTRAINT tx_block_number_id_fk FOREIGN KEY (tx_block_number) REFERENCES blocks(block_number) MATCH FULL;


--
-- PostgreSQL database dump complete
--

