### Variáveis, Valores & Tipos

- Estrutura básica: 
- package main.
- func main: é aqui que tudo começa, é aqui que tudo acaba.
- import.

- Packages:
    - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
    - Notação: pacote.Identificador. Exemplo: fmt.Println()
    - Documentação: fmt.Println.

- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Variáveis não utilizadas? Não pode, coloque _ nelas. [blank identifier]
- ...funções variádicas.

### Operador Curto de declaração

- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática, esse operador atribui um tipo baseado no valor da variável
    - Só pode repetir se houverem variáveis novas 
    - a marmota é != do assignment operator (operador de atribuição)
    - Marmota só funciona dentro de codeblocks 

- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores, operandos
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões 
    - expressão → qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope

- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
    - = para atribuir valores a variáveis já existentes