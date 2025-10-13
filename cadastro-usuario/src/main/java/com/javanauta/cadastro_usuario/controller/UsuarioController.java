package com.javanauta.cadastro_usuario.controller;

import com.javanauta.cadastro_usuario.business.UsuarioService;
import com.javanauta.cadastro_usuario.infrastructure.entitys.Usuario;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

// Indica que esta classe é um controlador REST, responsável por receber requisições HTTP
@RestController
// Define o endpoint base: todas as rotas começarão com "/usuario"
@RequestMapping("/usuario")
// Cria automaticamente um construtor com os campos 'final' (injeção de dependência via construtor)
@RequiredArgsConstructor
public class UsuarioController {

    // Injeção do serviço que contém as regras de negócio de usuário
    private final UsuarioService usuarioService;

    // Endpoint para salvar um novo usuário (HTTP POST)
    @PostMapping
    public ResponseEntity<Void> salvarUsuario(@RequestBody Usuario usuario) {
        // Chama o serviço para salvar o usuário recebido no corpo da requisição
        usuarioService.salvarUsuario(usuario);
        // Retorna uma resposta 200 OK sem corpo
        return ResponseEntity.ok().build();
    }

    // Endpoint para buscar um usuário pelo email (HTTP GET)
    @GetMapping
    public ResponseEntity<Usuario> buscarUsuarioPorEmail(@RequestParam String email) {
        // Chama o serviço para buscar o usuário pelo email passado como parâmetro na URL
        return ResponseEntity.ok(usuarioService.buscarUsuarioPorEmail(email));
    }

    // Endpoint para deletar um usuário pelo email (HTTP DELETE)
    @DeleteMapping
    public ResponseEntity<Void> deletarUsuarioPorEmail(@RequestParam String email) {
        // Chama o serviço para deletar o usuário pelo email informado
        usuarioService.deletarUsuarioPorEmail(email);
        // Retorna uma resposta 200 OK sem corpo
        return ResponseEntity.ok().build();
    }

    // Endpoint para atualizar os dados de um usuário existente (HTTP PUT)
    @PutMapping
    public ResponseEntity<Void> atualizarUsuarioPorId(@RequestParam Integer id,
                                                      @RequestBody Usuario usuario) {
        // Chama o serviço para atualizar o usuário com o ID informado
        usuarioService.atualizarUsuarioPorId(id, usuario);
        // Retorna uma resposta 200 OK sem corpo
        return ResponseEntity.ok().build();
    }
}
