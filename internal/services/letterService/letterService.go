package letterService

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"runtime"
	"sync"

	"mailganer/internal/entities"
	"mailganer/internal/services/subscriberService"

	gomail "gopkg.in/mail.v2"
)

type (
	storage interface {
		GetTemplate(ctx context.Context, id uint)(string, error)
	}

	queue interface {

	}

	taskQueue interface {
		Send(templateID uint, recipientsID []uint)error
	}

	letterService struct {
		storage storage
		queue queue
		tq taskQueue
		ss *subscriberService.SubscriberService
	}

)

func NewLetterService(s storage, q queue, tq taskQueue, ss *subscriberService.SubscriberService)*letterService{
	return &letterService{s, q, tq, ss}
}

func(ms *letterService)Send(ctx context.Context, letter *entities.Letter)error{

	err := ms.tq.Send(letter.TemplateID, letter.RecipientsID)
	if err != nil {
		return fmt.Errorf("(ms *letterService)Send #1 %s", err.Error())
	}


	var (
		from = "AugustWarm16@yandex.ru"
		host = "smtp.yandex.ru"
		pass = "AugustWarm1643"
		port = 587
	)

	//получаем шаблон из базы
	blankTemplate, err := ms.storage.GetTemplate(ctx, letter.TemplateID)
	if err != nil {
		return fmt.Errorf("(ms *MailService)Send #1 %s", err.Error())
	}

	tmpl, err := template.New("mail").Parse(blankTemplate)
	if err != nil {
		return fmt.Errorf("(ms *MailService)Send #2 %s", err.Error())
	}

	recipients, err := ms.ss.GetStorage().GetList(letter.RecipientsID)
	if err != nil {
		return fmt.Errorf("(ms *MailService)Send #3 %s", err.Error())
	}

	messages := getMessages(recipients, tmpl, from)

	d := gomail.NewDialer(host, port, from, pass)
	err = d.DialAndSend(messages...)
	if err != nil {
		log.Printf("(ms *MailService)Send #4 %s", err.Error())
	}

	return nil
}

func getMessages(recipients []*entities.Subscriber, tmpl *template.Template, from string)[]*gomail.Message{
	var (
		messages = make([]*gomail.Message, 0)
		wg = &sync.WaitGroup{}
		mu = &sync.Mutex{}
		subscriberChan = make(chan *entities.Subscriber, len(recipients))
		workerPoolSize = runtime.NumCPU()
	)

	for i:=0; i<=workerPoolSize; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for r := range subscriberChan{
				prepareMessage(r, tmpl, from, mu, messages)
			}
		}()
	}

	for _,r := range recipients{
		subscriberChan <- r
	}

	close(subscriberChan)

	wg.Wait()
	return messages
}

func prepareMessage(r *entities.Subscriber, t *template.Template, from string, mu *sync.Mutex, messages []*gomail.Message){
	var (
		buffer = &bytes.Buffer{}
		m = gomail.NewMessage()
	)

	err := t.Execute(buffer, r)
	if err != nil {
		log.Printf("(ms *MailService)Send #4 %s", err.Error())
	}

	m.SetHeader("From", from)
	m.SetHeader("To", r.Email)
	//m.SetHeader("Subject", r.Subject)
	m.SetHeader("text/html", buffer.String())

	mu.Lock()
	messages = append(messages, m)
	mu.Unlock()
}