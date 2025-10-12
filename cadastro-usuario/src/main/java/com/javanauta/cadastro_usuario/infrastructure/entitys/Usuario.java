package com.javanauta.cadastro_usuario.infrastructure.entitys;

import jakarta.persistence.*;
import lombok.*;

// Gera automaticamente os métodos getters para todos os campos da classe
@Getter

// Gera automaticamente os métodos setters para todos os campos da classe
@Setter

// Cria um construtor com todos os atributos como parâmetros
@AllArgsConstructor

// Cria um construtor vazio (sem parâmetros)
@NoArgsConstructor

// Permite construir objetos usando o padrão Builder (ex: Usuario.builder().email("...").nome("...").build())
@Builder

// Define o nome da tabela no banco de dados como "usuario"
@Table(name = "usuario")

// Indica que esta classe é uma entidade JPA, ou seja, será mapeada para uma tabela no banco
@Entity
public class Usuario {

    // Define o campo "id" como chave primária da tabela
    @Id
    // Especifica que o valor do ID será gerado automaticamente (auto incremento)
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Integer id;

    // Mapeia o campo "email" para a coluna "email" no banco e garante que ele seja único
    @Column(name = "email", unique = true)
    private String email;

    // Mapeia o campo "nome" para a coluna "nome" no banco
    @Column(name = "nome")
    private String nome;
}