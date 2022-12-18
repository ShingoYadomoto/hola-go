package hola_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHupaiCalculater_peko(t *testing.T) {
	type fields struct {
		standard StandardHoluPattern
	}
	tests := []struct {
		name   string
		fields fields
		want   []HandType
	}{
		{
			name: "true case. ipeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{一盃口},
		},
		{
			name: "false case. ipeko case. 3 same shuntsu",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{白, 白, 白}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
					},
				},
			},
			want: []HandType{一盃口},
		},
		{
			name: "true case. ryanpeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
					},
				},
			},
			want: []HandType{二盃口},
		},
		{
			name: "false case. ipeko case. 3 same shuntsu",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{東, 東, 東}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{},
		},
		{
			name: "false case. fulou case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{一萬, 二萬, 三萬}},
					},
					FulouMentsu: FulouMentsuList{
						{mentsu: mentsu{pais: []pai{白, 白, 白}}},
					},
				},
			},
			want: []HandType{},
		},
		{
			name: "false case. not same shuntsu case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{東, 東, 東}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := HupaiCalculater{
				standard: tt.fields.standard,
			}
			assert.Equalf(t, tt.want, hc.peko(), "ipeko()")
		})
	}
}

func TestFullHupaiCalculater_Hupai(t *testing.T) {
	type fields struct {
		fullParrern FullHoluPattern
		zhuangfeng  Zhuangfeng
		zifeng      Zifeng
		isTsumo     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   PossibleAllHands
	}{
		{
			name: "",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{四萬, 五萬, 六萬}, holuPai: &四萬},
								{pais: []pai{三筒, 四筒, 五筒}},
								{pais: []pai{五索, 六索, 七索}},
								{pais: []pai{八索, 八索, 八索}},
							},
							Head:    二索,
							HoluPai: 四萬,
						},
					},
					IsTsumo: true,
				},
				zhuangfeng: 東場,
				zifeng:     東家,
				isTsumo:    true,
			},
			want: []AllHands{{門前清自摸和}},
		},

		{
			name: "",
			fields: fields{
				fullParrern: FullHoluPattern{
					Standard: []StandardHoluPattern{
						{
							Mentsu: []mentsu{
								{pais: []pai{七索, 八索, 九索}, holuPai: &七索},
								{pais: []pai{七索, 八索, 九索}},
								{pais: []pai{七筒, 八筒, 九筒}},
								{pais: []pai{五筒, 六筒, 七筒}},
							},
							Head:    四索,
							HoluPai: 七索,
						},
					},
					IsTsumo: true,
				},
				zhuangfeng: 東場,
				zifeng:     東家,
				isTsumo:    true,
			},
			want: []AllHands{{門前清自摸和, 一盃口}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fhc := FullHupaiCalculater{
				fullParrern: tt.fields.fullParrern,
				zhuangfeng:  tt.fields.zhuangfeng,
				zifeng:      tt.fields.zifeng,
				isTsumo:     tt.fields.isTsumo,
			}
			assert.Equalf(t, tt.want.String(), fhc.Hupai().String(), "Hupai()")
		})
	}
}
