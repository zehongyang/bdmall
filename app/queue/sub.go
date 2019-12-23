package queue

import "bdmall/utils"

var EmailChan = make(chan map[string]string)

//异步发送邮件
func Sub()  {
	for data := range EmailChan {
		go func(to,subject,content string) {
			_ = utils.SendEmail(to, subject, content)
		}(data["to"],data["subject"],data["content"])
	}
}
