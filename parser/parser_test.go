package parser

import "testing"

func TestBytesToInt64(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult int64
		wantErr    bool
	}{
		// TODO: Add test cases.
		// {
		// 	name:       "TestBytesToInt64-1000",
		// 	args:       args{bs: []byte("123456")},
		// 	wantResult: 123456,
		// 	wantErr:    false,
		// },
		// {
		// 	name:       "TestBytesToInt64-1001",
		// 	args:       args{bs: []byte("-123456")},
		// 	wantResult: -123456,
		// 	wantErr:    false,
		// },
		// {
		// 	name:       "TestBytesToInt64-0800",
		// 	args:       args{bs: []byte("-01234567")},
		// 	wantResult: -01234567,
		// 	wantErr:    false,
		// },
		// {
		// 	name:       "TestBytesToInt64-0801",
		// 	args:       args{bs: []byte("01234567")},
		// 	wantResult: 01234567,
		// 	wantErr:    false,
		// },
		// {
		// 	name:       "TestBytesToInt64-0802",
		// 	args:       args{bs: []byte("-012345678")},
		// 	wantResult: 0,
		// 	wantErr:    true,
		// },
		// {
		// 	name:       "TestBytesToInt64-0803",
		// 	args:       args{bs: []byte("012345678")},
		// 	wantResult: 0,
		// 	wantErr:    true,
		// },
		{
			name:       "TestBytesToInt64-1600",
			args:       args{bs: []byte("0x1234567aBf")},
			wantResult: 0x1234567aBf,
			wantErr:    false,
		},
		{
			name:       "TestBytesToInt64-1601",
			args:       args{bs: []byte("-0x1234567aBf")},
			wantResult: -0x1234567aBf,
			wantErr:    false,
		},
		// {
		// 	name:       "TestBytesToInt64-1602",
		// 	args:       args{bs: []byte("-0x1234567aBg")},
		// 	wantResult: 0,
		// 	wantErr:    true,
		// },
		// {
		// 	name:       "TestBytesToInt64-1603",
		// 	args:       args{bs: []byte("-0x1234567aBg")},
		// 	wantResult: 0,
		// 	wantErr:    true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := BytesToInt64(tt.args.bs)
			if (err != nil) != tt.wantErr {
				t.Errorf("name= %s, BytesToInt64() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("name= %s, BytesToInt64() = %v, want %v", tt.name, gotResult, tt.wantResult)
			}
		})
	}
}
