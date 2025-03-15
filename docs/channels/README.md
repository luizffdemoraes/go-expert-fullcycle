**Channels.:**

##  Por que usar Channels?

Permitem sincronização entre goroutines.
Evitam o uso explícito de locks (Mutex).
Tornam o código mais legível e fácil de entender.
Seguem o princípio de Go: "Não compartilhe memória, compartilhe mensagens".

___


## 🔹 Tipos de Channels
1️⃣ Unbuffered (Sem Buffer - Padrão)

A comunicação é bloqueante, ou seja, a goroutine pausa até que o dado seja recebido.

2️⃣ Buffered (Com Buffer)

O canal pode armazenar múltiplos valores antes de bloquear.


[⬅ Voltar para o README principal](/README.MD)



