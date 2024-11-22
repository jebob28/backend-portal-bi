CREATE TABLE public.bi (
	id serial4 NOT NULL,
	link varchar(255) NULL,
	id_setor int4 NULL,
	CONSTRAINT bi_pkey PRIMARY KEY (id)
);

CREATE TABLE public.setor (
	id serial4 NOT NULL,
	setor varchar(100) NOT NULL,
	CONSTRAINT setor_pkey PRIMARY KEY (id)
);

CREATE TABLE public.usuario (
    id serial4 NOT NULL,
    nome varchar(255) NULL,
    id_setor int4 NULL,
    id_permisao int4 NULL, -- Coluna adicionada para chave estrangeira
    email text NULL,
    login varchar(100) NULL,
    senha varchar(100) NULL,
    CONSTRAINT usuario_pkey PRIMARY KEY (id),
    CONSTRAINT usuario_id_setor_fkey FOREIGN KEY (id_setor) REFERENCES public.setor(id),
    CONSTRAINT fk_usuario_permisao FOREIGN KEY (id_permisao) REFERENCES public.permisao(id)
);

CREATE TABLE public.permisao (
	id serial4 NOT NULL,
	nome varchar(100) NULL,
	CONSTRAINT permisao_pkey PRIMARY KEY (id)
);
