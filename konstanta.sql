PGDMP      )    ,                 z         	   konstanta    14.3    14.3     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16430 	   konstanta    DATABASE     f   CREATE DATABASE konstanta WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';
    DROP DATABASE konstanta;
                postgres    false            �            1259    16432    transactions    TABLE     G  CREATE TABLE public.transactions (
    t_id integer NOT NULL,
    user_id integer NOT NULL,
    email character varying(255) NOT NULL,
    amount double precision NOT NULL,
    currency character varying(255) NOT NULL,
    create_time date NOT NULL,
    update_time date NOT NULL,
    status character varying(255) NOT NULL
);
     DROP TABLE public.transactions;
       public         heap    postgres    false            �            1259    16431    transactions_t_id_seq    SEQUENCE     �   CREATE SEQUENCE public.transactions_t_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.transactions_t_id_seq;
       public          postgres    false    210            �           0    0    transactions_t_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.transactions_t_id_seq OWNED BY public.transactions.t_id;
          public          postgres    false    209            \           2604    16435    transactions t_id    DEFAULT     v   ALTER TABLE ONLY public.transactions ALTER COLUMN t_id SET DEFAULT nextval('public.transactions_t_id_seq'::regclass);
 @   ALTER TABLE public.transactions ALTER COLUMN t_id DROP DEFAULT;
       public          postgres    false    210    209    210            �          0    16432    transactions 
   TABLE DATA           p   COPY public.transactions (t_id, user_id, email, amount, currency, create_time, update_time, status) FROM stdin;
    public          postgres    false    210   �       �           0    0    transactions_t_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.transactions_t_id_seq', 1, false);
          public          postgres    false    209            ^           2606    16439    transactions transactions_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (t_id);
 H   ALTER TABLE ONLY public.transactions DROP CONSTRAINT transactions_pkey;
       public            postgres    false    210            �      x^����� � �     