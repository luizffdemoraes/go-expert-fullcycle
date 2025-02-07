**Ponteiros**

1️⃣ **O operador & (E comercial - "endereço de")**
O operador & é usado para obter o endereço de memória de uma variável.
Ele retorna o endereço da variável na memória, permitindo que um ponteiro armazene esse endereço.

✅ O &x retorna o endereço de memória onde x está armazenado.

2️⃣ **O operador * (Asterisco - "desreferência")**
Quando usado em uma declaração (int *ptr), indica que a variável é um ponteiro.
Quando usado antes de um ponteiro (*ptr), ele desreferencia o ponteiro, acessando o valor armazenado no endereço que ele contém.

✅ O *ptr acessa (ou modifica) o valor armazenado no endereço que ptr contém.


🔎 Resumindo:

| Expressão  | Significado |
|------------|------------|
| `&variavel` | Retorna o endereço de memória da variável. |
| `*ptr` | Desreferencia o ponteiro, acessando/modificando o valor armazenado. |
| `int *ptr;` | Declara um ponteiro para um inteiro. |

[⬅ Voltar para o README principal](/README.MD)