**scheduler.:**

## 1. scheduler
Um scheduler (agendador) é um componente do sistema (ou da runtime, no caso do Go) responsável por gerenciar a execução das tarefas – seja elas threads, goroutines ou processos. Ele decide qual tarefa será executada, quando e em qual núcleo de processamento, distribuindo os recursos disponíveis de forma eficiente.

## Preemptivo:
O sistema operacional interrompe automaticamente uma tarefa após um certo tempo de execução, independentemente de a tarefa solicitar ou não a pausa.
Exemplo: A maioria dos sistemas operacionais modernos (Windows, Linux, macOS) usam escalonamento preemptivo.

**Interrupção forçada:** O sistema pode interromper uma tarefa sem que ela precise cooperar. 

**Equidade:** Garante que todas as tarefas recebam tempo de CPU.
Complexidade: Requer mecanismos de interrupção e gerenciamento de contexto.


## Cooperativo:
Cada tarefa precisa explicitamente ceder o controle, por exemplo, chamando uma função de yield ou finalizando sua execução.
Exemplo: Algumas implementações antigas de sistemas operacionais móveis ou alguns ambientes específicos de programação (por exemplo, certos frameworks de corrotinas) utilizam esse modelo.

**Cessão voluntária:** Cada tarefa precisa ceder o controle para permitir que outras sejam executadas.

**Simplicidade:** Menor complexidade no gerenciamento de contexto.

**Risco:** Se uma tarefa não ceder o controle, pode causar problemas de responsividade.

## 2. Go scheduler
O scheduler do Go é uma parte fundamental da runtime que gerencia a execução das goroutines (as “green threads” do Go) e as distribui entre as threads do sistema operacional. Em vez de criar uma thread do sistema para cada tarefa (o que seria pesado e custoso em termos de memória e desempenho), o Go cria goroutines que são muito mais leves e são multiplexadas em um número menor de threads reais.

## Vantagens das Green Threads (Goroutines) em Go
**Leveza:**
- As goroutines têm overhead muito baixo, permitindo que você crie milhares ou milhões delas sem comprometer significativamente os recursos do sistema.

**Simplicidade:**
- A utilização da palavra-chave go torna muito fácil iniciar uma goroutine e escrever código concorrente sem a complexidade de lidar com threads nativas.

**Escalabilidade:**
- Com o gerenciamento de goroutines e o escalonamento M:N, aplicações Go podem aproveitar ao máximo os núcleos disponíveis e se adaptar a cargas de trabalho intensas de forma eficiente.







[⬅ Voltar para o README principal](/README.MD)
