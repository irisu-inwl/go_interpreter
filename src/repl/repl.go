package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

// Sengoku Collection is a divine anime.
const MONKEY_FACE = `            G#         &BGBB#&       #B5#         
           B7!5G  #GY?!~^^^^~!?5B& GJ?!!J         
          #7J?7?Y!^::::^^^^^^^::^~J?777!Y         
         &577?!^::^^^~^^^^^^^:^^^::^7J7~~?P&      
       #J~:^~^:::::~!^~!^^^::~:::::::~!^^::~J#    
     &J^:^~^::^^^^~?~!77^^^^^!!^^^^^^:~~^^^::^Y   
   &5^::^~^^~^^^^~?J!J77~^^~^~?!^^^7^^^77~^:^::7# 
 #Y~^^^^~~!7^^^~77?7YJ?77~^7?7?77~^?~^:!?77~^^^:7 
  ###P^~?777^^^!?J7.^!??77!?777~^777!:^7?777!^^J?B
     5^77777^^~7?!!!?7^~?7??~!??!^7?!~!?77777~:? &
     Y~77777~:!7?Y5PGY!.^?7^:!5G55J?777?777777^!& 
     P!77777?!~7GY~5PY:..^:..~Y5Y~GJ77?777777!:7  
     &Y7777777!7J^^77!.......^77!.!?777775YG&P7G  
       577?PY7^!!..::..........:.:7!~Y77G         
        YY#  &GJ?:......:^::.....:YP& PB          
        &        BY!^.........:!YB                
                    &B5?~^^^?G&                   
                     &#B?!!!G                     
                    #PGY~~~^7GG#                  
                   BPPP~....:YPPB&                
                &P7?PPPP!..^5PPP7!YB              
               5!~~~?PPPPJ75PP57~~~~5             
               5~~~~~75PPPPPPPJ!~~~~Y             
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
		// Lexer output
		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
