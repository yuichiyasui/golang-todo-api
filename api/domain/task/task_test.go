package task

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		id          string
		title       string
		description string
		status      TaskStatus
	}
	tests := []struct {
		name    string
		args    args
		want    *Task
		wantErr bool
	}{
		{
			name: "ドメインオブジェクトを作成する",
			args: args{
				id:          "1",
				title:       "タスク名",
				description: "タスク説明",
				status:      "TODO",
			},
			want: &Task{
				id:          "1",
				title:       "タスク名",
				description: "タスク説明",
				status:      "TODO",
			},
			wantErr: false,
		},
		{
			name: "statusが無効な値の場合TODOが設定される",
			args: args{
				id:          "1",
				title:       "タスク名",
				description: "タスク説明",
				status:      "invalid_status",
			},
			want: &Task{
				id:          "1",
				title:       "タスク名",
				description: "タスク説明",
				status:      "TODO",
			},
			wantErr: false,
		},
		{
			name: "titleが空の場合エラーを返す",
			args: args{
				id:          "1",
				title:       "",
				description: "タスク説明",
				status:      "TODO",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "titleが30文字を超える場合エラーを返す",
			args: args{
				id:          "1",
				title:       makeText(31),
				description: "タスク説明",
				status:      "TODO",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "descriptionが500文字を超える場合エラーを返す",
			args: args{
				id:          "1",
				title:       "タスク名",
				description: makeText(501),
				status:      "TODO",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.id, tt.args.title, tt.args.description, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeText(length int) string {
	var text string
	for i := 0; i < length; i++ {
		text += "a"
	}
	return text
}
