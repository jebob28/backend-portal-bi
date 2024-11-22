package dataops

import (
	bd_portal_bi "api-bi/bd_bi"
	"context"

	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func SelectSetor() ([]bd_portal_bi.Setor, error) {
	ctx := context.Background()

	// Configuração da conexão com o banco de dados
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 host=192.168.1.65 dbname=portal_bi sslmode=disable")
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}
	defer conn.Close(ctx)

	// Criação do objeto de consulta
	queries := bd_portal_bi.New(conn)

	// Busca pelos setores
	setores, err := queries.GetSetor(ctx)
	if err != nil {
		log.Println("Erro ao buscar setores:", err)
		return nil, err
	}

	// Retorna os setores
	return setores, nil
}

func SelectUser() ([]bd_portal_bi.GetUsuarioRow, error) {
	ctx := context.Background()
	// Configuração da conexão com o banco de dados
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 host=192.168.1.65 dbname=portal_bi sslmode=disable")
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}
	defer conn.Close(ctx)
	// Criação do objeto de consulta
	queries := bd_portal_bi.New(conn)
	// Busca pelos usuários
	usuarios, err := queries.GetUsuario(ctx)
	if err != nil {
		log.Println("Erro ao buscar usuários:", err)
		return nil, err
	}
	// Retorna os usuários
	return usuarios, nil
}

func SelectBi(id_link int) (bd_portal_bi.Bi, error) {
	ctx := context.Background()

	// Configuração da conexão com o banco de dados
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 host=192.168.1.65 dbname=portal_bi sslmode=disable")
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return bd_portal_bi.Bi{}, err
	}
	defer conn.Close(ctx)

	queries := bd_portal_bi.New(conn)

	// Criação do objeto idlink
	idlink := pgtype.Int4{Int32: int32(id_link), Valid: true}

	// Busca os dados do BI
	bi, err := queries.GetBi(ctx, idlink)
	if err != nil {
		log.Println("Erro ao buscar usuários:", err)
		return bd_portal_bi.Bi{}, err
	}

	return bi, nil
}
