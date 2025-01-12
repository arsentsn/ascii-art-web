// ascii.TestAsciiPrint
package ascii

import (
	"testing"
)

func TestAsciiPrint(t *testing.T) {
	tests := []struct {
		input      string
		bannerFile string
		want       string
	}{
		{"", "standard", ""},
		{"\n", "standard", "\n"},
		{"\nhello", "standard", `
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`},
		{"hello", "standard", ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`},
		{"HELLO", "standard", ` _    _   ______   _        _         ____   
| |  | | |  ____| | |      | |       / __ \  
| |__| | | |__    | |      | |      | |  | | 
|  __  | |  __|   | |      | |      | |  | | 
| |  | | | |____  | |____  | |____  | |__| | 
|_|  |_| |______| |______| |______|  \____/  
                                             
                                             
`},
		{"1Hello 2There", "standard", `     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`},
		{"Hello\nThere", "standard", ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`},
		{"Hello\n\nThere", "standard", ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`},
		{"{Hello & There #}", "standard", `   __  _    _          _   _                                _______   _                                    _  _    __    
  / / | |  | |        | | | |                 ___          |__   __| | |                                 _| || |_  \ \   
 | |  | |__| |   ___  | | | |   ___          ( _ )            | |    | |__     ___   _ __    ___        |_  __  _|  | |  
/ /   |  __  |  / _ \ | | | |  / _ \         / _ \/\          | |    |  _ \   / _ \ | '__|  / _ \        _| || |_    \ \ 
\ \   | |  | | |  __/ | | | | | (_) |       | (_>  <          | |    | | | | |  __/ | |    |  __/       |_  __  _|   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/         \___/\/          |_|    |_| |_|  \___| |_|     \___|         |_||_|    | |  
  \_\                                                                                                              /_/   
                                                                                                                         
`},
		{"hello There 1 to 2!", "standard", ` _              _   _                 _______   _                                            _                           _  
| |            | | | |               |__   __| | |                                 _        | |                  ____   | | 
| |__     ___  | | | |   ___            | |    | |__     ___   _ __    ___        / |       | |_    ___         |___ \  | | 
|  _ \   / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \       | |       | __|  / _ \          __) | | | 
| | | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/       | |       \ |_  | (_) |        / __/  |_| 
|_| |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|       |_|        \__|  \___/        |_____| (_) 
                                                                                                                            
                                                                                                                            
`},
		{"{|}~", "standard", `   __  _  __     /\/| 
  / / | | \ \   |/\/  
 | |  | |  | |        
/ /   | |   \ \       
\ \   | |   / /       
 | |  | |  | |        
  \_\ | | /_/         
      |_|             
`},
		{"RGB", "standard", ` _____     _____   ____   
|  __ \   / ____| |  _ \  
| |__) | | |  __  | |_) | 
|  _  /  | | |_ | |  _ <  
| | \ \  | |__| | | |_) | 
|_|  \_\  \_____| |____/  
                          
                          
`},
		{"123", "shadow", `                       
  _|   _|_|   _|_|_|   
_|_| _|    _|       _| 
  _|     _|     _|_|   
  _|   _|           _| 
  _| _|_|_|_| _|_|_|   
                       
                       
`},
		{"/(\")", "thinkertoy", `         o o    
    o  / | | \  
   /  o       o 
  o   |       | 
 /    o       o 
o      \     /  
                
                
`},
	}

	for _, test := range tests {
		MakeAsciiArtTable("../banners/" + test.bannerFile + ".txt")
		if got := AsciiPrint(test.input); got != test.want {
			t.Errorf("\n got:\n %s  \nwant:\n %s", got, test.want)
		}
	}
}
