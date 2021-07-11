# togodo

Schedule notifications.

```db.json``` describes timestamp (YYYYMMDDhhmm), recipient and notification message.
Togodo will send a notification email, signaling the event deadline

Events:
```json
{
    "events": [{
            "InEvt": 202107111956,
            "Rcpt": "albertodeprezzo@gmail.com",
            "Msg": "reminder 1"
        },
        {
            "InEvt": 202107111957,
            "Rcpt": "albertodeprezzo@gmail.com",
            "Msg": "reminder 2"
        },
        {
            "InEvt": 202107111958,
            "Rcpt": "albertodeprezzo@gmail.com",
            "Msg": "reminder 3"
        },
        {
            "InEvt": 202107111959,
            "Rcpt": "albertodeprezzo@gmail.com",
            "Msg": "reminder 4"
        }
    ]
}
```

Deploy Development 

```bash
# deploy
# Requirements:
export EMAIL_HOST="smtp.gmail.com"
export EMAIL_URI="smtp.gmail.com:587"
export EMAIL_USERNAME="ABCDEF@gmail.com"
export EMAIL_PASSWD=XXXXXXXXX;
# run 
go run main.go


```

Test
```golang
go test -v ./...
```



### TODO
- [ ] Events CRUD
- [ ] Web panel
- [ ] No need of service reload on events update