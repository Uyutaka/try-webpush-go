package main

import (
	"encoding/json"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
)

const (
	subscription    = `{"endpoint":"https://updates.push.services.mozilla.com/wpush/v2/gAAAAABj_…Xt3KGEfklFYvXaNFl17DYIzbNUb5buKJFB8So_2DTV8PZ-zb19QaNbum1Eyo","expirationTime":null,"keys":{"auth":"hsIQLvabvs3CqPcTx0DI7w","p256dh":"BGzvyOc0NcDKiWJfSNaOxshxCfktf2nmuSrIJ8PLNGngzxQgyjiKrSZHfu8SB_LQmwDUquDZ577TzboZjffxzpY"}}`
	vapidPublicKey  = "BMZvcUluk_cARMEsN4nbO-MLxx3s124NYnrCqhE3XT5QPCxQUFCEqh1T4WylJ4jCFZDUh68RiMA1-cFDWwRzMtY"
	vapidPrivateKey = "k0U5ulyDJGJZsg_pVgy0lhR97rXlFl_HPhQ1tOnYbLA"
)

func main() {
	// Decode subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
		Subscriber:      "ssss@example.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             86400,
	})
	if err != nil {
		// TODO: Handle error
		log.Println(err)

		log.Println("error happened")
		return
	}
	defer resp.Body.Close()
}
