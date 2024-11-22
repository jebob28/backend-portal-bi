package dataops

import (
	bd_portal_bi "api-bi/bd_bi"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// funcao Ataulizar o setor
func UpdateSetor(id_setor int, nome_setor string) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 dbname=portal_bi host=192.168.1.65 sslmode=disable ")
	if err != nil {
		log.Fatal(err)
		return "erro no update", err
	}
	defer conn.Close(ctx)
	queries := bd_portal_bi.New(conn)

	parametros := bd_portal_bi.UpdateSetorParams{
		ID:    int32(id_setor),
		Setor: nome_setor}
	err = queries.UpdateSetor(ctx, parametros)
	if err != nil {
		fmt.Println("erro ao atualizar setor")
		return "", fmt.Errorf("erro ao atualizar o setor: %v", err)
	}

	return "Setor atualizado com sucesso !!", nil

}

// funcao atualizar usuario
func AtualizarUser(nome_user string, id_setor, id int) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 dbname=portal_bi host=192.168.1.65 sslmode=disable ")
	if err != nil {
		log.Fatal(err)
		return "erro no update", err
	}
	defer conn.Close(ctx)
	queries := bd_portal_bi.New(conn)

	parametros := bd_portal_bi.UpdateUsuarioParams{
		Nome:    pgtype.Text{String: nome_user, Valid: true},
		IDSetor: pgtype.Int4{Int32: int32(id_setor), Valid: true},
		ID:      int32(id),
	}

	err = queries.UpdateUsuario(ctx, parametros)
	if err != nil {
		fmt.Println("erro ao atualizar usuario")
		return "", fmt.Errorf("erro ao atualizar o usuario: %v", err)
	}
	return "Usuario atualizado com sucesso !!", nil
}

// funcao atualizar atualizar o link do power-bi
func AtualizarBi(id_setor, id int, link string) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 dbname=portal_bi host=192.168.1.65 sslmode=disable ")
	if err != nil {
		log.Fatal(err)
		return "erro no update", err
	}
	defer conn.Close(ctx)
	queries := bd_portal_bi.New(conn)

	parametros := bd_portal_bi.UpdateBiParams{
		ID:      int32(id),
		IDSetor: pgtype.Int4{Int32: int32(id_setor), Valid: true},
		Link:    pgtype.Text{String: link, Valid: true},
	}

	err = queries.UpdateBi(ctx, parametros)
	if err != nil {
		fmt.Println("erro ao atualizar Bi")
		return "", fmt.Errorf("erro ao atualizar o Bi: %v", err)
	}

	return "Bi atualizado com sucesso !!", nil

}
