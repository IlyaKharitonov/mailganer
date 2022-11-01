package celeryWorkerService

import (
	"bytes"
	"html/template"
	"log"
	"runtime"
	"sync"

	"github.com/gocelery/gocelery"
	gomail "gopkg.in/mail.v2"

	"mailganer/internal/entities"
	"mailganer/internal/services/subscriberService"
)

type storage interface {
	GetTemplate(id uint)(string, error)
}

type celeryWorker struct {
	client *gocelery.CeleryClient
	ss *subscriberService.SubscriberService
	dialer  *gomail.Dialer
	storage storage
	config  *entities.Config
}

func NewCeleryWorker(celery *gocelery.CeleryClient, ss *subscriberService.SubscriberService, dialer *gomail.Dialer, storage storage, config *entities.Config)*celeryWorker {
		return &celeryWorker{celery, ss, dialer, storage, config}
}

func (cw *celeryWorker)Run()error {

	go func() {
		cw.client.Register("Send", cw.Send)
		cw.client.StartWorker()
		cw.client.WaitForStopWorker()
	}()

	return nil
}
//templateID uint, recipientsID[]uint
func (cw *celeryWorker)Send(args map[string]interface{}){

	templateID, recipientsID := cw.ParseArgs(args)
	//получаем шаблон из базы
	blankTemplate, err := cw.storage.GetTemplate(templateID)
	if err != nil {
		log.Println("(cw *celeryWorker)Send #1 %s", err.Error())
	}

	tmpl, err := template.New("mail").Parse(blankTemplate)
	if err != nil {
		log.Println("(cw *celeryWorker)Send #2 %s", err.Error())
	}

	recipients, err := cw.ss.GetStorage().GetList(recipientsID)
	if err != nil {
		log.Println("(cw *celeryWorker)Send #3 %s", err.Error())
	}

	messages := getMessages(recipients, tmpl, cw.config.Smtp.Login)
	err = cw.dialer.DialAndSend(messages...)
	if err != nil {
		log.Println("(cw *celeryWorker)Send #4 %s", err.Error())
	}

}

func (cw *celeryWorker)ParseArgs(args map[string]interface{})(uint, []uint){
	var (
		templateID uint
		recipientsID = make([]uint,0)
	)

	t, ok := args["templateID"]
	if !ok {
		log.Println("(cw *celeryWorker)ParseArg #1; Error no args['templateID'] elem")
	}
	tfloat, ok := t.(float64)
	if !ok {
		log.Println("(cw *celeryWorker)ParseArg #2; Error no convert['templateID'] elem")
	}
	templateID = uint(tfloat)

	recipients, ok := args["recipientsID"]
	if !ok {
		log.Println("(cw *celeryWorker)ParseArg #3; Error no args['recipientsID'] elem")
	}

	switch rec := recipients.(type) {
	case []interface{}:
		for _,r := range rec{
			rfloat, ok := r.(float64)
			if !ok {
				log.Println("(cw *celeryWorker)ParseArg #2; Error no convert['templateID'] elem")
			}
			recipientsID = append(recipientsID, uint(rfloat))
		}
	default:
		log.Println("(cw *celeryWorker)ParseArg #3; Error not []inteface{}['templateID'] elem")
	}

	return templateID, recipientsID
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