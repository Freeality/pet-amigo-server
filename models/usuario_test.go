package models

import "testing"

func TestNewUsuarioValido(t *testing.T) {
	usuario, err := NewUsuario("usuario1", "email1@email.com", "1234")

	if err != nil {
		t.Errorf("Erro tentando criar usuário: %v", err)
		return
	}
	t.Logf("Tudo certo: %v", usuario)
}

func TestNewUsuarioNomeInvalido(t *testing.T) {
	usuario, err := NewUsuario("us", "email", "")

	if err != nil {
		t.Logf("Erro identificado com sucesso: %v", err)
		return
	}
	t.Errorf("Usuário inválido não poderia ser criado: %v", usuario)
}

func TestNewUsuarioEmailInvalido(t *testing.T) {
	usuario, err := NewUsuario("Nome_Sobrenome", "email", "")

	if err != nil {
		t.Logf("Erro identificado com sucesso: %v", err)
		return
	}
	t.Errorf("Usuário inválido não poderia ser criado: %v", usuario)
}