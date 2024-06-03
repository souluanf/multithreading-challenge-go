# Desafio: Busca de CEP com Multithreading e APIs

## Descrição

Neste desafio, você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas. As duas requisições serão feitas simultaneamente para as seguintes APIs:
- https://brasilapi.com.br/api/cep/v1/{cep}
- http://viacep.com.br/ws/{cep}/json/

Os requisitos para este desafio são:
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Como Executar

### Pré-requisitos

- Go 1.18+ instalado

### Clonar o Repositório

```sh
git clone https://github.com/souluanf/multithreading-challenge-go.git
cd multithreading-challenge-go
```

### Executar a Aplicação

Para executar a aplicação, forneça um CEP como argumento:

```sh
go run main.go <CEP>
```

Exemplo:

```sh
go run main.go 08210010
```

## Exemplo de Saída

```
CEP: 08210010
Buscando em brasilapi...
Buscando em viacep...
Resultado mais rápido: brasilapi
{
        "cep": "08210010",
        "city": "São Paulo",
        "neighborhood": "Itaquera",
        "service": "open-cep",
        "state": "SP",
        "street": "Rua Barra de Guabiraba"
}
```
