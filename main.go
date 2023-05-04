package main

import (
	"fmt"
	"os"
	"context"
	"log"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println((event.Parameters))
		fmt.Println(event.Event)
		fmt.Println()


	}
}


func main(){
	
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5196625699367-5211392229715-9H74Tlcmb7jdSjcNrRpipnSF")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A056X0G354G-5223964825729-7739601cfd26188a6f1d8ee8fd798f10366140485b26b35523868fd69005fd56")

	
bot := slacker.NewClient(os.Getenv(("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

go printCommandEvents(bot.printCommandEvents())

bot.command("ping")


}