package dataops

import (
	bd_portal_bi "api-bi/bd_bi"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Funcao deletar setor
func DeletarSetor(id int) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 host=192.168.1.65 dbname=portal_bi sslmode=disable")
	if err != nil {
		return "Erro ao deletar", err
	}
	defer conn.Close(ctx)
	queries := bd_portal_bi.New(conn)
	err = queries.DeleteSetor(ctx, int32(id))
	if err != nil {
		fmt.Println("erro ao deletar setor")
		return "Erro ao deletar", fmt.Errorf("erro ao deletar setor: %v", err)
	}
	return "Setor deletado com sucesso", nil

}

// Funcao deletar usuario
func DeletarUsuario(id int) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 host=192.168.1.65 dbname=portal_bi sslmode=disable")
	if err != nil {
		return "Erro ao deletar", err
	}
	defer conn.Close(ctx)
	queries := bd_portal_bi.New(conn)
	err = queries.DeleteUsuario(ctx, int32(id))
	if err != nil {
		fmt.Println("erro deletar usuario")
		return "", fmt.Errorf("erro ao deletar usuario: %v", err)
	}
	return "Usuario deletado com sucesso", nil

}

// funcao deletar link do BI
func DeletarPortalBi(id int) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 host=192.168.1.65 dbname=portal_bi sslmode=disable")
	if err != nil {
		return "Erro ao deletar", err

	}
	defer conn.Close(ctx)
	queries := bd_portal_bi.New(conn)
	err = queries.DeleteBi(ctx, int32(id))
	if err != nil {
		fmt.Println("erro ao deletar portal bi")
		return "Erro ao deletar", fmt.Errorf("erro ao deletar portal bi: %v", err)
	}
	return "Portal Bi deletado com sucesso", nil
}
