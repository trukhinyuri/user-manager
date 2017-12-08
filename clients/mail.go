package clients

import (
	"git.containerum.net/ch/mail-templater/upstreams"
	"git.containerum.net/ch/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

type MailClient struct {
	rest *resty.Client
	log  *logrus.Logger
}

func NewMailClient(serverUrl string) *MailClient {
	log := logrus.WithField("component", "mail_client").Logger
	client := resty.New().SetHostURL(serverUrl).SetLogger(log.Writer())
	return &MailClient{
		rest: client,
		log:  log,
	}
}

func (mc *MailClient) sendOneTemplate(tmplName string, recipient *upstreams.Recipient) error {
	req := &upstreams.SendRequest{}
	req.Delay = 0
	req.Message.Recipients = append(req.Message.Recipients, *recipient)
	resp, err := mc.rest.R().
		SetBody(req).
		SetResult(upstreams.SendResponse{}).
		SetError(utils.Error{}).
		Post("/templates/" + tmplName)
	if err != nil {
		return err
	}
	return resp.Error().(*utils.Error)
}

func (mc *MailClient) SendConfirmationMail(recipient *upstreams.Recipient) error {
	mc.log.Info("Sending confirmation mail to", recipient.Email)
	return mc.sendOneTemplate("confirm_reg", recipient)
}

func (mc *MailClient) SendActivationMail(recipient *upstreams.Recipient) error {
	mc.log.Info("Sending confirmation mail to", recipient.Email)
	return mc.sendOneTemplate("activate_acc", recipient)
}