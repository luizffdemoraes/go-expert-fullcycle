# Multithreading

Em Go, o termo "processamento" geralmente se refere à forma como a linguagem trata a execução de tarefas, tanto de maneira concorrente quanto paralela. Go foi projetado para facilitar a criação de aplicações escaláveis e de alto desempenho, graças à sua abordagem nativa para concorrência, baseada em **goroutines** e **channels**.

---

## 1. Goroutines

### O que são:
- **Goroutines** são funções que podem ser executadas concorrentemente.  
- São muito leves em comparação com threads tradicionais de sistemas operacionais e são gerenciadas pela runtime do Go.

---

## 2. Channels

### O que são:
- **Channels** são mecanismos de comunicação entre goroutines.  
- Eles permitem que as goroutines se comuniquem e sincronizem o acesso a dados de forma segura, evitando condições de corrida.

---

## 3. Concorrência vs. Paralelismo

### Concorrência:
- Trata da capacidade de gerenciar várias tarefas de forma intercalada (simultaneamente), mesmo em um único núcleo. Em Go, isso é feito com goroutines e channels.

- As goroutines são concorrentes, permitindo escrever código que pareça "executar ao mesmo tempo", embora, na prática, elas sejam escalonadas pela runtime do Go.

### Paralelismo:
- Ocorre quando várias tarefas realmente são executadas ao mesmo tempo, em múltiplos núcleos de processamento.
- O Go permite paralelismo configurando a quantidade de núcleos que serão utilizados com a função `runtime.GOMAXPROCS`.

- Refere-se à execução simultânea de tarefas em múltiplos núcleos, o que pode aumentar a performance, mas exige cuidado na sincronização do acesso à memória.

- Trata da capacidade de gerenciar diversas tarefas ao mesmo tempo, permitindo que elas progridam intercaladamente, mesmo que em um único núcleo. Em Go, isso é feito com goroutines e channels.

### Observação:
- Embora o paralelismo permita que tarefas sejam processadas ao mesmo tempo, isso pode levar a conflitos de memória se múltiplas goroutines acessarem ou modificarem os mesmos dados simultaneamente. Esse tipo de condição pode resultar em inconsistências e erros inesperados.

### Solução:
- Para evitar esses problemas, é recomendável utilizar mecanismos de sincronização, como mutexes. Um mutex (mutual exclusion lock) garante que apenas uma goroutine acesse uma seção crítica do código por vez. Assim, mesmo quando tarefas são processadas paralelamente, o acesso aos dados compartilhados é controlado, permitindo que o valor seja atualizado de forma segura e liberado somente após a finalização do processo em execução.

---

## 4. Processamento em Go: O Que Isso Significa?

### Eficiência:
- As goroutines são extremamente leves, permitindo a criação de milhares ou até milhões delas sem sobrecarregar o sistema.

### Facilidade de Uso:
- A linguagem Go simplifica a criação e o gerenciamento de tarefas concorrentes com a palavra-chave `go` e canais para comunicação.

### Escalabilidade:
- Graças à concorrência e ao paralelismo, aplicações em Go podem escalar eficientemente em sistemas com múltiplos núcleos, melhorando o desempenho sem complicações significativas.

---

## 5. Conclusão

Processamento em Go envolve o uso de goroutines e channels para criar aplicações concorrentes e paralelas. Essa abordagem torna o desenvolvimento mais simples e escalável, permitindo que tarefas sejam executadas de forma eficiente sem a complexidade das threads tradicionais.

---

## 6. Condições de Corrida e Mutex

Quando goroutines compartilham dados, é essencial garantir que elas não modifiquem os dados simultaneamente sem controle, evitando resultados inesperados. Para isso, podemos usar um **mutex**.

### O que é um Mutex?
- **Mutex** (mutual exclusion lock) é uma primitiva que permite bloquear uma seção crítica do código, garantindo que apenas uma goroutine possa executá-la de cada vez.
- No Go, o pacote `sync` fornece o tipo `Mutex`.

---

[⬅ Voltar para o README principal](/README.MD)