package com.javanauta.cadastro_usuario.controller;

import com.javanauta.cadastro_usuario.business.UsuarioService;
import com.javanauta.cadastro_usuario.dto.request.UsuarioRequestDTO;
import com.javanauta.cadastro_usuario.dto.request.UsuarioUpdateRequestDTO;
import com.javanauta.cadastro_usuario.dto.response.UsuarioResponseDTO;
import jakarta.validation.Valid;
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
    public ResponseEntity<UsuarioResponseDTO> salvarUsuario(@Valid @RequestBody UsuarioRequestDTO usuarioRequest) {
        // Chama o serviço para salvar o usuário recebido no corpo da requisição
        // O serviço retorna o DTO de resposta com os dados do usuário criado
        UsuarioResponseDTO usuarioResponse = usuarioService.salvarUsuario(usuarioRequest);
        // Retorna uma resposta 200 OK com os dados do usuário criado
        return ResponseEntity.ok(usuarioResponse);
    }

    // Endpoint para buscar um usuário pelo email (HTTP GET)
    @GetMapping
    public ResponseEntity<UsuarioResponseDTO> buscarUsuarioPorEmail(@RequestParam String email) {
        // Chama o serviço para buscar o usuário pelo email passado como parâmetro na URL
        // O serviço retorna o DTO de resposta
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
    public ResponseEntity<UsuarioResponseDTO> atualizarUsuarioPorId(@RequestParam Integer id,
                                                                     @Valid @RequestBody UsuarioUpdateRequestDTO usuarioUpdateRequest) {
        // Chama o serviço para atualizar o usuário com o ID informado
        // O serviço retorna o DTO de resposta com os dados atualizados
        UsuarioResponseDTO usuarioResponse = usuarioService.atualizarUsuarioPorId(id, usuarioUpdateRequest);
        // Retorna uma resposta 200 OK com os dados do usuário atualizado
        return ResponseEntity.ok(usuarioResponse);
    }
}
