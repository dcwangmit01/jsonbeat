package beater

import ("github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
        "github.com/elastic/beats/libbeat/common"
        "github.com/fatih/structs"
        "time"
        "os"
        "encoding/json"
	"bufio")

type JsonBeat struct {
	ConfigSettings ConfigSettings
	Events publisher.Client
	Done chan struct{}
}

func New() *JsonBeat {
	return &JsonBeat{}
}

func (j *JsonBeat) Config(b *beat.Beat) error {
	err := cfgfile.Read(&j.ConfigSettings, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}
	return nil
}

func (j *JsonBeat) Setup(b *beat.Beat) error {
	j.Events = b.Events
	j.Done = make(chan struct{})
	return nil
}

func (j *JsonBeat) Run(b *beat.Beat) error {
        var err error
	err = nil

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
	    logp.Info("line %", s.Text())
	    
            var f interface{}
            err := json.Unmarshal(s.Bytes(), &f)

            m := f.(map[string]interface{})
            logp.Info("type%", m["event_type"])
                
            if err != nil {
                    logp.Err("Failed to parse as JSON line=%s err=%v", s.Text(), err)
            }
                    
  	    // parse the result	
            event := common.MapStr{}
	    event["@timestamp"] = common.Time(time.Now()) // mark time of event
	    event["type"] = m["event_type"]
	    event["count"] = "1"
	    event["params"] = structs.Map(j.ConfigSettings.Input)
	    event["payload"] = f

	    j.Events.PublishEvent(event)
        }

	// Wait until the signal to quit
	/*
	select {
	case <-t.Done:
		logp.Info("quitting")
		return nil
	}
	*/
	return err
}

func (j *JsonBeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (j *JsonBeat) Stop() {
	close(j.Done)
}
