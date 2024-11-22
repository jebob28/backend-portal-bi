-- name: CreateUsuarios :exec 
INSERT INTO public.usuario
(nome, id_setor,email,login,senha,id_permisao)
VALUES($1,$2,$3,$4,$5,$6);

-- name: CreateSetor :exec
INSERT INTO public.setor
(setor)
VALUES($1);

-- name: CreateLink :exec
INSERT INTO public.bi
(link, id_setor)
VALUES($1, $2);

-- name: GetUsuario :many 
SELECT id, nome, id_setor,email,login,senha,id_permisao
FROM public.usuario;

-- name: GetSetor :many 
SELECT id, setor
FROM public.setor;

-- name: GetBi :one 
SELECT id, link, id_setor
FROM public.bi WHERE id_setor=$1; 

-- name: UpdateUsuario :exec 
UPDATE public.usuario
SET nome=$1, id_setor=$2
WHERE id=$3;

-- name: UpdateSetor :exec
UPDATE public.setor
SET setor=$2
WHERE id=$1;

-- name: UpdateBi :exec
UPDATE public.bi
SET link=$3, id_setor=$2
WHERE id=$1;

-- name: DeleteUsuario :exec
DELETE FROM public.usuario
WHERE id=$1;

-- name: DeleteSetor :exec
DELETE FROM public.setor
WHERE id=$1;

-- name: DeleteBi :exec
DELETE FROM public.bi
WHERE id=$1;