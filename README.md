# dingtalk bot

    package main
    
    import (
    	"github.com/zhengkai/dingo"
    )
    
    func main() {
    
    	secret := `SEC59e4f4369a9753025609e34b4203b32ed13a85e78c0129ac8e2eb2e2010398e8`
    	token := `8f41e6348e28f26eebf1568ba8b89b167962f2ebbfa7b27690a406c06da197fe`
    
    	dingo.NewBot(token, secret).Text(`send`)
    }
