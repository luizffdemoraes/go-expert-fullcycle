**Dados Compostos.:**

Em Go, slices e arrays são tipos de dados compostos usados para armazenar sequências de elementos do mesmo tipo.

Eles pertencem à categoria de tipos de dados compostos (composite types), pois são estruturas que agrupam múltiplos valores.

🚀 Por que um slice dobra de tamanho ao adicionar novos elementos?
Isso acontece por causa do gerenciamento interno de memória em Go. O append() duplica a capacidade do slice quando ele atinge o limite, para otimizar o desempenho e evitar realocações frequentes.

🧐 O que acontece por baixo dos panos?
Criamos um slice com tamanho 2 e capacidade 2.
Ao adicionar o terceiro elemento (30), a capacidade dobra para 4.
Ao adicionar mais dois (40 e 50), a capacidade dobra novamente para 8.
⚡ Como o crescimento funciona?
O Go gerencia a alocação da seguinte forma:

Se a capacidade inicial for pequena, ela dobra a cada realocação.
Se a capacidade for grande (>1024 elementos), o crescimento será gradual (~25% por vez).
Isso minimiza chamadas ao garbage collector e melhora a eficiência.

📌 Como evitar crescimento desnecessário?
Se já souber quantos elementos serão adicionados, use make() para definir a capacidade:

```go
slc := make([]int, 2, 10) // Define capacidade maior desde o início
```

## 📌 Nome Técnico dos Tipos

| Tipo  | Nome Técnico  |
|-------|--------------|
| **Array** | *Array Type* |
| **Slice** | *Slice Type* |


## 📍 Diferença entre Array e Slice

| Característica                  | Array (`[N]T`)          | Slice (`[]T`)                   |
|---------------------------------|-------------------------|---------------------------------|
| **Tamanho fixo?**               | ✅ Sim                  | ❌ Não (dinâmico)               |
| **Armazena os dados diretamente?** | ✅ Sim                  | ❌ Não (aponta para um array subjacente) |
| **Pode crescer?**                | ❌ Não                  | ✅ Sim (usando `append()`)      |
| **Tipo diferente de slice?**     | ✅ Sim (`[5]int` ≠ `[10]int`) | ❌ Não (qualquer `[]int` é compatível) |

## 🔚 Resumo

| Ação                      | `len(slc)` | `cap(slc)` | Alocação de Memória          |
|---------------------------|-----------|-----------|------------------------------|
| Criando `make([]int, 2, 2)` | 2         | 2         | Capacidade inicial          |
| `append(slc, 30)`        | 3         | 4         | Capacidade dobra            |
| `append(slc, 40, 50)`    | 5         | 8         | Capacidade dobra            |

[⬅ Voltar para o README principal](/README.MD)
