package beater

import ("github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
        "github.com/elastic/beats/libbeat/common"
        "github.com/fatih/structs"
        "encoding/json"
	"github.com/davecgh/go-spew/spew"
        "time"
	"bufio"
        "os")

type JsonBeat struct {
	ConfigSettings ConfigSettings
	Events publisher.Client
	Done chan struct{}
}

type ConfigSettings struct {
	Input JbConfig
}

type JbConfig struct {
	BeatConfig Config
}

type Config struct {
	Json_Elasticsearch_Type_Field *string
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

	logp.Info(spew.Sdump(j.ConfigSettings))

	esTypeField := "type" // default value
	if j.ConfigSettings.Input.BeatConfig.Json_Elasticsearch_Type_Field != nil {
	        esTypeField = *j.ConfigSettings.Input.BeatConfig.Json_Elasticsearch_Type_Field
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var esTypeVal string
		var payload interface {}

		// parse the json
		var f interface{}
		err := json.Unmarshal(s.Bytes(), &f)

		// figure out the payload and type
		if err != nil { // unparseable JSON events
			logp.Err("Not sending unparsable-JSON: line=%s err=%v", s.Text(), err)
			continue
		/*
			logp.Err("Sending unparsable JSON as raw unparseable_event: line=%s err=%v", s.Text(), err)
			esTypeVal = "unparseable_event"
			payload = s.Text()
                */
		} else { // parseable JSON events
			m := f.(map[string]interface{}) // cast
			val, ok := m[esTypeField]
			if ok == true {
				esTypeVal = val.(string)
				payload = f
			} else {
				esTypeVal = "unknown_event"
				payload = f
			}
		}

		event := common.MapStr{}
		event["@timestamp"] = common.Time(time.Now()) // mark time of event
		event["type"] = esTypeVal
		event["count"] = "1"
		event["params"] = structs.Map(j.ConfigSettings.Input.BeatConfig)
		event["payload"] = payload
		j.Events.PublishEvent(event)
        }

	// Do not exit until all pending events have been sent (or attempted)
	// TODO: Fix this to not sleep, but to check pending events
	time.Sleep(5 * time.Second)

	logp.Info("exiting")
	return err
}

func (j *JsonBeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (j *JsonBeat) Stop() {
	close(j.Done)
}
