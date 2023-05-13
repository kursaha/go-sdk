## Signal API

This is a Go SDK for sending signals to the Engage Data Drive API endpoint.

### Usage

1. Import the SDK:

```go
"github.com/kursaha/go-sdk/kursaha"
"github.com/kursaha/go-sdk/kursaha/edd"
```

2. Create a new instance of the `SignalService`:

```go
kursahaClient := kursaha.NewKursaha("<YOUR-API-KEY>")
signal := kursahaClient.NewSignalService()
```

3. Use the `SendSignal` method to send a signal:

```go
signal := edd.Signal{
    EmitterID:           "unique-emitter-id",
    StepNodeID:          "step-node-id",
    Data:                map[string]interface{}{},
    EventFlowIdentifier: "uuid",
}
err := client.Edd.SendSignal([]edd.Signal{})
if err != nil {
    print(err.Error())
} else {
    print("successfully send signal!")
}
```

## Sendmail API

This is a Go SDK for sending emails using the Mailkeets API endpoint.

### Usage

1. Import the SDK:

```go
"github.com/kursaha/go-sdk/kursaha"
"github.com/kursaha/go-sdk/kursaha/mailkeets"
```

2. Create a new instance of the `Kursaha`:

```go
kursahaClient := kursaha.NewKursaha("<YOUR-API-KEY>")
```

3. Use the `SendEmail` method to send an email:

```go
dto := mailkeets.MailRequestDto{
    FromName:          "Bob",
    FromAddress:       "bob@joe.com",
    To:                "receiver@gmail.com",
    Subject:           "Hello from golang",
    ContentType:       "text/plain",
    Body:              "This is go lang sample",
    RequestIdentifier: uuid.New().String(),
    UnsubscribedList:  "",
}
err := client.Mk.SendEmail(dto)
if err != nil {
    print(err.Error())
} else {
    print("successfully send mail!")
}
```