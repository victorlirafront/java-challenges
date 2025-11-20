package com.javanauta.cadastro_usuario.business;

import com.javanauta.cadastro_usuario.infrastructure.entitys.Usuario;
import com.javanauta.cadastro_usuario.infrastructure.repository.UsuarioRepository;
import org.springframework.stereotype.Service;

@Service // Indica que essa classe é um componente de serviço do Spring (camada de negócio)
public class UsuarioService {

    // Repositório responsável por interagir com o banco de dados
    private final UsuarioRepository repository;

    // Injeção de dependência do repositório via construtor
    public UsuarioService(UsuarioRepository repository) {
        this.repository = repository;
    }

    // Método para salvar um novo usuário no banco
    public void salvarUsuario(Usuario usuario){
        repository.saveAndFlush(usuario); // salva e força a escrita imediata no banco
    }

    // Busca um usuário pelo e-mail
    public Usuario buscarUsuarioPorEmail(String email){
        // Caso não encontre o e-mail, lança uma exceção
        return repository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("Email não encontrado"));
    }

    // Deleta um usuário com base no e-mail
    public void deletarUsuarioPorEmail(String email){
        repository.deleteByEmail(email);
    }

    // Atualiza um usuário com base no ID informado
    public void atualizarUsuarioPorId(Integer id, Usuario usuario){
        // Busca o usuário existente no banco
        Usuario usuarioEntity = repository.findById(id)
                .orElseThrow(() -> new RuntimeException("Usuario não encontrado"));

        // Cria uma nova instância do usuário com os dados atualizados
        // Se algum campo vier nulo, mantém o valor anterior
        Usuario usuarioAtualizado = Usuario.builder()
                .email(usuario.getEmail() != null ? usuario.getEmail() : usuarioEntity.getEmail())
                .nome(usuario.getNome() != null ? usuario.getNome() : usuarioEntity.getNome())
                .id(usuarioEntity.getId()) // mantém o mesmo ID
                .build();

        // Salva as alterações no banco
        repository.saveAndFlush(usuarioAtualizado);
    }
}
