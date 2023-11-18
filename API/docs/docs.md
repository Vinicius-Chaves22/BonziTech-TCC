# BonziTech-TCC API

API hospedada em: http://154.56.41.79:3304/

## Sumário

* [Status Codes](#status-codes)
* [Responses](#responses)
* [Código de sessão](#código-de-sessão-codsessao) 
* [/api/auth](#apiauth)
* [/api/categorias](#apicategorias)
* [/api/clientes](#apiclientes)
  * [/api/clientes/telefones](#apiclientestelefones)
* [/api/componentes](#apicomponentes)
* [/api/entradas](#apientradas)
* [/api/estoque](#apiestoque)
* [/api/fabricantes](#apifabricantes)
* [/api/ordem-servico](#apiordem-servico)
* [/api/sessao](#apisessao)
* [/api/subcategorias](#apisubcategorias)
* [/api/usuarios](#apiusuarios)
* [/api/vendas](#apivendas)

## Status Codes

Os status de respostas possíveis para esta API são:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |

## Responses

Os endpoints terão sempre 2 possíveis responses, 
sendo elas:

```javascript
{
  ...
  "message": string
}
```

```javascript
{
  "error": string
}
```

O atributo `message` estará presente caso uma request seja concluída com sucesso, e ausente caso contrário.

O atributo `error` estará presente caso uma request não seja devidamente concluída, retornando o devido erro.

Além disso, os endpoints podem retornar outros atributos específicos daquele endpoint, mas sempre
serão retornados estes dois atributos.

## Código de sessão (codSessao):

Muitas funções da API necessitam da identificação do usuário para serem efetudas, por isso, sempre, ao
fazer uma request à API, deve ser passado, pelos headers, o código de sessão do usuário, pela chave (key)
`codSessao` ou `Codsessao`.

Para ações como login, ou ping, não é necessário o código de sessão do usuário.

## Endpoints:

#### Verificar conexão com API

```http
GET /api/ping
```


## /api/auth 

#### Validar login

```http
POST /api/auth/login
```

Request body:

```javascript
{
  "email": string
  "senha": string
}
```

Exemplo:

```javascript
{
  "email": "teste",
  "senha": "teste123"
}
```

Response:

```javascript
{
  "codSessao": Number
}
```


## /api/categorias

#### Retornar todas categorias

```http
GET /api/categorias
```

Response:
```javascript
{
  "categorias": []object || null
}
```

#### Retornar dados e componentes de uma categoria

```http
GET /api/categorias/${codCat}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codCat`      | `string` | **Required**. Código da categoria a ser mostrada |

Response:
```javascript
{
  "componentes": []object,
  "categoria": {
    "codCat": Number,
    "nomeCat": string,
    "unidMedida": string,
    "montagem": boolean,
    "apelido": string
  }
}
```

#### Criar categoria

```http
POST /api/categorias
```

Request body:
```javascript
{
  "nomeCat": string,
  "unidMedida": string,
  "montagem": boolean,
  "apelido": string
}
```

Exemplo:
```javascript
{
  "nomeCat": "Mangueiras mais brabas de 2012",
  "unidMedida": "cm",
  "montagem": false,
  "apelido": "MMB"
}
```

#### Editar categoria

```http
PUT /api/categorias
```

Request body:
```javascript
{
  "codCat": Number,
  "nomeCat": string,
  "unidMedida": string,
  "montagem": boolean,
  "apelido": string
}
```

#### Excluir categoria

```http
DELETE /api/categorias/${codCat}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codCat`      | `string` | **Required**. Código da categoria a ser mostrada |


## /api/clientes

#### Mostrar todos clientes

```http
GET /api/clientes
```

Response:
```javascript
{
  "clientes": []object || null
}
```

#### Mostrar dados de um cliente

```http
GET /api/clientes/${codCli}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codCli`      | `string` | **Required**. Código do cliente a ser mostrado |

Response:
```javascript
{
  "cliente": {
    "codCli": number,
    "nomeEmpresa": string,
    "nome": string,
    "tipo": string || null,
    "diaReg": string,
    "endereco": string || null,
    "bairro": string || null,
    "cidade": string,
    "estado": string,
    "cep": string || null,
    "email": string || null
  },
  "telefones": []object || null
}
```

#### Cadastrar novo cliente

```http
POST /api/clientes
```

Request body:
```javascript
{
  "nomeEmpresa": string,
  "nome": string,
  "tipo": string || null,
  "endereco": string || null,
  "bairro": string || null,
  "cidade": string,
  "estado": string,
  "cep": string || null,
  "email": string || null
}
```

Exemplo:
```javascript
{
  "nomeEmpresa": "Beer Lanches",
  "nome": "Alberto",
  "tipo": null,
  "endereco": null,
  "bairro": null,
  "cidade": "Santa Rosa",
  "estado": "RS",
  "cep": null,
  "email": "beer@lanches.rs"
}
```

#### Atualizar cliente

```http
PUT /api/clientes
```

Request body:
```javascript
{
  "codCli": Number,
  "nomeEmpresa": string,
  "nome": string,
  "tipo": string || null,
  "endereco": string || null,
  "bairro": string || null,
  "cidade": string,
  "estado": string,
  "cep": string || null,
  "email": string || null
}
```

Exemplo:
```javascript
{
  "codCli": 2,
  "nomeEmpresa": "Beer Lanches",
  "nome": "sr. Beer Lanches",
  "tipo": null,
  "endereco": null,
  "bairro": null,
  "cidade": "Santa Rosa",
  "estado": "RS",
  "cep": null,
  "email": "beer@lanches.rs"
}
```

#### Deletar cliente

```http
DELETE /api/clientes/${codCli}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codCli`      | `string` | **Required**. Código do cliente a ser removido do estoque |


## /api/clientes/telefones

#### Cadastrar números de telefone de cliente

```http
POST /api/clientes/telefones
```

Vale notar que este endpoint permite que você envie múltiplos objetos de
uma só vez por meio de uma array.

Request body:
```javascript
[
  {
    "codCli": number,
    "telefone": string,
    "nomeTel": string, 
    "tipoContato": string || null,
    "tipoCli": string || null
  },
  ...
]
```

Exemplo:
```javascript
[
  {
    "codCli": 1,
    "telefone": "(19) 99999-9999",
    "nomeTel": "Telefone pessoal", 
    "tipoContato": null,
    "tipoCli": null
  }
]
```

#### Atualizar cadastro de telefone

```http
PUT /api/clientes/telefones
```

Request body:
```javascript
{
  "codTel": number,
  "telefone": string,
  "nomeTel": string, 
  "tipoContato": string || null,
  "tipoCli": string || null
}
```

Exemplo:
```javascript
{
  "codTel": 1,
  "telefone": "(19) 99999-9999",
  "nomeTel": "Telefone pessoal", 
  "tipoContato": null,
  "tipoCli": null
}
```

#### Excluir telefone

```http
DELETE /api/clientes/telefones/${codTel}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codTel`      | `string` | **Required**. Código do telefone a ser excluído |


## /api/componentes

#### Mostrar todos componentes

```http
GET /api/componentes
```

Response:
```javascript
{
  "componentes": []object || null
}
```

#### Mostrar componente

```http
GET /api/componentes/${codComp}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codComp`      | `string` | **Required**. Código de componente a ser mostrado |

Response:
```javascript
{
  "componente": {
    "codComp": Number,
    "codPeca": Number,
    "especificacao": string,
    "codCat": Number,
    "codSubcat": Number || null,
    "diamInterno": string || null,
    "diamExterno": Number || null,
    "diamNominal": string || null,
    "medidaD": Number || null,
    "costura": boolean || null,
    "prensadoReusavel": string || null,
    "mangueira": string || null,
    "material": string || null,
    "norma": string || null,
    "bitola": Number || null,
    "valorEntrada": Number,
    "valorVenda": Number
  }
}
```

#### Criar novo componente

```http
POST /api/componentes
```

Request body:
```javascript
{
  "codPeca": Number,
  "especificacao": string,
  "codCat": Number,
  "codSubcat": Number || null,
  "diamInterno": string || null,
  "diamExterno": Number || null,
  "diamNominal": string || null,
  "medidaD": Number || null,
  "costura": boolean || null,
  "prensadoReusavel": string || null,
  "mangueira": string || null,
  "material": string || null,
  "norma": string || null,
  "bitola": Number || null,
  "valorEntrada": Number,
  "valorVenda": Number
}
```

Exemplo:
```javascript
{
  "codPeca": "ABCDE",
  "especificacao": "sla man",
  "codCat": 2,
  "codSubcat": null,
  "diamInterno": null,
  "diamExterno": null,
  "diamNominal": null,
  "medidaD": null,
  "costura": null,
  "prensadoReusavel": null,
  "mangueira": null,
  "material": null,
  "norma": null,
  "bitola": null,
  "valorEntrada": 10,
  "valorVenda": 85
}
```

#### Atualizar componente

```http
PUT /api/componentes
```

Request body:
```javascript
{
  "codComp": Number,
  "codPeca": Number,
  "especificacao": string,
  "codCat": Number,
  "codSubcat": Number || null,
  "diamInterno": string || null,
  "diamExterno": Number || null,
  "diamNominal": string || null,
  "medidaD": Number || null,
  "costura": boolean || null,
  "prensadoReusavel": string || null,
  "mangueira": string || null,
  "material": string || null,
  "norma": string || null,
  "bitola": Number || null,
  "valorEntrada": Number,
  "valorVenda": Number
}
```

Exemplo:
```javascript
{
  "codComp": 1,
  "codPeca": "ABCDE",
  "especificacao": "sla man",
  "codCat": 2,
  "codSubcat": null,
  "diamInterno": null,
  "diamExterno": null,
  "diamNominal": null,
  "medidaD": null,
  "costura": null,
  "prensadoReusavel": null,
  "mangueira": null,
  "material": null,
  "norma": null,
  "bitola": null,
  "valorEntrada": 10,
  "valorVenda": 85
}
```

#### Deletar componente

```http
DELETE /api/componentes/${codComp}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codComp`      | `string` | **Required**. Código do componente a ser removido do estoque |

## /api/entradas

#### Mostrar todas entradas

```http
GET /api/entradas
```

Response:
```javascript
{
  "entradas": []object || null
}
```

#### Mostrar entrada

```http
GET /api/entradas/${codEntd}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codEntd`      | `string` | **Required**. Código da entrada a ser retornada |

Response:
```javascript
{
  "entrada": {
    "codEntd": Number,
    "codFab": Number || null,
    "nomeFab": string || null,
    "dataVenda": string,
    "notaFiscal": string || null,
    "valorTotal": Number
  },
  "componentes": []object || null
}
```

#### Criar entrada

```http
POST /api/entradas
```

Request body:
```javascript
{
  "codFab": Number || null,
  "dataVenda": string,
  "notaFiscal": string || null
}
```

Exemplo:
```javascript
{
  "codFab": 3,
  "dataVenda": "2022-01-01",
  "notaFiscal": null
}
```

#### Registrar componentes de uma entrada

```http
POST /api/entradas/componentes
```

Vale notar que, este endpoint permite que você envie múltiplos objetos de
uma só vez por meio de uma array.

Request body:
```javascript
[
  {
    "codEntd": Number,
    "codComp": Number,
    "quantidade": Number,
    "valorUnit": Number
  },
  ...
]
```

#### Atualizar entrada

```http
PUT /api/entradas
```

Request body:
```javascript
{
  "codCompEntd": Number,
  "codFab": Number || null,
  "dataVenda": string,
  "notaFiscal": string || null
}
```

Exemplo:
```javascript
{
  "codCompEntd": 1,
  "codFab": null,
  "dataVenda": "2022-01-01",
  "notaFiscal": null
}
```

#### Atualizar componente de uma entrada

```http
PUT /api/entradas/componentes
```

Request body:
```javascript
{
  "codCompEntd": Number,
  "codEntd": Number,
  "codComp": Number,
  "quantidade": Number,
  "valorUnit": Number
}
```

#### Deletar entrada

```http
DELETE /api/entradas/${codEntd}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codEntd`      | `string` | **Required**. Código da entrada a ser excluída |


## /api/estoque

#### Mostrar estoque

```http
GET /api/estoque
```

Response:
```javascript
{
  "estoque": []object || null
}
```

#### Adicionar componente ao estoque

```http
POST /api/estoque
```

Request body:
```javascript
{
  "codComp": Number,
  "min": Number,
  "max": Number,
  "quantidade": Number 
}
```

Exemplo:
```javascript
{
  "codComp": 12,
  "min": 1,        // caso min seja nulo, ele será 0
  "max": 100,      // caso max seja nulo, ele será 10000000
  "quantidade": 20
}
```

#### Editar dados de componente em estoque

```http
PUT /api/estoque
```

Request body:
```javascript
{
  "codComp": Number,
  "min": Number,
  "max": Number,
  "quantidade": Number 
}
```

#### Remover componente de estoque

```http
DELETE /api/estoque/${codComp}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codComp`      | `string` | **Required**. Código do componente a ser removido do estoque |

## /api/fabricantes

#### Retornar todos fabricantes

```http
GET /api/fabricantes
```

Response:
```javascript
{
  "fabricantes": []object || null
}
```

#### Retornar dados de um fabricante

```http
GET /api/fabricantes/${codFab}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codFab`      | `string` | **Required**. Código do fabricante a ser mostrado |

Response:
```javascript
{
  "fabricante": {
    "nome": string,
    "nomeContato": string || null,
    "razaoSocial": string || null,
    "telefone": string || null,
    "celular": string || null,
    "fax": string || null,
    "endereco": string || null,
    "cidade": string || null,
    "estado": string || null,
    "cep": string || null
  }
}
```

#### Cadastrar fabricante

```http
POST /api/fabricantes
```

Request body:
```javascript
{
  "nome": string,
  "nomeContato": string, // opcional
  "razaoSocial": string, // opcional
  "telefone": string,    // opcional
  "celular": string,     // opcional
  "fax": string,         // opcional
  "endereco": string,    // opcional
  "cidade": string,      // opcional
  "estado": string,      // opcional
  "cep": string          // opcional
}
```

Exemplo:
```javascript
{
  "nome": "Fabrício Caminhões & Motocas",
  "nomeContato": "Fabrício", 
  "telefone": "(19) 99999-9999",    
  "cidade": "Campinas"
}
```

#### Atualizar dados de fabricante

```http
PUT /api/fabricantes
```

Request body:
```javascript
{
  "codFab": number,
  "nome": string,
  "nomeContato": string || null, 
  "razaoSocial": string || null, 
  "telefone": string || null,    
  "celular": string || null,     
  "fax": string || null,         
  "endereco": string || null,
  "cidade": string || null,      
  "estado": string || null,      
  "cep": string || null          
}
```

Exemplo:
```javascript
{
  "cidFab": 1,
  "nome": "Fabrício Caminhões & Motocas",
  "nomeContato": "Fabrício", 
  "razaoSocial": null,
  "telefone": "(19) 99999-9999",   
  "celular": null,
  "fax": null,
  "endereco": null,
  "cidade": "Campinas",
  "estado": "SP",
  "cep": null
}
```

#### Excluir fabricante

```http
DELETE /api/fabricantes/${codFab}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codFab`      | `string` | **Required**. Código do fabricante a ser excluído |


## /api/ordem-servico

#### Mostrar todas ordens de serviço

```http
GET /api/ordem-servico
```

Response:
```javascript
{
  "ordensServico": []object || null
}
```

#### Mostrar dados de oderm de serviço

```http
GET /api/ordem-servico/${codOS}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codOS`      | `string` | **Required**. Código da ordem de serviço a ser retornada |

Response:
```javascript
{
  "ordemServico": {
    "codOs": Number,
    "dataEmissao": string,
    "codCli": Number,
    "pedido": string,
    "concluida": boolean
  }
}
```

#### Criar nova ordem de serviço

```http
POST /api/ordem-servico
```

Request body:
```javascript
{
  "dataEmissao": string,
  "codCli": Number,
  "pedido": string,
  "concluida": boolean
}
```

Exemplo:
```javascript
{
  "dataEmissao": "1970-01-01",
  "codCli": 2,
  "pedido": "1 km de mangueira",
  "concluida": false
}
```

#### Atualizar ordem de serviço

```http
PUT /api/ordem-servico
```

Request body:
```javascript
{
  "codOS": Number,
  "dataEmissao": string,
  "codCli": Number,
  "pedido": string,
  "concluida": boolean
}
```

Exemplo:
```javascript
{
  "codOS": 1,
  "dataEmissao": "1970-01-01",
  "codCli": 2,
  "pedido": "1 km de mangueira",
  "concluida": false
}
```

#### Deletar ordem de serviço

```http
DELETE /api/ordem-servico/${codOS}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codOS`      | `string` | **Required**. Código da ordem de serviço a ser excluída |


## /api/sessao

#### Retornar informações de sessão de usuário

```http
GET /api/sessao/${codSessao}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codSessao`      | `string` | **Required**. Código da sessão a ser procurada |

Response:

```javascript
{
  "codSessao": Number,
  "codUsuario": Number,
  "entrada": string,
  "saida": object
}
```

#### Criar sessão

```http
POST /api/sessao
```

Request body:

```javascript
{
  "codUsuario": Number
}
```

Exemplo:
```javascript
{
  "codUsuario": 10000000
}
```

Response:
```javascript
{
  "codSessao": Number
}
```

#### Encerrar sessão

```http
DELETE /api/sessao/${codSessao}
```


## /api/subcategorias

#### Retornar subcategorias de categoria

```http
GET /api/subcategorias/categoria/${codCat}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codCat`      | `string` | **Required**. Código da categoria principal |

Response:
```javascript
{
  "subcategorias": []object || null
}
```

#### Mostrar componentes da subcategoria e dados subcategoria

```http
GET /api/subcategorias/subcategoria/${codSubcat}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codSubcat`      | `string` | **Required**. Código da subcategoria a ser procurada |

Response:
```javascript
{
  "componentes": []object || null,
  "subcategoria": {
    "codSubcat": Number,
    "codCat": Number,
    "nome": string
  }
}
```

#### Criar subcategoria

```http
POST /api/subcategorias
```

Request body:
```javascript
{
  "codCat": Number, // Código da categoria principal da subcategoria
  "nome": string
}
```

Exemplo:
```javascript
{
  "codCat": 12, 
  "nome": "Mangueiras teste" 
}
```

#### Atualizar subcategoria

```http
PUT /api/subcategorias
```

Request body:
```javascript
{
  "codSubcat": Number,
  "codCat": Number, // Código da categoria principal da subcategoria
  "nome": string
}
```

Exemplo:
```javascript
{
  "codSubcat": 10,
  "codCat": 12, 
  "nome": "Mangueiras teste" 
}
```

#### Deletar subcategoria

```http
DELETE /api/subcategorias/${codSubcat}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codSubcat`      | `string` | **Required**. Código da subcategoria a ser excluída |

## /api/usuarios

#### Retornar todos os usuários do sistema

```http
GET /api/usuarios
```

Response:

```javascript
{
  "usuarios": []object || null
}
```

#### Retornar um usuário do sistema

```http
GET /api/usuarios/${codUsu}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codUsu`      | `string` | **Required**. Código do usuário a ser  procurado |

Response:

```javascript
{
  "usuario": object
}
```

#### Criar novo usuário

```http
POST /api/usuarios
```

Request body:

```javascript
{
  "permissoes": string,
  "nome": string,
  "email": string,
  "senha": string
}
```

Exemplo:
```javascript
{
  "permissoes": "Leitura",
  "nome": "Adalberto R.",
  "email": "adalbertorocha@gmail.com",
  "senha": "teste123"
}
```

#### Atualizar um usuário

```http
PUT /api/usuarios
```

Request body:

```javascript
{
  "codUsuario": Number,
  "permissoes": string,
  "nome": string,
  "email": string,
  "senha": string
}
```

Exemplo:
```javascript
{
  "permissoes": "Leitura",
  "nome": "Adalberto R.",
  "email": "adalbertorocha@gmail.com",
  "senha": "teste123"
}
```

#### Desativar um usuário

```http
DELETE /api/usuarios/${codUsu}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codUsu`      | `string` | **Required**. Código do usuário a ser inativado |


## /api/vendas

#### Mostrar todas vendas

```http
GET /api/vendas
```

Response:
```javascript
{
  "vendas": []object || null
}
```

#### Mostrar uma venda

```http
GET /api/vendas/${codVenda}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codVenda`      | `string` | **Required**. Código de venda a ser encontrada |


Response:
```javascript
{
  "venda": {
    "codVenda": Number,
    "dataVenda": string,
    "codCli": Number || null,
    "nomeCli": string || null,
    "codOS": Number,
    "valorTotal": Number,
    "descricao": string || null
  },
  "componentes": []object || null
}
```

#### Registrar nova venda

```http
POST /api/vendas
```

Request body:
```javascript
{
  "dataVenda": string,
  "codCli": Number || null,
  "codOS": Number,
  "descricao": string || null
}
```

Exemplo:
```javascript
{
  "dataVenda": "2020-01-15",
  "codCli": 1,
  "codOS": 2,
  "descricao": "Venda de 1 km de magueira ao Gustavo web"
}
```

#### Atualizar venda

```http
PUT /api/vendas
```

Request body:
```javascript
{
  "codVenda": Number,
  "dataVenda": string,
  "codCli": Number || null,
  "codOS": Number,
  "descricao": string || null
}
```

Exemplo:
```javascript
{
  "codVenda": 1
  "dataVenda": "2020-01-01",
  "codCli": 1,
  "codOS": 3,
  "descricao": "Venda de 2 km de mangueira ao Gustavo web"
}
```

#### Excluir venda

```http
DELETE /api/vendas/${codVendas}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codVendas`      | `string` | **Required**. Código da venda a ser excluída |