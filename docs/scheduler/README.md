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







[⬅ Voltar para o README principal](/README.MD)
