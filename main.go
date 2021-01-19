// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
// 	"strconv"
	"time"
	"math/rand"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
// 				quota, err := bot.GetMessageQuota().Do()
// 				if err != nil {
// 					log.Println("Quota err:", err)
// 				}else {
// 					log.Println("Quota err:", linebot.EventTypeMessage)
// 				}								
				if message.Text == "我要看妞妞"{
					t1 := time.NewTimer(3 * time.Second)
					rand.Seed(time.Now().Unix())
					image := []string{
						"https://i.imgur.com/z5yOT1e.jpg",
						"https://i.imgur.com/Wxa4lzR.jpg",
						"https://i.imgur.com/NPQy2Cn.jpg",
						"https://i.imgur.com/VjV59Dk.jpg",
						"https://i.imgur.com/fGvy47i.jpg",
						"https://i.imgur.com/pPQI1LN.jpg",
						"https://i.imgur.com/pEjnhSy.jpg",
					}
					<- t1.C
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("汪！"), linebot.NewImageMessage(image[rand.Intn(len(image))] , image[rand.Intn(len(image))])).Do(); err != nil {
					log.Print(err)
					}
				}
				if message.Text == "叫我起床"{
					t1 := time.NewTimer(20 * time.Second)
					<- t1.C
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("起床！！！"), linebot.NewImageMessage("https://i.imgur.com/URlBkOk.jpg" , "https://i.imgur.com/URlBkOk.jpg")).Do(); err != nil {
					log.Print(err)
					}
				}
			}
		}
	}
}
