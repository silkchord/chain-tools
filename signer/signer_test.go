package singer

import (
	"reflect"
	"testing"

	"github.com/songqianba/chain-tools/protocol/btm"
	"github.com/songqianba/chain-tools/signer/element"
	"github.com/songqianba/chain-tools/signer/mocks"
)

func TestSigner_Sign(t *testing.T) {
	type fields struct {
		memc string
	}
	type args struct {
		rawTx            string
		signInstructions []element.SigningInstructionTemp
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantWitnesses [][]string
		wantErr       bool
	}{
		{
			name: "test singer linger desire",
			fields: fields{
				memc: "salon invest obtain tomato submit jelly clump drum grape usual heavy wear",
			},
			args: args{
				rawTx:            mocks.MockMap["2c1826322ec81743e7932fe856ac2c27017de1fa0fbb3b6e1969210a7df3e3bc"].RawTx,
				signInstructions: mocks.MockMap["2c1826322ec81743e7932fe856ac2c27017de1fa0fbb3b6e1969210a7df3e3bc"].Instructions,
			},
			wantWitnesses: [][]string{{"4a6bb2f7d744a37b273c995c691a1b161a75fe1602c02f8868f092dc6f3390a482ff32d7b810a38ddfd9c31ac5852008dc32b6bbe7e1f774a3e09b99661f7302"}},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewSigner(tt.fields.memc)
			if err != nil {
				t.Errorf("NewSigner() error = %v", err)
				return
			}

			tx, err := btm.UnserializeTx(tt.args.rawTx)
			if err != nil {
				t.Errorf("decode raw_tx error = %v", err)
				return
			}

			gotWitnesses, err := s.Sign(btm.NewWrapTx(tx), tt.args.signInstructions...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotWitnesses, tt.wantWitnesses) {
				t.Errorf("Sign() gotWitnesses = %v, want %v", gotWitnesses, tt.wantWitnesses)
			}
		})
	}
}
