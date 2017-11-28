package pycodec

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	input := `[{'user': {'id': 123, 'name': 'hello'}, 'created_datetime': '2017-11-20T15:18:08.941064+00:00'}, {'user': ['test', 42, 42.42, None, False, True]}]`
	type entry struct {
		User interface{} `pycodec:"user"`
	}
	var got []entry
	err := Unmarshal([]byte(input), &got)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	want := []entry{
		{User: map[string]interface{}{"id": 123, "name": "hello"}},
		{User: []interface{}{"test", 42, 42.42, nil, false, true}},
	}

	gotJson, _ := json.Marshal(got)
	wantJson, _ := json.Marshal(want)

	if string(gotJson) != string(wantJson) {
		t.Errorf("error\ngot  %+v\nwant %+v", got, want)
	}
}
