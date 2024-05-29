# Repositório de Estudo de Golang

Este repositório contém as aulas e exemplos práticos do curso de Golang, organizado por [Roberto Camargo](https://github.com/Lajuro). Cada aula aborda um tópico específico da linguagem Go, com exemplos de código e explicações detalhadas.

## Índice

- [Repositório de Estudo de Golang](#repositório-de-estudo-de-golang)
  - [Índice](#índice)
  - [Introdução](#introdução)
  - [Instalação](#instalação)
  - [Resumo das Aulas](#resumo-das-aulas)
    - [Aula 1: Introdução ao Curso de Golang](#aula-1-introdução-ao-curso-de-golang)
    - [Aula 2: Instalação do Golang e Configuração do VS Code](#aula-2-instalação-do-golang-e-configuração-do-vs-code)
    - [Aula 3: Criando um "Hello, World!" Personalizado](#aula-3-criando-um-hello-world-personalizado)
    - [Aula 4: Trabalhando com Pacotes em Go](#aula-4-trabalhando-com-pacotes-em-go)
    - [Aula 5: Adicionando Pacotes Externos](#aula-5-adicionando-pacotes-externos)
    - [Aula 6: Variáveis, Constantes e Tipos de Dados](#aula-6-variáveis-constantes-e-tipos-de-dados)
    - [Aula 8: Funções](#aula-8-funções)
    - [Aula 9: Operadores](#aula-9-operadores)
    - [Aula 10: Structs](#aula-10-structs)
    - [Aula 11: "Herança"](#aula-11-herança)
    - [Aula 12: Ponteiros](#aula-12-ponteiros)
    - [Aula 13: Arrays e Slices](#aula-13-arrays-e-slices)
    - [Aula 14: Arrays Internos](#aula-14-arrays-internos)
    - [Aula 15: Maps](#aula-15-maps)
  - [Como Contribuir](#como-contribuir)
  - [Licença](#licença)
  - [Contato](#contato)

## Introdução

Este repositório foi criado para armazenar e organizar o material de estudo do curso de Golang. A linguagem Go é conhecida por sua simplicidade e eficiência, sendo uma das melhores escolhas para desenvolvimento de software de alto desempenho.

## Instalação

Siga as instruções abaixo para configurar o ambiente de desenvolvimento para Go:

1. Acesse [golang.org](https://golang.org) e baixe a versão adequada para seu sistema operacional.
2. Siga as instruções de instalação fornecidas no site.
3. Verifique a instalação executando `go version` no terminal.

## Resumo das Aulas

### Aula 1: Introdução ao Curso de Golang

**Descrição:** Apresentação do curso de Golang, incluindo a estrutura do curso e os motivos para aprender a linguagem. Foi destacada a eficiência, simplicidade e suporte pelo Google.

### Aula 2: Instalação do Golang e Configuração do VS Code

**Descrição:** Instruções detalhadas sobre como instalar Golang e configurar o VS Code com as extensões necessárias para desenvolvimento em Go.

### Aula 3: Criando um "Hello, World!" Personalizado

**Descrição:** Implementação de um programa "Hello, World!" interativo que solicita o nome do usuário e exibe uma saudação personalizada.

### Aula 4: Trabalhando com Pacotes em Go

**Descrição:** Criação e utilização de pacotes em Go, incluindo a criação de pacotes auxiliares e a importação de pacotes externos.

### Aula 5: Adicionando Pacotes Externos

**Descrição:** Adição de pacotes externos ao projeto Go usando `go get` e gerenciamento de dependências com `go.mod`.

### Aula 6: Variáveis, Constantes e Tipos de Dados

**Descrição:** Explicação detalhada sobre variáveis, constantes e tipos de dados em Go, abordando inteiros, floats, strings, booleanos e o tipo `error`.

### Aula 8: Funções

**Descrição:** Introdução às funções em Go, incluindo definição, chamada, e uso de funções como variáveis. Também foi abordado como as funções podem retornar múltiplos valores e como ignorar valores retornados.

**Tópicos Cobertos:**

- Definindo funções em Go
- Parâmetros e retorno de funções
- Chamando funções
- Funções como variáveis
- Funções com múltiplos retornos
- Ignorando valores retornados

### Aula 9: Operadores

**Descrição:** Introdução aos operadores em Go, cobrindo operadores aritméticos, de atribuição, relacionais, lógicos e unitários. Também foi discutida a ausência do operador ternário e como substituí-lo com `if`.

**Tópicos Cobertos:**

- Operadores aritméticos: +, -, *, /, %
- Operadores de atribuição: =, :=, +=, -=, *=, /=
- Operadores relacionais: ==, !=, >, <, >=, <=
- Operadores lógicos: &&, ||, !
- Operadores unitários: ++, --
- Substituição do operador ternário com `if`

### Aula 10: Structs

**Descrição:** Introdução aos structs em Go, a estrutura de dados que permite agrupar múltiplos campos. Um struct é semelhante a uma classe em linguagens orientadas a objetos, mas sem suporte a métodos ou herança.

**Tópicos Cobertos:**

- Definindo structs
- Instanciando structs
- Valor zero de structs
- Structs aninhados

### Aula 11: "Herança"

**Descrição:** Introdução ao conceito de pseudo-herança em Go utilizando structs. Em Go, a herança tradicional de linguagens OO é substituída pela composição de structs.

**Tópicos Cobertos:**

- Definindo structs base e derivados
- Composição de structs
- Instanciando structs compostos
- Acessando campos de structs compostos
- Benefícios da composição

### Aula 12: Ponteiros

**Descrição:** Introdução aos ponteiros em Go, explicando o conceito de ponteiros como referências de memória e como utilizá-los para acessar e modificar dados diretamente.

**Tópicos Cobertos:**

- Declaração e inicialização de ponteiros
- Atribuição de endereços de memória a ponteiros
- Desreferenciação de ponteiros
- Exemplos práticos de uso de ponteiros

### Aula 13: Arrays e Slices

**Descrição:** Introdução a arrays e slices em Go, explicando a diferença entre essas duas estruturas de dados e como utilizá-las.

**Tópicos Cobertos:**

- Declaração e inicialização de arrays
- Acesso e modificação de elementos em arrays
- Declaração e inicialização de slices
- Adição de elementos a slices usando a função `append`
- Criação de slices a partir de arrays
- Uso da função `reflect.TypeOf` para verificar o tipo de uma variável
- Compreensão de slices como ponteiros para arrays subjacentes

### Aula 14: Arrays Internos

**Descrição:** Explicação sobre a função `make` e como arrays internos funcionam em Go, incluindo o relacionamento entre slices e arrays.

**Tópicos Cobertos:**

- Uso da função `make` para alocar memória
- Diferença entre tamanho e capacidade de slices
- Crescimento dinâmico de slices
- Entendimento de arrays internos e como slices referenciam arrays subjacentes
- Exemplos práticos e uso das funções `len` e `cap`

### Aula 15: Maps

**Descrição:** Introdução aos maps em Go, uma estrutura chave-valor rígida onde tanto as chaves quanto os valores devem ser do mesmo tipo. Cobertura dos conceitos de criação, manipulação e deleção de elementos em maps.

**Tópicos Cobertos:**

- Criação de maps usando a palavra-chave `make`
- Tipagem das chaves e valores em maps
- Diferenças entre maps e structs
- Acessando e modificando valores em maps
- Deletando chaves de maps
- Aninhamento de maps
- Exemplos práticos de manipulação de maps

## Como Contribuir

1. Faça um fork deste repositório.
2. Crie uma branch para sua feature (`git checkout -b feature/fooBar`).
3. Commit suas mudanças (`git commit -m 'feat: Adicionar fooBar'`).
4. Envie para o branch (`git push origin feature/fooBar`).
5. Crie um novo Pull Request.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Para mais informações, entre em contato:

- **Nome:** Roberto Camargo
- **LinkedIn:** [Roberto Camargo](https://www.linkedin.com/in/robertocamargo96/)
- **GitHub:** [Lajuro](https://github.com/Lajuro)
