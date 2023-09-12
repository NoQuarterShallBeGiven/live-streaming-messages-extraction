package _go

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"reflect"
	"testing"
)

func TestConfig_message(t *testing.T) {
	type fields struct {
		Commands map[string]string
	}
	type args struct {
		mess []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Commands: tt.fields.Commands,
			}
			c.message(tt.args.mess)
		})
	}
}

func TestConfig_parseCommand(t *testing.T) {
	type fields struct {
		Commands map[string]string
	}
	type args struct {
		command string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Commands: tt.fields.Commands,
			}
			c.parseCommand(tt.args.command)
		})
	}
}

func TestConfig_start(t *testing.T) {
	type fields struct {
		Commands map[string]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Commands: tt.fields.Commands,
			}
			c.start()
		})
	}
}

func TestConfig_writeConfig(t *testing.T) {
	type fields struct {
		Commands map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Commands: tt.fields.Commands,
			}
			if err := c.writeConfig(); (err != nil) != tt.wantErr {
				t.Errorf("writeConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_colorString(t *testing.T) {
	type args struct {
		FgColor int
		input   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testing colorString()", args: struct {
			FgColor int
			input   string
		}{FgColor: 1, input: "test"}, want: "\u001B[38;5;1mtest\u001B[0m"}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorString(tt.args.FgColor, tt.args.input); got != tt.want {
				fmt.Println(got)
				t.Errorf("colorString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testing if file exist", args: struct{ filename string }{filename: "main_test.go"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exists(tt.args.filename); got != tt.want {
				t.Errorf("exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notification(t *testing.T) {
	type args struct {
		title string
		note  string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testing notification", args: struct {
			title string
			note  string
		}{title: "test", note: "test Passed"}},
	}
	app.NewWithID("odysee-livechat")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notification(tt.args.title, tt.args.note)
		})
	}
}

func Test_speak(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testing speech", args: struct{ input string }{input: "testing speech, test passed!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			speak(tt.args.input)
		})
	}
}

func Test_system(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "testing system()", args: struct{ command string }{command: "whoami"}, want: []byte("scott\n"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := system(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("system() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("system() got = %v, want %v", got, tt.want)
			}
		})
	}
}
