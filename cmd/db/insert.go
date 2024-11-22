package dataops

import (
	bd_portal_bi "api-bi/bd_bi"
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// funcao cadastrar Setor
func CadastrarSetor(setor string) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 dbname=portal_bi host=192.168.1.65 sslmode=disable")
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return "", fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}
	defer conn.Close(ctx)

	// Execução da query para criar o setor

	queries := bd_portal_bi.New(conn)
	err = queries.CreateSetor(ctx, setor)
	if err != nil {
		log.Println("Erro ao criar o setor:", err)
		return "", fmt.Errorf("erro ao criar o setor: %v", err)
	}

	// Retorno de sucesso
	return "Setor criado com sucesso", nil
}

// funcao cadastrar usuario
func CadastrarUsuario(user, login, email, senha string, idsetor, permisao int) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 dbname=portal_bi host=192.168.1.65 sslmode=disable")
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return "", err
	}
	defer conn.Close(ctx)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Erro ao gerar hash da senha:", err)
		return "", fmt.Errorf("erro ao gerar hash da senha: %v", err)
	}

	queries := bd_portal_bi.New(conn)
	params := bd_portal_bi.CreateUsuariosParams{
		Nome:       pgtype.Text{String: user, Valid: true},
		IDSetor:    pgtype.Int4{Int32: int32(idsetor), Valid: true},
		IDPermisao: pgtype.Int4{Int32: int32(permisao), Valid: true},
		Email:      pgtype.Text{String: email, Valid: true},
		Login:      pgtype.Text{String: login, Valid: true},
		Senha:      pgtype.Text{String: string(hashedPassword), Valid: true},
	}
	err = queries.CreateUsuarios(ctx, params)
	if err != nil {
		fmt.Println("Erro ao criar o usuario")
		return "", fmt.Errorf("erro ao criar usuario: %v", err)
	}
	return "Usuario criado com sucesso", nil
}

// funcao cadastrar link do power bi
func SalvarBi(link string, id_setor int) (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=sa password=H0rt@1!ca5 dbname=portal_bi host=192.168.1.65 sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return "Erro ao conectar ao banco de dados", err
	}
	defer conn.Close(ctx)
	queries := bd_portal_bi.New(conn)
	params := bd_portal_bi.CreateLinkParams{
		Link:    pgtype.Text{String: link, Valid: true},
		IDSetor: pgtype.Int4{Int32: int32(id_setor), Valid: true},
	}
	err = queries.CreateLink(ctx, params)
	if err != nil {
		fmt.Println("Erro ao salvar o link")
		return "", fmt.Errorf("erro ao salvar link: %v", err)

	}

	return "link do Bi cadastrado com sucesso", nil

}
